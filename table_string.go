package main

import (
	"errors"
	"time"
)

// TableString - table for items string key
type TableString struct {
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	FlushTimeout time.Duration `json:"flush_timeout"`
	Folder       string        `json:"data_folder"`

	// Not save
	BaseFolder string          `json:"-"`
	TableMap   *TableMapString `json:"-"`

	// Tech
	lastFlush time.Time
}

// FullPath get full path to stor
func (ti *TableString) FullPath() string {
	return ti.BaseFolder + ti.Folder
}

// CreateTableString create table
func CreateTableString(tblType string, name string, baseFolder string, path string, flushTimeout time.Duration, eo ErrorsOut) (ti *TableString, err error) {
	ti = &TableString{
		Name:         name,
		Type:         tblType,
		Folder:       path,
		FlushTimeout: flushTimeout,

		BaseFolder: baseFolder,
	}

	if tblType == TableMapStringName {
		ti.TableMap = CreateTableMapString(ti.FullPath())

		sfok, err := ti.TableMap.StorageExists()
		if err != nil {
			return nil, err
		}
		if sfok {
			err = ti.TableMap.Load()
		} else {
			err = ti.TableMap.Flush()
		}
		if err != nil {
			return nil, err
		}

		go func() {
			for {
				e := ti.TableMap.Flush()
				if e != nil && eo != nil {
					eo(e)
				}
				time.Sleep(ti.FlushTimeout)
			}
		}()

	} else {
		return nil, errors.New("CreateTable: table type " + tblType + " not support")
	}

	return ti, nil
}

// Init table for start
func (ti *TableString) Init(eo ErrorsOut) error {

	if ti.Type == TableMapStringName {
		ti.TableMap = CreateTableMapString(ti.FullPath())
		err := ti.TableMap.Load()
		if err != nil {
			return err
		}

		go func() {
			for {
				e := ti.TableMap.Flush()
				if e != nil && eo != nil {
					eo(e)
				}
				time.Sleep(ti.FlushTimeout)
			}
		}()

	} else {
		return errors.New("Init: table type " + ti.Type + " not support")
	}

	return nil
}

// IndexIntCreate create int index
func (ti *TableString) IndexIntCreate(colName string) (parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		parErr = ti.TableMap.IndexIntAdd(colName, "ix_i_"+colName)
		return parErr, nil
	}

	return errors.New("IndexIntCreate: table type " + ti.Type + " not support"), nil

}

// IndexStringCreate create int index
func (ti *TableString) IndexStringCreate(colName string) (parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		parErr = ti.TableMap.IndexStringAdd(colName, "ix_s_"+colName)
		return parErr, nil
	}

	return errors.New("IndexStringCreate: table type " + ti.Type + " not support"), nil

}

// Set set value
func (ti *TableString) Set(itm ItemString) (itmOut ItemString, isNew bool, isRvEqual bool, parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		itmOut, isNew, isRvEqual, intErr = ti.TableMap.Set(itm)
		return itmOut, isNew, isRvEqual, parErr, intErr
	}

	return itmOut, isNew, isRvEqual, errors.New("Set: table type " + ti.Type + " not support"), nil

}

// Get get value
func (ti *TableString) Get(key string) (itm ItemString, ok bool, parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		itm, ok, intErr = ti.TableMap.Get(key)
		return itm, ok, parErr, intErr
	}

	return itm, ok, errors.New("Get: table type " + ti.Type + " not support"), nil

}

// GetAll get value
func (ti *TableString) GetAll(limitStop int, includeDeleted bool) (itms []ItemString, parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		itms, parErr, intErr = ti.TableMap.GetAll(limitStop, includeDeleted)
		return itms, parErr, intErr
	}

	return itms, errors.New("GetAll: table type " + ti.Type + " not support"), nil

}

// GetList get values List
func (ti *TableString) GetList(keys []string, limitStop int) (itms []ItemString, parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		itms, intErr = ti.TableMap.GetList(keys, limitStop)
		return itms, parErr, intErr
	}

	return itms, errors.New("GetList: table type " + ti.Type + " not support"), nil

}

// GetKeysIndexInt get values List
func (ti *TableString) GetKeysIndexInt(colName string, val int64, limitStop int) (keys []string, parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		keys, parErr, intErr = ti.TableMap.GetKeysByIndexInt(colName, val, limitStop)
		return keys, parErr, intErr
	}

	return keys, errors.New("GetKeysIndexInt: table type " + ti.Type + " not support"), nil

}

// GetIndexInt get values List
func (ti *TableString) GetIndexInt(colName string, vals []int64, limitStop int) (itms []ItemString, parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		itms, parErr, intErr = ti.TableMap.GetListByIndexInt(colName, vals, limitStop)
		return itms, parErr, intErr
	}

	return itms, errors.New("GetIndexInt: table type " + ti.Type + " not support"), nil

}

// GetKeysIndexString get values List
func (ti *TableString) GetKeysIndexString(colName string, val string, limitStop int) (keys []string, parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		keys, parErr, intErr = ti.TableMap.GetKeysByIndexString(colName, val, limitStop)
		return keys, parErr, intErr
	}

	return keys, errors.New("GetKeysIndexString: table type " + ti.Type + " not support"), nil

}

// GetIndexString get values List
func (ti *TableString) GetIndexString(colName string, vals []string, limitStop int) (itms []ItemString, parErr error, intErr error) {
	if ti.Type == TableMapStringName {
		itms, parErr, intErr = ti.TableMap.GetListByIndexString(colName, vals, limitStop)
		return itms, parErr, intErr
	}

	return itms, errors.New("GetIndexString: table type " + ti.Type + " not support"), nil

}
