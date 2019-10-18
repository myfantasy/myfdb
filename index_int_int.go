package main

import (
	"encoding/json"
	"errors"
	"sync"
)

// IndexIntInt - Index from int64 indexed fileds and int64 base value
type IndexIntInt struct {
	Ix       map[int64][]int64 `json:"ix,omitempty"`
	mx       sync.RWMutex
	mxF      sync.Mutex
	needSave bool
}

// CreateIndexIntInt - create index
func CreateIndexIntInt() (i *IndexIntInt) {
	i = &IndexIntInt{Ix: make(map[int64][]int64),
		needSave: true}

	return i
}

// Len indexed vals count
func (mi *IndexIntInt) Len() (r int) {

	mi.mx.RLock()

	r = len(mi.Ix)

	mi.mx.RUnlock()

	return r
}

// Get one items
func (mi *IndexIntInt) Get(val int64) (keys []int64, err error) {

	mi.mx.RLock()

	s, ok := mi.Ix[val]
	if ok {
		keys = s
	}

	mi.mx.RUnlock()

	return keys, nil
}

// GetList list keys limit by limitStop
func (mi *IndexIntInt) GetList(vals []int64, limitStop int) (keys []int64, err error) {

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
func (mi *IndexIntInt) Set(key int64, newValue []int64, oldValue []int64) error {

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
func (mi *IndexIntInt) Flush(path string) error {

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
			return ErrorNew("IndexIntIntFlush json marshal "+path, err)
		}

		err = FileReplace(path, b, 0660)

		if err != nil {
			mi.needSave = true
			return ErrorNew("IndexIntIntFlush json write file "+path, err)
		}
	}

	return nil
}

// Load - load from file
func (mi *IndexIntInt) Load(path string) error {

	mi.mxF.Lock()
	defer mi.mxF.Unlock()
	mi.mx.Lock()
	defer mi.mx.Unlock()

	d, e, err := FileLoad(path)

	if err != nil {
		return ErrorNew("IndexIntIntLoad load file "+path, err)
	}

	if !e {
		return errors.New("IndexIntIntLoad Not Found file " + path)
	}

	err = json.Unmarshal(d, mi)
	if err != nil {
		return ErrorNew("IndexIntIntLoad unmarshal "+path, err)
	}
	mi.needSave = false

	return nil
}
