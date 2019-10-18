package main

import (
	"encoding/json"
	"errors"
	"sync"
)

// IndexStringString - Index from int64 indexed fileds and string base value
type IndexStringString struct {
	Ix       map[string][]string `json:"ix,omitempty"`
	mx       sync.RWMutex
	mxF      sync.Mutex
	needSave bool
}

// CreateIndexStringString - create index
func CreateIndexStringString() (i *IndexStringString) {
	i = &IndexStringString{Ix: make(map[string][]string),
		needSave: true}

	return i
}

// Len indexed vals count
func (mi *IndexStringString) Len() (r int) {

	mi.mx.RLock()

	r = len(mi.Ix)

	mi.mx.RUnlock()

	return r
}

// Get one items
func (mi *IndexStringString) Get(val string) (keys []string, err error) {

	mi.mx.RLock()

	s, ok := mi.Ix[val]
	if ok {
		keys = s
	}

	mi.mx.RUnlock()

	return keys, nil
}

// GetList list keys limit by limitStop
func (mi *IndexStringString) GetList(vals []string, limitStop int) (keys []string, err error) {

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
func (mi *IndexStringString) Set(key string, newValue []string, oldValue []string) error {

	mi.mx.Lock()

	mi.needSave = true

	for _, v := range oldValue {
		s, ok := mi.Ix[v]
		if ok {
			mi.Ix[v] = SliceRemoveString(s, key)
		}
	}

	for _, v := range newValue {
		s, ok := mi.Ix[v]
		if ok {
			mi.Ix[v] = append(s, key)
		} else {
			mi.Ix[v] = append(make([]string, 0), key)
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
func (mi *IndexStringString) Flush(path string) error {

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
			return ErrorNew("IndexStringStringFlush json marshal "+path, err)
		}

		err = FileReplace(path, b, 0660)

		if err != nil {
			mi.needSave = true
			return ErrorNew("IndexStringStringFlush json write file "+path, err)
		}
	}

	return nil
}

// Load - load from file
func (mi *IndexStringString) Load(path string) error {

	mi.mxF.Lock()
	defer mi.mxF.Unlock()
	mi.mx.Lock()
	defer mi.mx.Unlock()

	d, e, err := FileLoad(path)

	if err != nil {
		return ErrorNew("IndexStringStringLoad load file "+path, err)
	}

	if !e {
		return errors.New("IndexStringStringLoad Not Found file " + path)
	}

	err = json.Unmarshal(d, mi)
	if err != nil {
		return ErrorNew("IndexStringStringLoad unmarshal "+path, err)
	}
	mi.needSave = false

	return nil
}
