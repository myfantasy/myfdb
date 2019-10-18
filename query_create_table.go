package main

import (
	"encoding/json"
	"time"
)

// QueryCreateTable Create table query
type QueryCreateTable struct {
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	FlushTimeout time.Duration `json:"flush_timeout"`
}

// StructGet get database primary struct
func (db *DB) StructGet() ([]byte, error) {
	db.mx.Lock()
	defer db.mx.Unlock()

	return json.Marshal(db)
}

// CreateTableIntFromJSON creates table from json
func (db *DB) CreateTableIntFromJSON(d []byte) (parErr error, intErr error) {

	tblP := QueryCreateTable{
		FlushTimeout: db.DefaultFlushTimeout,
	}

	err := json.Unmarshal(d, &tblP)
	if err != nil {
		return ErrorNew("fail json ", err), nil
	}

	parErr, intErr = db.CreateTableInt(tblP.Type, tblP.Name, tblP.FlushTimeout)
	if parErr != nil {
		return parErr, nil
	}
	if intErr != nil {
		return nil, ErrorNew("fail create table ", intErr)
	}

	err = db.Flush()
	if err != nil {
		return nil, ErrorNew("fail flush db struct ", err)
	}

	return nil, nil
}

// CreateTableStringFromJSON creates table from json
func (db *DB) CreateTableStringFromJSON(d []byte) (parErr error, intErr error) {

	tblP := QueryCreateTable{
		FlushTimeout: db.DefaultFlushTimeout,
	}

	err := json.Unmarshal(d, &tblP)
	if err != nil {
		return ErrorNew("fail json ", err), nil
	}

	parErr, intErr = db.CreateTableString(tblP.Type, tblP.Name, tblP.FlushTimeout)
	if parErr != nil {
		return parErr, nil
	}
	if intErr != nil {
		return nil, ErrorNew("fail create table ", intErr)
	}

	err = db.Flush()
	if err != nil {
		return nil, ErrorNew("fail flush db struct ", err)
	}

	return nil, nil
}
