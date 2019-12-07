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
	IntTable    map[string]*TableInt    `json:"int_tables"`
	StringTable map[string]*TableString `json:"string_tables"`
	BaseStorage string                  `json:"base_storage"`

	NameInCluster string `json:"cluster_name"`

	DefaultFlushTimeout time.Duration `json:"default_flush_timeout"`

	MasterTokens []string `json:"master_tokens"`
	tokenOk      map[string]bool

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

	db.tokenOk = make(map[string]bool)

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

// CheckToken check token exists
func (db *DB) CheckToken(token string) bool {
	ok := false
	exit := false
	db.mx.RLock()
	if len(db.MasterTokens) == 0 {
		ok = true
	} else if token == "" {
		exit = true
	} else {
		_, ok = db.tokenOk[token]
	}
	db.mx.RUnlock()

	if ok || exit {
		return ok
	}

	h := HashGet(token)
	db.mx.RLock()
	for _, v := range db.MasterTokens {
		if v == h {
			ok = true
			break
		}
	}
	db.mx.RUnlock()

	if ok {
		db.mx.Lock()
		db.tokenOk[token] = true
		db.mx.Unlock()
	}

	return ok
}

// RMToken remove exists token
func (db *DB) RMToken(token string) {

	h := HashGet(token)
	db.mx.RLock()

	db.MasterTokens = SliceRemoveString(db.MasterTokens, h)
	db.tokenOk = make(map[string]bool)
	db.mx.RUnlock()
}

// AddToken Add new token
func (db *DB) AddToken(token string) {

	if token == "" {
		return
	}
	ok := db.CheckToken(token)
	h := HashGet(token)
	db.mx.RLock()

	if len(db.MasterTokens) == 0 || !ok {
		db.MasterTokens = append(db.MasterTokens, h)
	}

	db.mx.RUnlock()
}
