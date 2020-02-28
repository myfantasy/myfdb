package main

import (
	"encoding/json"

	"github.com/myfantasy/mdp"
)

// DBItemQueryGet - process query get item
func DBItemQueryGet(data []byte) (s mdp.ItemsGet) {

	var q mdp.ItemsGetQuery

	err := json.Unmarshal(data, &q)

	if err != nil {
		s.ParamsErr = mdp.ErrorNew("DBQueryGet fail unmarchal query", err)
		return s
	}

	s = db.ItemGet(q)
	//	s.ParamsErr = mdp.ErrorS("Query not implemented")
	return s
}

// DBItemQuerySet - process query set item
func DBItemQuerySet(data []byte) (s mdp.ItemsGet) {

	var q mdp.ItemsSetQuery

	err := json.Unmarshal(data, &q)

	if err != nil {
		s.ParamsErr = mdp.ErrorNew("DBQueryGet fail unmarchal query", err)
		return s
	}

	s = db.ItemSet(q)
	//	s.ParamsErr = mdp.ErrorS("Query not implemented")
	return s
}

// DBStructQueryGet - process query struct get
func DBStructQueryGet(data []byte) (s mdp.StructGet) {

	var q mdp.StructGetQuery

	err := json.Unmarshal(data, &q)

	if err != nil {
		s.ParamsErr = mdp.ErrorNew("DBQueryGet fail unmarchal query", err)
		return s
	}

	s = db.StructGet(q)
	//	s.ParamsErr = mdp.ErrorS("Query not implemented")
	return s
}

// DBStructQuerySet - process query struct set
func DBStructQuerySet(data []byte) (s mdp.StructGet) {

	var q mdp.StructSetQuery

	err := json.Unmarshal(data, &q)

	if err != nil {
		s.ParamsErr = mdp.ErrorNew("DBQueryGet fail unmarchal query", err)
		return s
	}

	s = db.StructSet(q)
	//	s.ParamsErr = mdp.ErrorS("Query not implemented")
	return s
}

// DBStructStorageQueryGet - process query struct get
func DBStructStorageQueryGet(data []byte) (s mdp.StructStorageGet) {

	var q mdp.StructStorageGetQuery

	err := json.Unmarshal(data, &q)

	if err != nil {
		s.ParamsErr = mdp.ErrorNew("DBQueryGet fail unmarchal query", err)
		return s
	}

	s = db.StructStorageGet(q)
	//	s.ParamsErr = mdp.ErrorS("Query not implemented")
	return s
}

// DBStructStorageQuerySet - process query struct set
func DBStructStorageQuerySet(data []byte) (s mdp.StructStorageGet) {

	var q mdp.StructStorageSetQuery

	err := json.Unmarshal(data, &q)

	if err != nil {
		s.ParamsErr = mdp.ErrorNew("DBQueryGet fail unmarchal query", err)
		return s
	}

	s = db.StructStorageSet(q)
	//	s.ParamsErr = mdp.ErrorS("Query not implemented")
	return s
}
