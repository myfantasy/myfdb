package main

import (
	"encoding/json"
	"errors"
	"sync"
)

// IndexIntString - Index from int64 indexed fileds and string base value
type IndexIntString struct {
	Ix       map[string][]int64 `json:"ix,omitempty"`
	mx       sync.RWMutex
	mxF      sync.Mutex
	needSave bool
}

// CreateIndexIntString - create index
func CreateIndexIntString() (i *IndexIntString) {
	i = &IndexIntString{Ix: make(map[string][]int64),
		needSave: true}

	return i
}

// Len indexed vals count
func (mi *IndexIntString) Len() (r int) {

	mi.mx.RLock()

	r = len(mi.Ix)

	mi.mx.RUnlock()

	return r
}

// Get one items
func (mi *IndexIntString) Get(val string) (keys []int64, err error) {

	mi.mx.RLock()

	s, ok := mi.Ix[val]
	if ok {
		keys = s
	}

	mi.mx.RUnlock()

	return keys, nil
}

// GetList list keys limit by limitStop
func (mi *IndexIntString) GetList(vals []string, limitStop int) (keys []int64, err error) {

	for _, val := range vals {
		ikeys, err := mi.Get(val)
		if err != nil {
			return keys, err
		}
		keys = append(keys, ikeys...)
		if limitStop > -1 && len(keys) > limitStop {
			break
		}
	}

	return keys, nil
}

// Set new item
func (mi *IndexIntString) Set(key int64, newValue []string, oldValue []string) error {

	mi.mx.Lock()

	mi.needSave = true

	for _, v := range oldValue {
		s, ok := mi.Ix[v]
		if ok {
			mi.Ix[v] = SliceRemoveInt(s, key)
		}
	}

	for _, v := range newValue {
		s, ok := mi.Ix[v]
		if ok {
			mi.Ix[v] = append(s, key)
		} else {
			mi.Ix[v] = append(make([]int64, 0), key)
		}
	}

	for _, v := range oldValue {
		s, ok := mi.Ix[v]
		if ok && len(s) == 0 {
			delete(mi.Ix, v)
		}
	}

	mi.mx.Unlock()

	return nil
}

// Flush - flush on disk
func (mi *IndexIntString) Flush(path string) error {

	mi.mxF.Lock()
	defer mi.mxF.Unlock()
	if mi.needSave {
		mi.mx.Lock()
		mi.needSave = false
		mi.mx.Unlock()

		mi.mx.RLock()
		b, err := json.Marshal(mi)
		mi.mx.RUnlock()

		if err != nil {
			mi.needSave = true
			return ErrorNew("IndexIntStringFlush json marshal "+path, err)
		}

		err = FileReplace(path, b, 0660)

		if err != nil {
			mi.needSave = true
			return ErrorNew("IndexIntStringFlush json write file "+path, err)
		}
	}

	return nil
}

// Load - load from file
func (mi *IndexIntString) Load(path string) error {

	mi.mxF.Lock()
	defer mi.mxF.Unlock()
	mi.mx.Lock()
	defer mi.mx.Unlock()

	d, e, err := FileLoad(path)

	if err != nil {
		return ErrorNew("IndexIntStringLoad load file "+path, err)
	}

	if !e {
		return errors.New("IndexIntStringLoad Not Found file " + path)
	}

	err = json.Unmarshal(d, mi)
	if err != nil {
		return ErrorNew("IndexIntStringLoad unmarshal "+path, err)
	}
	mi.needSave = false

	return nil
}
