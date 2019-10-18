package main

import (
	"encoding/json"
)

// QueryGetIntItem set int item
type QueryGetIntItem struct {
	TableName      string  `json:"tbl_name"`
	Key            *int64  `json:"key"`
	Keys           []int64 `json:"keys"`
	Limit          int     `json:"limit"`
	Short          bool    `json:"short"`
	All            bool    `json:"all"`
	IncludeDeleted bool    `json:"include_deleted"`
}

// QueryGetStringItem set string item
type QueryGetStringItem struct {
	TableName      string   `json:"tbl_name"`
	Key            *string  `json:"key"`
	Keys           []string `json:"keys"`
	Limit          int      `json:"limit"`
	Short          bool     `json:"short"`
	All            bool     `json:"all"`
	IncludeDeleted bool     `json:"include_deleted"`
}

// GetItemFromTableIntFromJSON get item from int table
func (db *DB) GetItemFromTableIntFromJSON(d []byte) (data []byte, parErr error, intErr error) {

	tblP := QueryGetIntItem{
		Limit: -1,
	}

	err := json.Unmarshal(d, &tblP)
	if err != nil {
		parErr = ErrorNew("fail json ", err)
		ro := make(map[string]interface{})

		if parErr != nil {
			ro["params_err"] = parErr
		}
		if parErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}

	tbl, err := db.TableIntGet(tblP.TableName)
	if err != nil {
		parErr = ErrorNew("Table get: ", err)

		ro := make(map[string]interface{})

		if parErr != nil {
			ro["params_err"] = parErr
		}
		if intErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}

	if tblP.All {
		ro := make(map[string]interface{})
		itms, parErr, intErr := tbl.GetAll(tblP.Limit, tblP.IncludeDeleted)

		ro["items"] = itms
		if parErr != nil {
			ro["params_err"] = parErr
		}
		if intErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}

	if tblP.Key != nil {
		ro := make(map[string]interface{})
		itm, ok, parErr, intErr := tbl.Get(*(tblP.Key))

		if tblP.Short {
			ro["itm"] = itm.Stat()
		} else {
			ro["itm"] = itm
		}

		ro["is_new"] = ok

		if parErr != nil {
			ro["params_err"] = parErr
		}
		if intErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}

	{
		itmsOut := make([]interface{}, 0, 0)
		itms, parErr, intErr := tbl.GetList(tblP.Keys, tblP.Limit)
		for _, v := range itms {
			if tblP.Short {
				itmsOut = append(itmsOut, v.Stat())
			} else {
				itmsOut = append(itmsOut, v)
			}

		}

		ro := make(map[string]interface{})

		ro["itms"] = itmsOut
		if parErr != nil {
			ro["params_err"] = parErr
		}
		if parErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}
}

// GetItemFromTableStringFromJSON get item from string table
func (db *DB) GetItemFromTableStringFromJSON(d []byte) (data []byte, parErr error, intErr error) {

	tblP := QueryGetStringItem{
		Limit: -1,
	}

	err := json.Unmarshal(d, &tblP)
	if err != nil {
		parErr = ErrorNew("fail json ", err)
		ro := make(map[string]interface{})

		if parErr != nil {
			ro["params_err"] = parErr
		}
		if parErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}

	tbl, err := db.TableStringGet(tblP.TableName)
	if err != nil {
		parErr = ErrorNew("Table get: ", err)

		ro := make(map[string]interface{})

		if parErr != nil {
			ro["params_err"] = parErr
		}
		if intErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}

	if tblP.All {
		ro := make(map[string]interface{})
		itms, parErr, intErr := tbl.GetAll(tblP.Limit, tblP.IncludeDeleted)

		ro["items"] = itms
		if parErr != nil {
			ro["params_err"] = parErr
		}
		if intErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}

	if tblP.Key != nil {
		ro := make(map[string]interface{})
		itm, ok, parErr, intErr := tbl.Get(*(tblP.Key))

		if tblP.Short {
			ro["itm"] = itm.Stat()
		} else {
			ro["itm"] = itm
		}

		ro["is_new"] = ok

		if parErr != nil {
			ro["params_err"] = parErr
		}
		if intErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}

	{
		itmsOut := make([]interface{}, 0, 0)
		itms, parErr, intErr := tbl.GetList(tblP.Keys, tblP.Limit)
		for _, v := range itms {
			if tblP.Short {
				itmsOut = append(itmsOut, v.Stat())
			} else {
				itmsOut = append(itmsOut, v)
			}

		}

		ro := make(map[string]interface{})

		ro["itms"] = itmsOut
		if parErr != nil {
			ro["params_err"] = parErr
		}
		if parErr != nil {
			ro["internal_err"] = intErr
		}

		data, err = json.Marshal(ro)
		if err != nil {
			return data, nil, ErrorNew("fail json marshal", err)
		}
		return data, parErr, intErr
	}
}
