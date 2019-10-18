package main

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"sync"
	"time"
)

// DB - database struct
type DB struct {
	IntTable    map[string]*TableInt
	StringTable map[string]*TableString
	BaseStorage string

	DefaultFlushTimeout time.Duration

	eo ErrorsOut

	mx sync.RWMutex
}

// ErrorsOut log errors func
type ErrorsOut func(e error)

// CreateTableInt create int table
func (db *DB) CreateTableInt(tblType string, name string, flushTimeout time.Duration) (parErr error, intErr error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	_, ok := db.IntTable[name]
	if ok {
		return errors.New("Table with name " + name + " already exists"), nil
	}
	_, ok = db.StringTable[name]
	if ok {
		return errors.New("Table with name " + name + " already exists int String tables"), nil
	}

	ti, err := CreateTableInt(tblType, name, db.BaseStorage, filepath.ToSlash(name+"/"), flushTimeout, db.eo)

	if err != nil {
		return nil, ErrorNew("DBCreateTableInt CreateTableInt with name "+name, err)
	}

	db.IntTable[name] = ti

	return nil, nil

}

// CreateTableString create String table
func (db *DB) CreateTableString(tblType string, name string, flushTimeout time.Duration) (parErr error, intErr error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	_, ok := db.StringTable[name]
	if ok {
		return errors.New("Table with name " + name + " already exists"), nil
	}
	_, ok = db.IntTable[name]
	if ok {
		return errors.New("Table with name " + name + " already exists int Int tables"), nil
	}

	ti, err := CreateTableString(tblType, name, db.BaseStorage, filepath.ToSlash(name+"/"), flushTimeout, db.eo)

	if err != nil {
		return nil, ErrorNew("DBCreateTableString CreateTableString with name "+name, err)
	}

	db.StringTable[name] = ti

	return nil, nil

}

// CreateDB create DB struct
func CreateDB(path string, defaultFlushTimeout time.Duration, eo ErrorsOut) (db *DB, err error) {

	d, e, err := FileLoad(path + "struct.json")
	if err != nil {
		return db, ErrorNew("CreateDB load file "+path+"struct.json", err)
	}

	if e {
		err = json.Unmarshal(d, &db)
		if err != nil {
			return db, ErrorNew("CreateDB unmarshal "+path+"struct.json", err)
		}
		db.BaseStorage = path
		db.eo = eo
	} else {
		db = &DB{
			BaseStorage: path,
			IntTable:    make(map[string]*TableInt),
			StringTable: make(map[string]*TableString),

			DefaultFlushTimeout: defaultFlushTimeout,

			eo: eo,
		}
	}

	for n, v := range db.IntTable {
		v.BaseFolder = db.BaseStorage
		err = v.Init(db.eo)
		if err != nil {
			return db, ErrorNew("CreateDB init int table "+n+"", err)
		}
	}

	for n, v := range db.StringTable {
		v.BaseFolder = db.BaseStorage
		err = v.Init(db.eo)
		if err != nil {
			return db, ErrorNew("CreateDB init string table "+n+"", err)
		}
	}

	return db, nil
}

// Flush db struct on disk
func (db *DB) Flush() error {
	db.mx.Lock()
	defer db.mx.Unlock()

	err := MkDirIfNotExists(db.BaseStorage, 0760)
	if err != nil {
		return ErrorNew("DBFlush mkdir "+db.BaseStorage, err)
	}

	b, err := json.Marshal(db)

	if err != nil {
		return ErrorNew("DBFlush json marshal "+db.BaseStorage, err)
	}

	err = FileReplace(db.BaseStorage+"struct.json", b, 0660)

	if err != nil {
		return ErrorNew("DBFlush json write file "+db.BaseStorage+"struct.json", err)
	}

	return nil
}

// TableIntGet Get int table
func (db *DB) TableIntGet(tblName string) (tbl *TableInt, parErr error) {
	db.mx.RLock()
	tbl, ok := db.IntTable[tblName]
	db.mx.RUnlock()

	if !ok {
		return nil, errors.New("Table with name " + tblName + " not exists")
	}

	return tbl, nil

}

// TableStringGet Get int table
func (db *DB) TableStringGet(tblName string) (tbl *TableString, parErr error) {
	db.mx.RLock()
	tbl, ok := db.StringTable[tblName]
	db.mx.RUnlock()

	if !ok {
		return nil, errors.New("Table with name " + tblName + " not exists")
	}

	return tbl, nil

}
