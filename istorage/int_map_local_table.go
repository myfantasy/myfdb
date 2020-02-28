package istorage

import (
	"sync"
	"time"

	"github.com/myfantasy/mdp"
	"github.com/myfantasy/myfdb/logger"
)

// IntMapLocalTableTableType table type of int_map_local_table
const IntMapLocalTableTableType = "int_map_local_table"

// IntMapLocalTable Local table
type IntMapLocalTable struct {
	TableDefinition mdp.TableDefinition
	Storage         *IntMapStorage

	mx sync.RWMutex
}

// Flush flush table on storage
func (imt *IntMapLocalTable) Flush() (err *mdp.Error) {
	if imt.Storage != nil {
		if imt.Storage.StoragePlace != "" {
			return imt.Storage.FlushToDisc()
		}
	}
	return mdp.ErrorS("IntMapLocalTable.Flush: Storage is nil")
}

// IntMapLocalTableCreate load table from struct
func IntMapLocalTableCreate(s mdp.TableDefinition) (tbl *IntMapLocalTable, paramsErr *mdp.Error, internalErr *mdp.Error) {

	tbl = &IntMapLocalTable{}

	tbl.TableDefinition = s

	tbl.Storage = IntMapStorageCreate(tbl.TableDefinition.StoragePlace)
	go tbl.DoSaveIteration()

	return tbl, nil, nil
}

// DoSaveIteration iteration for flush to disc
func (imt *IntMapLocalTable) DoSaveIteration() {

	for !imt.TableDefinition.IsDeleted {
		if imt.TableDefinition.FlushTimeout > 0 {
			time.Sleep(imt.TableDefinition.FlushTimeout)
		} else {
			time.Sleep(time.Second)
		}
		imt.mx.RLock()
		if imt.Storage.StoragePlace != "" {
			err := imt.Storage.FlushToDisc()
			if err != nil {
				logger.InternalProcessError(mdp.ErrorNew("IntMapLocalTable Write on disc fail for table "+imt.TableDefinition.TableName+" ("+imt.TableDefinition.UniqueID+")", err))
			}
		}
		imt.mx.RUnlock()
	}
}

// IntMapLocalTableLoad load table from struct
func IntMapLocalTableLoad(s mdp.TableDefinition) (tbl *IntMapLocalTable, paramsErr *mdp.Error, internalErr *mdp.Error) {

	tbl = &IntMapLocalTable{TableDefinition: s}

	if tbl.TableDefinition.StoragePlace != "" {
		tbl.Storage = &IntMapStorage{StoragePlace: tbl.TableDefinition.StoragePlace}
		err := tbl.Storage.LoadFromDisc()
		if err != nil {
			return tbl, nil, mdp.ErrorNew("IntMapLocalTableLoad fail load from disc for table "+s.TableName+" ("+s.UniqueID+")", err)
		}
	} else {
		tbl.Storage = IntMapStorageCreate(tbl.TableDefinition.StoragePlace)
	}

	go tbl.DoSaveIteration()

	return tbl, nil, nil
}

// GetStruct get table struct
func (imt *IntMapLocalTable) GetStruct() (s mdp.TableDefinition, err *mdp.Error) {

	imt.mx.RLock()
	defer imt.mx.RUnlock()

	if imt.Storage == nil {
		return s, mdp.ErrorS("IntMapLocalTable.GetStruct: Storage not init")
	}

	s = imt.TableDefinition.ClearLocalInfo()
	return s, nil

}

// GetStructFull get table full struct
func (imt *IntMapLocalTable) GetStructFull() (s mdp.TableDefinition) {

	imt.mx.RLock()
	defer imt.mx.RUnlock()

	s = imt.TableDefinition
	return s

}

// GetItem get item
func (imt *IntMapLocalTable) GetItem(key int64) (s mdp.ItemInt, ok bool, paramsErr *mdp.Error, internalErr *mdp.Error) {
	imt.mx.RLock()
	defer imt.mx.RUnlock()

	if imt.Storage == nil {
		return s, false, nil, mdp.ErrorS("IntMapLocalTable.GetItem: Storage not init")
	}

	s, ok, paramsErr, internalErr = imt.Storage.GetItem(key)

	return s, ok, paramsErr, internalErr
}

// SetItem set item
func (imt *IntMapLocalTable) SetItem(s mdp.ItemInt) (r mdp.ItemInt, paramsErr *mdp.Error, internalErr *mdp.Error) {
	imt.mx.RLock()
	defer imt.mx.RUnlock()

	if imt.Storage == nil {
		return r, nil, mdp.ErrorS("IntMapLocalTable.SetItem: Storage not init")
	}

	r, paramsErr, internalErr = imt.Storage.SetItem(s)

	return r, paramsErr, internalErr
}

// import (
// 	"encoding/json"
// 	"time"

// 	"github.com/myfantasy/mdp"
// 	"github.com/myfantasy/myfdb/logger"
// )

// // IntMapLocalTableTableType table type of int_map_local_table
// const IntMapLocalTableTableType = "int_map_local_table"

// // IntMapLocalTablePublicParams Local table
// type IntMapLocalTablePublicParams struct {
// 	FlushTimeout time.Duration `json:"flush_duration,omitempty"`
// }

// // IntMapLocalTable Local table
// type IntMapLocalTable struct {
// 	UniqueID     string
// 	TableName    string
// 	StoragePlace string

// 	PublicParams IntMapLocalTablePublicParams

// 	MetaData json.RawMessage

// 	ItemsStruct *mdp.ItemStruct
// 	Version     int64
// 	IsDeleted   bool

// 	Storage *IntMapStorage
// }

// // GetStruct get table full struct
// func (tbl *IntMapLocalTable) GetStruct() (s mdp.TableDefinition) {

// 	pps, _ := json.Marshal(tbl.PublicParams)

// 	s.UniqueID = tbl.UniqueID
// 	s.TableName = tbl.TableName
// 	s.TableType = IntMapLocalTableTableType
// 	s.KeyType = mdp.KeyTypeInt
// 	s.StoragePlace = tbl.StoragePlace
// 	s.PublicParams = pps
// 	s.MetaData = tbl.MetaData
// 	s.ItemsStruct = tbl.ItemsStruct
// 	s.Version = tbl.Version
// 	s.IsDeleted = tbl.IsDeleted

// 	return s
// }

// // IntMapLocalTableLoad load table from struct
// func IntMapLocalTableLoad(s mdp.TableDefinition) (tbl *IntMapLocalTable, paramsErr *mdp.Error, internalErr *mdp.Error) {

// 	tbl = &IntMapLocalTable{
// 		UniqueID:     s.UniqueID,
// 		TableName:    s.TableName,
// 		StoragePlace: s.StoragePlace,
// 		MetaData:     s.MetaData,
// 		ItemsStruct:  s.ItemsStruct,
// 		Version:      s.Version,
// 		IsDeleted:    s.IsDeleted,
// 	}

// 	err := json.Unmarshal(s.PublicParams, &tbl.PublicParams)
// 	if err != nil {
// 		return tbl, nil, mdp.ErrorNew("IntMapLocalTableLoad fail unmarshal for table "+s.TableName+" ("+s.UniqueID+")", err)
// 	}

// 	tbl.Storage = &IntMapStorage{StoragePlace: tbl.StoragePlace}

// 	if tbl.Storage.StoragePlace != "" {
// 		err = tbl.Storage.LoadFromDisc()
// 		return tbl, nil, mdp.ErrorNew("IntMapLocalTableLoad fail load from disc for table "+s.TableName+" ("+s.UniqueID+")", err)
// 	}

// 	go func() {
// 		for !tbl.IsDeleted {
// 			if tbl.PublicParams.FlushTimeout > 0 {
// 				time.Sleep(tbl.PublicParams.FlushTimeout)
// 			} else {
// 				time.Sleep(time.Second)
// 			}

// 			if tbl.Storage.StoragePlace != "" {
// 				err := tbl.Storage.FlushToDisc()
// 				logger.InternalProcessError(mdp.ErrorNew("IntMapLocalTable Write fail load from disc for table "+s.TableName+" ("+s.UniqueID+")", err))
// 			}
// 		}
// 	}()

// 	return tbl, nil, nil
// }

// // IntMapLocalTableCreate load table from struct
// func IntMapLocalTableCreate(s mdp.TableDefinition) (tbl *IntMapLocalTable, paramsErr *mdp.Error, internalErr *mdp.Error) {

// 	tbl = &IntMapLocalTable{
// 		UniqueID:     s.UniqueID,
// 		TableName:    s.TableName,
// 		StoragePlace: s.StoragePlace,
// 		MetaData:     s.MetaData,
// 		ItemsStruct:  s.ItemsStruct,
// 		Version:      s.Version,
// 		IsDeleted:    s.IsDeleted,
// 	}

// 	err := json.Unmarshal(s.PublicParams, &tbl.PublicParams)
// 	if err != nil {
// 		return tbl, nil, mdp.ErrorNew("IntMapLocalTableLoad fail unmarshal for table "+s.TableName+" ("+s.UniqueID+")", err)
// 	}

// 	tbl.Storage = &IntMapStorage{StoragePlace: tbl.StoragePlace}

// 	if tbl.Storage.StoragePlace != "" {
// 		err = tbl.Storage.LoadFromDisc()
// 		return tbl, nil, mdp.ErrorNew("IntMapLocalTableLoad fail load from disc for table "+s.TableName+" ("+s.UniqueID+")", err)
// 	}

// 	go func() {
// 		for !tbl.IsDeleted {
// 			if tbl.PublicParams.FlushTimeout > 0 {
// 				time.Sleep(tbl.PublicParams.FlushTimeout)
// 			} else {
// 				time.Sleep(time.Second)
// 			}

// 			if tbl.Storage.StoragePlace != "" {
// 				err := tbl.Storage.FlushToDisc()
// 				logger.InternalProcessError(mdp.ErrorNew("IntMapLocalTable Write fail load from disc for table "+s.TableName+" ("+s.UniqueID+")", err))
// 			}
// 		}
// 	}()

// 	return tbl, nil, nil
// }
