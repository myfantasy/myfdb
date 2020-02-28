package storage

import (
	"strconv"

	"github.com/myfantasy/mdp"
	"github.com/myfantasy/myfdb/ftools"
	"github.com/myfantasy/myfdb/generator"
	"github.com/myfantasy/myfdb/istorage"
)

// IntAbstractTable abstract table with int key
type IntAbstractTable struct {
	Type string

	Imlt *istorage.IntMapLocalTable
}

// Flush flush table on storage
func (iat *IntAbstractTable) Flush() (err *mdp.Error) {
	if iat.Type == istorage.IntMapLocalTableTableType && iat.Imlt != nil {
		return iat.Imlt.Flush()
	}
	return mdp.ErrorS("GetStruct: Not implement struct unknown type")
}

// GetStruct get structure
func (iat *IntAbstractTable) GetStruct() (s mdp.TableDefinition, paramsErr *mdp.Error, internalErr *mdp.Error) {
	if iat.Type == istorage.IntMapLocalTableTableType && iat.Imlt != nil {
		s, internalErr = iat.Imlt.GetStruct()
		return s, nil, internalErr
	}
	return s, nil, mdp.ErrorS("GetStruct: Not implement struct unknown type")
}

// GetStructFull get structure
func (iat *IntAbstractTable) GetStructFull() (s mdp.TableDefinition, paramsErr *mdp.Error, internalErr *mdp.Error) {
	if iat.Type == istorage.IntMapLocalTableTableType && iat.Imlt != nil {
		s = iat.Imlt.GetStructFull()
		return s, nil, nil
	}
	return s, nil, mdp.ErrorS("GetStructFull: Not implement struct unknown type")
}

// IntAbstractTableCreate create from structure
func IntAbstractTableCreate(s mdp.TableDefinition, db *DB) (iat *IntAbstractTable, paramsErr *mdp.Error, internalErr *mdp.Error) {
	if s.TableType == istorage.IntMapLocalTableTableType {
		iat = &IntAbstractTable{
			Type: s.TableType,
		}
		s.StoragePlace = db.DefaultSavePath + "int_map/" + s.TableName + "_" + strconv.Itoa(int(generator.RvGet())) + "/"
		ftools.MkDirIfNotExists(s.StoragePlace, 0760)
		if s.FlushTimeout == 0 {
			s.FlushTimeout = db.DBFlushTimeout
		}
		tbl, pErr, iErr := istorage.IntMapLocalTableCreate(s)
		iat.Imlt = tbl
		return iat, pErr, iErr
	}
	return iat, nil, mdp.ErrorS("IntAbstractTableCreate: Not implement struct unknown type")
}

// IntAbstractTableLoad load from structure
func IntAbstractTableLoad(s mdp.TableDefinition) (iat *IntAbstractTable, paramsErr *mdp.Error, internalErr *mdp.Error) {
	if s.TableType == istorage.IntMapLocalTableTableType {
		iat = &IntAbstractTable{
			Type: s.TableType,
		}

		iat.Imlt, paramsErr, internalErr = istorage.IntMapLocalTableLoad(s)

		return iat, paramsErr, internalErr
	}

	// if s.TableType == istorage.IntMapLocalTableTableType {
	// 	iat = &IntAbstractTable{
	// 		TableName: s.TableName,
	// 		Type:      s.TableType,
	// 	}
	// 	tbl, pErr, iErr := istorage.IntMapLocalTableLoad(s)
	// 	iat.Imlt = tbl
	// 	return iat, pErr, iErr
	// }
	return iat, nil, mdp.ErrorS("IntAbstractTableLoad: Not implement struct unknown type")
}

// GetItem get item
func (iat *IntAbstractTable) GetItem(key int64) (s mdp.ItemInt, ok bool, paramsErr *mdp.Error, internalErr *mdp.Error) {

	if iat.Type == istorage.IntMapLocalTableTableType && iat.Imlt != nil {
		s, ok, paramsErr, internalErr = iat.Imlt.GetItem(key)
		return s, ok, paramsErr, internalErr
	}
	return s, false, nil, mdp.ErrorS("GetStructFull: Not implement struct unknown type")
}

// SetItem set item
func (iat *IntAbstractTable) SetItem(i mdp.ItemInt) (s mdp.ItemInt, paramsErr *mdp.Error, internalErr *mdp.Error) {

	if iat.Type == istorage.IntMapLocalTableTableType && iat.Imlt != nil {
		s, paramsErr, internalErr = iat.Imlt.SetItem(i)
		return s, paramsErr, internalErr
	}
	return s, nil, mdp.ErrorS("GetStructFull: Not implement struct unknown type")
}
