package main

import (
	"encoding/json"
	"errors"
)

// ItemInt - one row with key string
type ItemInt struct {
	Key       int64           `json:"key,omitempty"`
	Data      json.RawMessage `json:"d,omitempty"`
	Rv        int64           `json:"rv,omitempty"`
	IsRemoved bool            `json:"rm,omitempty"`
}

// ItemIntMakeFromJSONObject create ItemInt from json
func ItemIntMakeFromJSONObject(msg json.RawMessage, is *ItemStruct) (itm ItemInt, err error) {

	m := make(map[string]json.RawMessage)

	err = json.Unmarshal(msg, &m)

	if err != nil {
		return itm, err
	}

	var v json.RawMessage
	var ok bool

	if is == nil {
		v, ok = m["key"]
	} else {
		v, ok = m[is.KeyName]
	}

	if !ok && is == nil {
		return itm, errors.New("key not found")
	}
	if !ok {
		return itm, errors.New("key (" + is.KeyName + ") not found")
	}

	err = json.Unmarshal(v, &itm.Key)

	if err != nil {
		return itm, ErrorNew("key unmurshal fail", err)
	}

	itm.Data = msg
	return itm, err

}

// ItemIntStruct - one row with key string
type ItemIntStruct struct {
	Key           int64               `json:"key,omitempty"`
	FieldsInt     map[string]int64    `json:"fi,omitempty"`
	FieldsString  map[string]string   `json:"fs,omitempty"`
	FieldsIntA    map[string][]int64  `json:"fia,omitempty"`
	FieldsStringA map[string][]string `json:"fsa,omitempty"`
	Data          *[]byte             `json:"d,omitempty"`
	Rv            int64               `json:"rv,omitempty"`
	IsRemoved     bool                `json:"rm,omitempty"`
}

// ItemIntStat - one row with key string
type ItemIntStat struct {
	Key       int64 `json:"key,omitempty"`
	Rv        int64 `json:"rv,omitempty"`
	IsRemoved bool  `json:"rm,omitempty"`
}

// Stat - get stat object
func (i ItemInt) Stat() ItemIntStat {
	return ItemIntStat{
		Key:       i.Key,
		Rv:        i.Rv,
		IsRemoved: i.IsRemoved,
	}
}
