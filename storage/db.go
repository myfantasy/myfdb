package storage

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"github.com/myfantasy/myfdb/generator"
	"github.com/myfantasy/myfdb/logger"

	"github.com/myfantasy/mdp"
	"github.com/myfantasy/myfdb/ftools"
)

// DB database
type DB struct {
	ServerName string

	IntTables map[string]*IntAbstractTable

	DefaultSavePath string
	FileSavePath    string

	DBFlushTimeout time.Duration

	mx sync.RWMutex

	stopChan chan<- bool
}

// DBWrite wirite struct
type DBWrite struct {
	ServerName string `json:"server_name"`

	DBFlushTimeout time.Duration `json:"flush_timeout"`

	IntTables map[string]mdp.TableDefinition `json:"int_tables"`

	DefaultSavePath string `json:"default_save_path"`
}

// DBWriteGet write struct get
func (db *DB) DBWriteGet() (dbw DBWrite, err *mdp.Error) {

	dbw.ServerName = db.ServerName
	dbw.DefaultSavePath = db.DefaultSavePath
	dbw.DBFlushTimeout = db.DBFlushTimeout

	dbw.IntTables = make(map[string]mdp.TableDefinition)
	for k, v := range db.IntTables {
		s, pe, ie := v.GetStructFull()
		if pe != nil {
			return dbw, pe
		}
		if ie != nil {
			return dbw, ie
		}
		dbw.IntTables[k] = s
	}

	return dbw, nil
}

// DBLoadFromWriteStruct DB load from struct
func DBLoadFromWriteStruct(fileSavePath string, dbFlushTimeout time.Duration, stopChan chan<- bool) (db *DB, err *mdp.Error) {
	dpath := fileSavePath + "struct.json"

	b, ex, e := ftools.FileLoad(dpath)

	if ex {
		if e != nil {
			return nil, mdp.ErrorNew("DBLoadFromWriteStruct json load file "+dpath+"", e)
		}

		var dbw DBWrite

		dbw.DBFlushTimeout = dbFlushTimeout

		e = json.Unmarshal(b, &dbw)
		if err != nil {
			return nil, mdp.ErrorNew("DBLoadFromWriteStruct unmarshal "+dpath+"", e)
		}

		db, err := dbw.DBLoadFromWriteStruct(fileSavePath)
		db.stopChan = stopChan

		return db, err

	}

	db = &DB{
		IntTables:       make(map[string]*IntAbstractTable),
		DefaultSavePath: fileSavePath,
		FileSavePath:    fileSavePath,
		DBFlushTimeout:  dbFlushTimeout,
		ServerName:      "db_" + strconv.Itoa(int(generator.RvGet())),
		stopChan:        stopChan,
	}

	err = db.Flush()

	return db, err
}

// DBLoadFromWriteStruct DB load from struct
func (dbw DBWrite) DBLoadFromWriteStruct(fileSavePath string) (db *DB, err *mdp.Error) {

	db = &DB{
		ServerName:      dbw.ServerName,
		DefaultSavePath: dbw.DefaultSavePath,
		FileSavePath:    fileSavePath,
		DBFlushTimeout:  dbw.DBFlushTimeout,
	}
	db.IntTables = make(map[string]*IntAbstractTable)

	for k, v := range dbw.IntTables {
		t, pe, ie := IntAbstractTableLoad(v)
		if pe != nil {
			return db, pe
		}
		if ie != nil {
			return db, ie
		}
		db.IntTables[k] = t
	}

	return db, nil
}

// Flush on disk
func (db *DB) Flush() (err *mdp.Error) {

	dbw, err := db.DBWriteGet()
	if err != nil {
		return err
	}

	b, e := json.MarshalIndent(dbw, "", "\t")

	dpath := db.FileSavePath + "struct.json"

	if e != nil {
		return mdp.ErrorNew("DB.Flush write file "+dpath+" fail Marshal JSON ", e)
	}

	e = ftools.FileReplace(dpath, b, 0660)

	if e != nil {
		return mdp.ErrorNew("DB.Flush write json write file "+dpath+"", e)
	}

	return nil

}

// CreateTable Create new table
func (db *DB) CreateTable(td mdp.TableDefinition) (tdOut mdp.TableDefinition, paramsErr *mdp.Error, internalErr *mdp.Error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	if td.TableName == "" {
		paramsErr = mdp.ErrorS("Name must be not empty")
		return tdOut, paramsErr, internalErr
	}

	if _, ok := db.IntTables[td.TableName]; ok {
		paramsErr = mdp.ErrorS("Table already exists")
		return tdOut, paramsErr, internalErr
	}

	if td.KeyType == mdp.KeyTypeInt {

		iat, paramsErr, internalErr := IntAbstractTableCreate(td, db)

		if paramsErr != nil || internalErr != nil {
			return tdOut, paramsErr, internalErr
		}

		db.IntTables[td.TableName] = iat

		internalErr = db.Flush()
		if internalErr != nil {
			return tdOut, paramsErr, internalErr
		}

		tdOut, paramsErr, internalErr = iat.GetStruct()
		return tdOut, paramsErr, internalErr

	}
	internalErr = mdp.ErrorS("String key not implemented")
	return tdOut, paramsErr, internalErr
}

// StructGet get struct by query
func (db *DB) StructGet(sgq mdp.StructGetQuery) (res mdp.StructGet) {

	db.mx.RLock()
	defer db.mx.RUnlock()

	if sgq.TableName != "" {

		t, ok := db.IntTables[sgq.TableName]

		if !ok {
			res.ParamsErr = mdp.ErrorS("Table not exists")
			return res
		}

		var s mdp.TableDefinition
		if sgq.LoadInternalInfo {
			s, res.ParamsErr, res.InternalErr = t.GetStructFull()
		} else {

			s, res.ParamsErr, res.InternalErr = t.GetStruct()
		}

		res.Tables = []mdp.TableDefinition{s}

	} else if sgq.LoadAll {

	}

	return res
}

// StructSet set struct by query
func (db *DB) StructSet(ssq mdp.StructSetQuery) (res mdp.StructGet) {

	if ssq.CreateTable != nil {
		tdOut, paramsErr, internalErr := db.CreateTable(*ssq.CreateTable)
		res.ParamsErr = paramsErr
		res.InternalErr = internalErr
		res.Tables = []mdp.TableDefinition{tdOut}
	}

	return res
}

// StructStorageGet get struct storage by query
func (db *DB) StructStorageGet(sgq mdp.StructStorageGetQuery) (res mdp.StructStorageGet) {

	db.mx.RLock()
	defer db.mx.RUnlock()

	if sgq.Name {
		res.Storage.StorageName = db.ServerName
	}

	if sgq.LocalPlace {
		res.Storage.LocalStoragePath = db.DefaultSavePath
	}

	return res
}

// StructStorageSet set struct storage by query
func (db *DB) StructStorageSet(sgq mdp.StructStorageSetQuery) (res mdp.StructStorageGet) {

	db.mx.Lock()
	defer db.mx.Unlock()

	if sgq.Name != "" {
		db.ServerName = sgq.Name
		err := db.Flush()
		if err != nil {
			res.InternalErr = err
			return res
		}

		res.Storage.StorageName = db.ServerName
	}

	if sgq.Stop {
		db.stopChan <- true
	}

	return res
}

// FlushAll flush all structs and tables errors out into logger.InternalProcessError
func (db *DB) FlushAll() {
	db.mx.Lock()
	defer db.mx.Unlock()
	err := db.Flush()
	if err != nil {
		logger.InternalProcessError(err)
	}

	for _, v := range db.IntTables {
		err = v.Flush()
		if err != nil {
			logger.InternalProcessError(err)
		}
	}

}

// TableIntGet get int table by name
func (db *DB) TableIntGet(tableName string) (iat *IntAbstractTable, paramsErr *mdp.Error) {
	db.mx.RLock()
	defer db.mx.RUnlock()

	var ok bool

	iat, ok = db.IntTables[tableName]
	if !ok {
		paramsErr = mdp.ErrorS("Table not exists")
	}

	return iat, paramsErr
}

// ItemGet get item by query
func (db *DB) ItemGet(igq mdp.ItemsGetQuery) (res mdp.ItemsGet) {

	if igq.IKey != 0 && igq.TableName != "" {

		iat, paramsErr := db.TableIntGet(igq.TableName)
		if paramsErr != nil {
			res.ParamsErr = paramsErr
			return
		}

		s, ok, paramsErr, internalErr := iat.GetItem(igq.IKey)

		if paramsErr != nil && internalErr != nil {
			res.ParamsErr = paramsErr
			res.InternalErr = internalErr
			return
		}

		if ok {
			res.Count = 1
			res.IItems = []mdp.ItemInt{s}
		}

	}

	return res
}

// ItemSet set item by query
func (db *DB) ItemSet(igq mdp.ItemsSetQuery) (res mdp.ItemsGet) {

	if igq.TableName != "" && igq.IItem != nil {

		iat, paramsErr := db.TableIntGet(igq.TableName)
		if paramsErr != nil {
			res.ParamsErr = paramsErr
			return
		}

		s, paramsErr, internalErr := iat.SetItem(*igq.IItem)

		if paramsErr != nil && internalErr != nil {
			res.ParamsErr = paramsErr
			res.InternalErr = internalErr
			return
		}

		res.Count = 1
		res.IItems = []mdp.ItemInt{s}

	}

	return res
}
