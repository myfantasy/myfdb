package main

import (
	"encoding/json"
	"errors"
	"sync"
)

// TableMapStringName - name of TableMapString type
const TableMapStringName = "table_map_string"

// TableMapString - table with int64 key
type TableMapString struct {
	Data          map[string]ItemString
	IntIndexes    map[string]*IndexStringInt
	StringIndexes map[string]*IndexStringString
	mx            sync.RWMutex
	mxF           sync.Mutex
	needSave      bool

	storageName string

	TableStruct TableStructData
}

// GetType type name return
func (tmi *TableMapString) GetType() string {
	return TableMapStringName
}

// CreateTableMapString - create index
func CreateTableMapString(storageName string) (i *TableMapString) {
	i = &TableMapString{
		Data:          make(map[string]ItemString),
		IntIndexes:    make(map[string]*IndexStringInt),
		StringIndexes: make(map[string]*IndexStringString),
		storageName:   storageName,
		needSave:      true}

	i.TableStruct.NamesIndexInt = make(map[string]string)
	i.TableStruct.NamesIndexString = make(map[string]string)

	return i
}

// IndexIntAdd - add new index
func (tmi *TableMapString) IndexIntAdd(colName string, ixName string) error {
	tmi.mx.Lock()

	for k, v := range tmi.TableStruct.NamesIndexInt {
		if ixName == v {
			tmi.mx.Unlock()
			return errors.New("Index exists " + ixName + " (" + colName + ") stor " + tmi.storageName)
		}

		if colName == k {
			tmi.mx.Unlock()
			return errors.New("Index exists " + ixName + " on column " + colName + " stor " + tmi.storageName)
		}
	}

	tmi.TableStruct.NamesIndexInt[colName] = ixName
	tmi.IntIndexes[colName] = CreateIndexStringInt()
	tmi.needSave = true
	tmi.mx.Unlock()

	return nil
}

// IndexStringAdd - add new index
func (tmi *TableMapString) IndexStringAdd(colName string, ixName string) error {
	tmi.mx.Lock()

	for k, v := range tmi.TableStruct.NamesIndexString {
		if ixName == v {
			tmi.mx.Unlock()
			return errors.New("Index exists " + ixName + " (" + colName + ") stor " + tmi.storageName)
		}

		if colName == k {
			tmi.mx.Unlock()
			return errors.New("Index exists " + ixName + " on column " + colName + " stor " + tmi.storageName)
		}
	}

	tmi.TableStruct.NamesIndexString[colName] = ixName
	tmi.StringIndexes[colName] = CreateIndexStringString()
	tmi.needSave = true
	tmi.mx.Unlock()

	return nil
}

// Len - items count
func (tmi *TableMapString) Len() (r int, err error) {

	tmi.mx.RLock()

	r = len(tmi.Data)

	tmi.mx.RUnlock()

	return r, nil
}

// Get one item
func (tmi *TableMapString) Get(key string) (itm ItemString, ok bool, err error) {

	tmi.mx.RLock()

	s, ok := tmi.Data[key]
	if ok {
		itm = s
	}

	tmi.mx.RUnlock()

	return itm, ok, nil
}

// GetAll items limitStop
func (tmi *TableMapString) GetAll(limitStop int, includeDeleted bool) (itms []ItemString, parErr error, intErr error) {

	tmi.mx.RLock()
	for _, v := range tmi.Data {
		if includeDeleted || !v.IsRemoved {
			itms = append(itms, v)
			if limitStop > -1 && len(itms) >= limitStop {
				break
			}
		}
	}
	tmi.mx.RUnlock()

	return itms, parErr, intErr
}

// GetList list keys limit by limitStop
func (tmi *TableMapString) GetList(keys []string, limitStop int) (itms []ItemString, err error) {

	var itm ItemString
	var ok bool
	for _, key := range keys {
		itm, ok, err = tmi.Get(key)
		if err != nil {
			return itms, err
		}
		if ok {
			itms = append(itms, itm)
			if limitStop > -1 && len(itms) >= limitStop {
				break
			}
		}
	}

	return itms, nil
}

// GetKeysByIndexInt list items by col vals limit by limitStop
func (tmi *TableMapString) GetKeysByIndexInt(colName string, val int64, limitStop int) (keys []string, parErr error, intErr error) {
	tmi.mx.RLock()

	ix, ok := tmi.IntIndexes[colName]
	if ok {
		keys, intErr = ix.Get(val)
	}

	tmi.mx.RUnlock()

	if !ok {
		return keys, errors.New("GetListByIndexInt: index " + colName + " not found"), nil
	}
	return keys, parErr, intErr

}

// GetListByIndexInt list items by col vals limit by limitStop
func (tmi *TableMapString) GetListByIndexInt(colName string, vals []int64, limitStop int) (itms []ItemString, parErr error, intErr error) {

	var itmT []ItemString
	for _, val := range vals {
		keys, parErr, intErr := tmi.GetKeysByIndexInt(colName, val, limitStop)
		if parErr != nil || intErr != nil {
			return itms, parErr, intErr
		}
		itmT, intErr = tmi.GetList(keys, limitStop-len(itms))
		if parErr != nil || intErr != nil {
			return itms, parErr, intErr
		}
		itms = append(itms, itmT...)
		if limitStop > -1 && len(itms) >= limitStop {
			break
		}

	}

	return itms, parErr, intErr
}

// GetKeysByIndexString list items by col vals limit by limitStop
func (tmi *TableMapString) GetKeysByIndexString(colName string, val string, limitStop int) (keys []string, parErr error, intErr error) {
	tmi.mx.RLock()

	ix, ok := tmi.StringIndexes[colName]
	if ok {
		keys, intErr = ix.Get(val)
	}

	tmi.mx.RUnlock()

	if !ok {
		return keys, errors.New("GetListByIndexInt: index " + colName + " not found"), nil
	}
	return keys, parErr, intErr

}

// GetListByIndexString list items by col vals limit by limitStop
func (tmi *TableMapString) GetListByIndexString(colName string, vals []string, limitStop int) (itms []ItemString, parErr error, intErr error) {

	var itmT []ItemString
	for _, val := range vals {
		keys, parErr, intErr := tmi.GetKeysByIndexString(colName, val, limitStop)
		if parErr != nil || intErr != nil {
			return itms, parErr, intErr
		}
		itmT, intErr = tmi.GetList(keys, limitStop-len(itms))
		if parErr != nil || intErr != nil {
			return itms, parErr, intErr
		}
		itms = append(itms, itmT...)
		if limitStop > -1 && len(itms) >= limitStop {
			break
		}

	}

	return itms, parErr, intErr
}

// Set new item
func (tmi *TableMapString) Set(itm ItemString) (itmOut ItemString, isNew bool, isRvEqual bool, err error) {

	if itm.Rv == 0 {
		itm.Rv = RvGet()
	}

	tmi.mx.Lock()

	itmOld, okOld := tmi.Data[itm.Key]

	if okOld {
		if itmOld.Rv == itm.Rv {
			tmi.mx.Unlock()
			return itmOld, false, true, nil
		}

		if itmOld.Rv > itm.Rv {
			tmi.mx.Unlock()
			return itmOld, false, false, nil
		}

	}

	tmi.Data[itm.Key] = itm

	//IntIndex
	{
		var okOO bool
		var okON bool
		var okLO bool
		var okLN bool

		var vOO int64
		var vON int64
		var vLO []int64
		var vLN []int64

		for k := range tmi.TableStruct.NamesIndexInt {
			if okOld {
				vOO, okOO = itmOld.FieldsInt[k]
				vLO, okLO = itmOld.FieldsIntA[k]
			}
			vON, okON = itm.FieldsInt[k]
			vLN, okLN = itm.FieldsIntA[k]

			if okOO || okLO || okON || okLN {
				if okOO {
					vLO = append(vLO, vOO)
				}

				if okON {
					vLN = append(vLN, vON)
				}

				tmi.IntIndexes[k].Set(itm.Key, vLN, vLO)
			}
		}
	}

	//StringIndex
	{
		var okOO bool
		var okON bool
		var okLO bool
		var okLN bool

		var vOO string
		var vON string
		var vLO []string
		var vLN []string

		for k := range tmi.TableStruct.NamesIndexString {
			if okOld {
				vOO, okOO = itmOld.FieldsString[k]
				vLO, okLO = itmOld.FieldsStringA[k]
			}
			vON, okON = itm.FieldsString[k]
			vLN, okLN = itm.FieldsStringA[k]

			if okOO || okLO || okON || okLN {
				if okOO {
					vLO = append(vLO, vOO)
				}

				if okON {
					vLN = append(vLN, vON)
				}

				tmi.StringIndexes[k].Set(itm.Key, vLN, vLO)
			}
		}
	}

	tmi.needSave = true
	tmi.mx.Unlock()

	return itm, true, false, nil
}

// Flush - flush on disk
func (tmi *TableMapString) Flush() error {

	tmi.mxF.Lock()
	defer tmi.mxF.Unlock()
	if tmi.needSave {
		err := MkDirIfNotExists(tmi.storageName, 0760)
		if err != nil {
			return ErrorNew("TableMapStringFlush mkdir "+tmi.storageName, err)
		}

		tmi.mx.Lock()
		tmi.needSave = false
		tmi.mx.Unlock()

		tmi.mx.RLock()
		b, err := json.Marshal(tmi.TableStruct)
		tmi.mx.RUnlock()
		if err != nil {
			tmi.needSave = true
			return ErrorNew("TableMapStringFlush json TableStruct marshal "+tmi.storageName, err)
		}

		err = FileReplace(tmi.storageName+"struct.json", b, 0660)

		if err != nil {
			tmi.needSave = true
			return ErrorNew("TableMapStringFlush json TableStruct write file "+tmi.storageName+"struct.json", err)
		}
		{
			itype := "ii"
			inames := make(map[string]string)

			tmi.mx.RLock()
			for k, v := range tmi.TableStruct.NamesIndexInt {
				inames[k] = v
			}
			tmi.mx.RUnlock()

			for k, v := range inames {
				iName := tmi.storageName + "i_" + itype + "_" + v + ".json"

				tmi.mx.RLock()
				err = tmi.IntIndexes[k].Flush(iName)
				tmi.mx.RUnlock()

				if err != nil {
					tmi.needSave = true
					return ErrorNew("TableMapStringFlush json Index marshal "+iName, err)
				}

			}
		}
		{
			itype := "ss"
			inames := make(map[string]string)

			tmi.mx.RLock()
			for k, v := range tmi.TableStruct.NamesIndexString {
				inames[k] = v
			}
			tmi.mx.RUnlock()

			for k, v := range inames {
				iName := tmi.storageName + "i_" + itype + "_" + v + ".json"

				tmi.mx.RLock()
				err = tmi.StringIndexes[k].Flush(iName)
				tmi.mx.RUnlock()

				if err != nil {
					tmi.needSave = true
					return ErrorNew("TableMapStringFlush json Index marshal "+iName, err)
				}

			}
		}

		tmi.mx.RLock()
		b, err = json.Marshal(tmi.Data)
		tmi.mx.RUnlock()

		if err != nil {
			tmi.needSave = true
			return ErrorNew("TableMapStringFlush json Data marshal "+tmi.storageName, err)
		}

		err = FileReplace(tmi.storageName+"data.json", b, 0660)

		if err != nil {
			tmi.needSave = true
			return ErrorNew("TableMapStringFlush json Data write file "+tmi.storageName+"data.json", err)
		}
	}
	return nil
}

// PathSet - set path for table
func (tmi *TableMapString) PathSet(path string) {
	tmi.storageName = path
}

// StorageExists - exists struct file
func (tmi *TableMapString) StorageExists() (bool, error) {

	tmi.mxF.Lock()
	defer tmi.mxF.Unlock()
	tmi.mx.Lock()
	defer tmi.mx.Unlock()

	// load struct

	_, e, err := FileLoad(tmi.storageName + "struct.json")

	if err != nil {
		return false, ErrorNew("TableMapStringLoad load file "+tmi.storageName+"struct.json", err)
	}

	return e, nil
}

// Load - load from file
func (tmi *TableMapString) Load() error {

	tmi.mxF.Lock()
	defer tmi.mxF.Unlock()
	tmi.mx.Lock()
	defer tmi.mx.Unlock()

	// load struct

	d, e, err := FileLoad(tmi.storageName + "struct.json")

	if err != nil {
		return ErrorNew("TableMapStringLoad load file "+tmi.storageName+"struct.json", err)
	}

	if !e {
		return errors.New("TableMapStringLoad Not Found file " + tmi.storageName + "struct.json")
	}

	err = json.Unmarshal(d, &tmi.TableStruct)
	if err != nil {
		return ErrorNew("TableMapStringLoad unmarshal "+tmi.storageName+"data.json", err)
	}

	// load data

	d, e, err = FileLoad(tmi.storageName + "data.json")

	if err != nil {
		return ErrorNew("TableMapStringLoad load file "+tmi.storageName+"data.json", err)
	}

	if !e {
		return errors.New("TableMapStringLoad Not Found file " + tmi.storageName + "data.json")
	}

	err = json.Unmarshal(d, &tmi.Data)
	if err != nil {
		return ErrorNew("TableMapStringLoad unmarshal "+tmi.storageName+"data.json", err)
	}

	{
		tmi.IntIndexes = make(map[string]*IndexStringInt)
		itype := "ii"
		for k, v := range tmi.TableStruct.NamesIndexInt {
			iName := tmi.storageName + "i_" + itype + "_" + v + ".json"

			tmi.IntIndexes[k] = CreateIndexStringInt()
			err = tmi.IntIndexes[k].Load(iName)
			if err != nil {
				return ErrorNew("TableMapStringLoad load index file "+iName, err)
			}
		}
	}

	{
		tmi.StringIndexes = make(map[string]*IndexStringString)
		itype := "ss"
		for k, v := range tmi.TableStruct.NamesIndexString {
			iName := tmi.storageName + "i_" + itype + "_" + v + ".json"

			tmi.StringIndexes[k] = CreateIndexStringString()
			err = tmi.StringIndexes[k].Load(iName)
			if err != nil {
				return ErrorNew("TableMapStringLoad load index file "+iName, err)
			}
		}
	}

	tmi.needSave = false

	return nil
}
