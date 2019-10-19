package main

import (
	"encoding/json"
)

// QueryAddOrRemoveToken Create or Remove Token query
type QueryAddOrRemoveToken struct {
	Token string `json:"token"`
}

//func (db *DB) CheckToken()

// RMTokenFromJSON remove token from json
func (db *DB) RMTokenFromJSON(d []byte) (parErr error, intErr error) {

	tblP := QueryAddOrRemoveToken{}

	err := json.Unmarshal(d, &tblP)
	if err != nil {
		return ErrorNew("fail json ", err), nil
	}

	db.RMToken(tblP.Token)

	err = db.Flush()
	if err != nil {
		return nil, ErrorNew("fail flush db struct ", err)
	}

	return nil, nil
}

// AddTokenFromJSON creates token from json
func (db *DB) AddTokenFromJSON(d []byte) (parErr error, intErr error) {

	tblP := QueryAddOrRemoveToken{}

	err := json.Unmarshal(d, &tblP)
	if err != nil {
		return ErrorNew("fail json ", err), nil
	}

	db.AddToken(tblP.Token)

	err = db.Flush()
	if err != nil {
		return nil, ErrorNew("fail flush db struct ", err)
	}

	return nil, nil
}
