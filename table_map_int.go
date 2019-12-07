package main

import (
	"encoding/json"
	"errors"
	"sync"
)

// TableMapIntName - name of TableMapInt type
const TableMapIntName = "table_map_int"

// TableMapInt - table with int64 key
type TableMapInt struct {
	Data          map[int64]ItemInt
	IntIndexes    map[string]*IndexIntInt
	StringIndexes map[string]*IndexIntString
	mx            sync.RWMutex
	mxF           sync.Mutex
	needSave      bool

	storageName string

	TableStruct TableStructData

	ItmStruct *ItemStruct
}

// TableStructData - struct table
type TableStructData struct {
	NamesIndexInt    map[string]string `json:"name_of_int_indexes,omitempty"`
	NamesIndexString map[string]string `json:"name_of_string_indexes,omitempty"`
}

// GetType type name return
func (tmi *TableMapInt) GetType() string {
	return TableMapIntName
}

// CreateTableMapInt - create index
func CreateTableMapInt(storageName string) (i *TableMapInt) {
	i = &TableMapInt{
		Data:          make(map[int64]ItemInt),
		IntIndexes:    make(map[string]*IndexIntInt),
		StringIndexes: make(map[string]*IndexIntString),
		storageName:   storageName,
		needSave:      true}

	i.TableStruct.NamesIndexInt = make(map[string]string)
	i.TableStruct.NamesIndexString = make(map[string]string)

	return i
}

// IndexIntAdd - add new index
func (tmi *TableMapInt) IndexIntAdd(colName string, ixName string) error {
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
	tmi.IntIndexes[colName] = CreateIndexIntInt()
	tmi.needSave = true
	tmi.mx.Unlock()

	return nil
}

// IndexStringAdd - add new index
func (tmi *TableMapInt) IndexStringAdd(colName string, ixName string) error {
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
	tmi.StringIndexes[colName] = CreateIndexIntString()
	tmi.needSave = true
	tmi.mx.Unlock()

	return nil
}

// Len - items count
func (tmi *TableMapInt) Len() (r int, err error) {

	tmi.mx.RLock()

	r = len(tmi.Data)

	tmi.mx.RUnlock()

	return r, nil
}

// Get one item
func (tmi *TableMapInt) Get(key int64) (itm ItemInt, ok bool, err error) {

	tmi.mx.RLock()

	s, ok := tmi.Data[key]
	if ok {
		itm = s
	}

	tmi.mx.RUnlock()

	return itm, ok, nil
}

// GetAll items limitStop
func (tmi *TableMapInt) GetAll(limitStop int, includeDeleted bool) (itms []ItemInt, parErr error, intErr error) {

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
func (tmi *TableMapInt) GetList(keys []int64, limitStop int) (itms []ItemInt, err error) {

	var itm ItemInt
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
func (tmi *TableMapInt) GetKeysByIndexInt(colName string, val int64, limitStop int) (keys []int64, parErr error, intErr error) {
	tmi.mx.RLock()

	ix, ok := tmi.IntIndexes[colName]
	if ok {
		keys, intErr = ix.Get(val)
	}

	tmi.mx.RUnlock()

	if !ok {
		return keys, errors.New("GetListByIndexInt: index " + colName + " not found"), intErr
	}
	return keys, parErr, intErr

}

// GetListByIndexInt list items by col vals limit by limitStop
func (tmi *TableMapInt) GetListByIndexInt(colName string, vals []int64, limitStop int) (itms []ItemInt, parErr error, intErr error) {

	var itmT []ItemInt
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
func (tmi *TableMapInt) GetKeysByIndexString(colName string, val string, limitStop int) (keys []int64, parErr error, intErr error) {
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
func (tmi *TableMapInt) GetListByIndexString(colName string, vals []string, limitStop int) (itms []ItemInt, parErr error, intErr error) {

	var itmT []ItemInt
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
func (tmi *TableMapInt) Set(itm ItemInt) (itmOut ItemInt, isNew bool, isRvEqual bool, err error) {

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

	if len(tmi.TableStruct.NamesIndexInt) > 0 || len(tmi.TableStruct.NamesIndexString) > 0 {

		iis, _, err := ItemIntMake(itm.Data, tmi.ItmStruct)
		if err != nil {
			return itmOut, isNew, isRvEqual, err
		}
		var iisOld ItemIntStruct
		if okOld {
			iisOld, _, err = ItemIntMake(itmOld.Data, tmi.ItmStruct)
			if err != nil {
				return itmOut, isNew, isRvEqual, err
			}
		}

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
					vOO, okOO = iisOld.FieldsInt[k]
					vLO, okLO = iisOld.FieldsIntA[k]
				}
				vON, okON = iis.FieldsInt[k]
				vLN, okLN = iis.FieldsIntA[k]

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
					vOO, okOO = iisOld.FieldsString[k]
					vLO, okLO = iisOld.FieldsStringA[k]
				}
				vON, okON = iis.FieldsString[k]
				vLN, okLN = iis.FieldsStringA[k]

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

	}

	tmi.needSave = true
	tmi.mx.Unlock()

	return itm, true, false, nil
}

// Flush - flush on disk
func (tmi *TableMapInt) Flush() error {

	tmi.mxF.Lock()
	defer tmi.mxF.Unlock()
	if tmi.needSave {
		err := MkDirIfNotExists(tmi.storageName, 0760)
		if err != nil {
			return ErrorNew("TableMapIntFlush mkdir "+tmi.storageName, err)
		}

		tmi.mx.Lock()
		tmi.needSave = false
		tmi.mx.Unlock()

		tmi.mx.RLock()
		b, err := json.Marshal(tmi.TableStruct)
		tmi.mx.RUnlock()
		if err != nil {
			tmi.needSave = true
			return ErrorNew("TableMapIntFlush json TableStruct marshal "+tmi.storageName, err)
		}

		err = FileReplace(tmi.storageName+"struct.json", b, 0660)

		if err != nil {
			tmi.needSave = true
			return ErrorNew("TableMapIntFlush json TableStruct write file "+tmi.storageName+"struct.json", err)
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
					return ErrorNew("TableMapIntFlush json Index marshal "+iName, err)
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
					return ErrorNew("TableMapIntFlush json Index marshal "+iName, err)
				}

			}
		}

		tmi.mx.RLock()
		b, err = json.Marshal(tmi.Data)
		tmi.mx.RUnlock()

		if err != nil {
			tmi.needSave = true
			return ErrorNew("TableMapIntFlush json Data marshal "+tmi.storageName, err)
		}

		err = FileReplace(tmi.storageName+"data.json", b, 0660)

		if err != nil {
			tmi.needSave = true
			return ErrorNew("TableMapIntFlush json Data write file "+tmi.storageName+"data.json", err)
		}
	}
	return nil
}

// PathSet - set path for table
func (tmi *TableMapInt) PathSet(path string) {
	tmi.storageName = path
}

// StorageExists - exists struct file
func (tmi *TableMapInt) StorageExists() (bool, error) {

	tmi.mxF.Lock()
	defer tmi.mxF.Unlock()
	tmi.mx.Lock()
	defer tmi.mx.Unlock()

	// load struct

	_, e, err := FileLoad(tmi.storageName + "struct.json")

	if err != nil {
		return false, ErrorNew("TableMapIntLoad load file "+tmi.storageName+"struct.json", err)
	}

	return e, nil
}

// Load - load from file
func (tmi *TableMapInt) Load() error {

	tmi.mxF.Lock()
	defer tmi.mxF.Unlock()
	tmi.mx.Lock()
	defer tmi.mx.Unlock()

	// load struct

	d, e, err := FileLoad(tmi.storageName + "struct.json")

	if err != nil {
		return ErrorNew("TableMapIntLoad load file "+tmi.storageName+"struct.json", err)
	}

	if !e {
		return errors.New("TableMapIntLoad Not Found file " + tmi.storageName + "struct.json")
	}

	err = json.Unmarshal(d, &tmi.TableStruct)
	if err != nil {
		return ErrorNew("TableMapIntLoad unmarshal "+tmi.storageName+"struct.json", err)
	}

	// load data

	d, e, err = FileLoad(tmi.storageName + "data.json")

	if err != nil {
		return ErrorNew("TableMapIntLoad load file "+tmi.storageName+"data.json", err)
	}

	if !e {
		return errors.New("TableMapIntLoad Not Found file " + tmi.storageName + "data.json")
	}

	err = json.Unmarshal(d, &tmi.Data)
	if err != nil {
		return ErrorNew("TableMapIntLoad unmarshal "+tmi.storageName+"data.json", err)
	}

	{
		tmi.IntIndexes = make(map[string]*IndexIntInt)
		itype := "ii"
		for k, v := range tmi.TableStruct.NamesIndexInt {
			iName := tmi.storageName + "i_" + itype + "_" + v + ".json"

			tmi.IntIndexes[k] = CreateIndexIntInt()
			err = tmi.IntIndexes[k].Load(iName)
			if err != nil {
				return ErrorNew("TableMapIntLoad load index file "+iName, err)
			}
		}
	}

	{
		tmi.StringIndexes = make(map[string]*IndexIntString)
		itype := "ss"
		for k, v := range tmi.TableStruct.NamesIndexString {
			iName := tmi.storageName + "i_" + itype + "_" + v + ".json"

			tmi.StringIndexes[k] = CreateIndexIntString()
			err = tmi.StringIndexes[k].Load(iName)
			if err != nil {
				return ErrorNew("TableMapIntLoad load index file "+iName, err)
			}
		}
	}

	tmi.needSave = false

	return nil
}
