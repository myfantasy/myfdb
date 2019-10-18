package main

import (
	"encoding/json"
)

// QuerySetIntItem set int item
type QuerySetIntItem struct {
	TableName string    `json:"tbl_name"`
	Itm       *ItemInt  `json:"itm"`
	Itms      []ItemInt `json:"itms"`
	Full      bool      `json:"full"`
}

// QuerySetStringItem set string item
type QuerySetStringItem struct {
	TableName string       `json:"tbl_name"`
	Itm       *ItemString  `json:"itm"`
	Itms      []ItemString `json:"itms"`
	Full      bool         `json:"full"`
}

// SetItemIntoTableIntFromJSON set item into int table
func (db *DB) SetItemIntoTableIntFromJSON(d []byte) (data []byte, parErr error, intErr error) {

	tblP := QuerySetIntItem{}

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

	if tblP.Itm != nil {
		ro := make(map[string]interface{})
		itmOut, isNew, isRvEqual, parErr, intErr := tbl.Set(*tblP.Itm)

		if tblP.Full {
			ro["out"] = itmOut
		} else {
			ro["out"] = itmOut.Stat()
		}

		ro["is_new"] = isNew
		ro["is_same"] = isRvEqual

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
		itmsOut := make([]map[string]interface{}, 0, len(tblP.Itms))
		for _, v := range tblP.Itms {
			rone := make(map[string]interface{})
			itmOut, isNew, isRvEqual, parE, intE := tbl.Set(v)

			if tblP.Full {
				rone["out"] = itmOut
			} else {
				rone["out"] = itmOut.Stat()
			}

			rone["is_new"] = isNew
			rone["is_same"] = isRvEqual

			if parE != nil {
				rone["params_err"] = parE
			}
			if intE != nil {
				rone["internal_err"] = intE
			}

			parErr = parE
			intErr = intE

			itmsOut = append(itmsOut, rone)

			if parErr != nil || intErr != nil {
				break
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

// SetItemIntoTableStringFromJSON set item into String table
func (db *DB) SetItemIntoTableStringFromJSON(d []byte) (data []byte, parErr error, intErr error) {

	tblP := QuerySetStringItem{}

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

	if tblP.Itm != nil {
		ro := make(map[string]interface{})
		itmOut, isNew, isRvEqual, parErr, intErr := tbl.Set(*tblP.Itm)

		if tblP.Full {
			ro["out"] = itmOut
		} else {
			ro["out"] = itmOut.Stat()
		}

		ro["is_new"] = isNew
		ro["is_same"] = isRvEqual

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
		itmsOut := make([]map[string]interface{}, 0, len(tblP.Itms))
		for _, v := range tblP.Itms {
			rone := make(map[string]interface{})
			itmOut, isNew, isRvEqual, parE, intE := tbl.Set(v)

			if tblP.Full {
				rone["out"] = itmOut
			} else {
				rone["out"] = itmOut.Stat()
			}

			rone["is_new"] = isNew
			rone["is_same"] = isRvEqual

			if parE != nil {
				rone["params_err"] = parE
			}
			if intE != nil {
				rone["internal_err"] = intE
			}

			parErr = parE
			intErr = intE

			itmsOut = append(itmsOut, rone)

			if parErr != nil || intErr != nil {
				break
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
