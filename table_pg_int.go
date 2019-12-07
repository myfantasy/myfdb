package main

import (
	"sync"
	"time"
)

// TablePGIntName - name of TablePGIntName type
const TablePGIntName = "table_map_int"

// TablePGInt - table with int64 key based on page
type TablePGInt struct {
	Data          map[int64]ItemInt
	IntIndexes    map[string]*IndexIntInt
	StringIndexes map[string]*IndexIntString

	mx       sync.RWMutex
	mxF      sync.Mutex
	needSave bool

	storageName string

	TableStruct TableStructData `json:"table_struct"`

	ItmStruct *ItemStruct `json:"item_struct"`
}

// PageInt page of data
type PageInt struct {
	FirstKey    int64 `json:"f_key"`
	LastKey     int64 `json:"l_key"`
	mx          sync.RWMutex
	storageName string

	items      *[]ItemInt
	needSave   bool
	lastUnload time.Time
	lastFlush  time.Time
}

// UnLoad unload page from memory
func (pi *PageInt) UnLoad() (err error) {

	pi.mx.Lock()
	if !pi.needSave {
		pi.items = nil
		pi.lastUnload = time.Now()
	}
	pi.mx.Unlock()

	return nil
}

// Flush Flush on disk
func (pi *PageInt) Flush() (err error) {

	pi.mx.Lock()
	defer pi.mx.Unlock()
	if pi.needSave {
		pi.items = nil
		pi.lastFlush = time.Now()
	}

	return nil
}
