package main

import (
	"encoding/json"
	"errors"
)

// IntType column type int
const IntType = "int"

// StringType column type string
const StringType = "string"

// QueryCreateIndex Create index query
type QueryCreateIndex struct {
	TableName  string `json:"tbl_name"`
	ColumnName string `json:"col_name"`

	// int or string
	ColumnType string `json:"col_type"`
}

// CreateIndexOnTableIntFromJSON creates string index on table from json
func (db *DB) CreateIndexOnTableIntFromJSON(d []byte) (parErr error, intErr error) {

	tblP := QueryCreateIndex{}

	err := json.Unmarshal(d, &tblP)
	if err != nil {
		return ErrorNew("fail json ", err), nil
	}

	if tblP.ColumnType != IntType && tblP.ColumnType != StringType {
		return errors.New("column type must be int or string"), nil
	}

	tbl, err := db.TableIntGet(tblP.TableName)
	if err != nil {
		return ErrorNew("Table get: ", err), nil
	}

	if tblP.ColumnType == IntType {

		parErr, intErr = tbl.IndexIntCreate(tblP.ColumnName)

	} else {
		parErr, intErr = tbl.IndexStringCreate(tblP.ColumnName)
	}

	return parErr, intErr
}

// CreateIndexOnTableStringFromJSON creates string index on table from json
func (db *DB) CreateIndexOnTableStringFromJSON(d []byte) (parErr error, intErr error) {

	tblP := QueryCreateIndex{}

	err := json.Unmarshal(d, &tblP)
	if err != nil {
		return ErrorNew("fail json ", err), nil
	}

	if tblP.ColumnType != IntType && tblP.ColumnType != StringType {
		return errors.New("column type must be int or string"), nil
	}

	tbl, err := db.TableStringGet(tblP.TableName)
	if err != nil {
		return ErrorNew("Table get: ", err), nil
	}

	if tblP.ColumnType == IntType {

		parErr, intErr = tbl.IndexIntCreate(tblP.ColumnName)

	} else {
		parErr, intErr = tbl.IndexStringCreate(tblP.ColumnName)
	}

	return parErr, intErr
}
