package main

import (
	"encoding/json"
)

// QueryGetItemStringIdx string val index
type QueryGetItemStringIdx struct {
	TableName string   `json:"tbl_name"`
	ColName   string   `json:"col_name"`
	Val       *string  `json:"val"`
	Vals      []string `json:"vals"`
	Limit     int      `json:"limit"`
	Short     bool     `json:"short"`
}

// GetItemFromTableIntFromByStringIndexJSON get item from int table by string index
func (db *DB) GetItemFromTableIntFromByStringIndexJSON(d []byte) (data []byte, parErr error, intErr error) {

	tblP := QueryGetItemStringIdx{
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

	if tblP.Val != nil {
		ro := make(map[string]interface{})
		keys, parErr, intErr := tbl.GetKeysIndexString(tblP.ColName, *(tblP.Val), tblP.Limit)

		ro["keys"] = keys

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
		itms, parErr, intErr := tbl.GetIndexString(tblP.ColName, tblP.Vals, tblP.Limit)
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

// GetItemFromTableStringFromByStringIndexJSON get item from string table by string index
func (db *DB) GetItemFromTableStringFromByStringIndexJSON(d []byte) (data []byte, parErr error, intErr error) {

	tblP := QueryGetItemStringIdx{
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

	if tblP.Val != nil {
		ro := make(map[string]interface{})
		keys, parErr, intErr := tbl.GetKeysIndexString(tblP.ColName, *(tblP.Val), tblP.Limit)

		ro["keys"] = keys

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
		itms, parErr, intErr := tbl.GetIndexString(tblP.ColName, tblP.Vals, tblP.Limit)
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
