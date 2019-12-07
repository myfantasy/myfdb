package main

import (
	"encoding/json"
	"errors"
)

// ItemStruct of item
type ItemStruct struct {
	KeyName    string   `json:"key_name,omitempty"`
	StringVals []string `json:"string_vals,omitempty"`
	IntVals    []string `json:"int_vals,omitempty"`
}

// ItemStructDefault get default value if null
func ItemStructDefault(is *ItemStruct) *ItemStruct {
	if is != nil {
		return is
	}
	return &ItemStruct{KeyName: "key"}
}

// ItemIntMake generate ItemInt from json { "key": , others fields ...  }
func ItemIntMake(r []byte, is *ItemStruct) (itm ItemIntStruct, ok bool, err error) {

	is = ItemStructDefault(is)
	m := make(map[string]json.RawMessage)

	err = json.Unmarshal(r, &m)

	itm.FieldsInt = make(map[string]int64)
	itm.FieldsString = make(map[string]string)
	itm.FieldsIntA = make(map[string][]int64)
	itm.FieldsStringA = make(map[string][]string)

	if err != nil {
		return itm, false, err
	}

	v, ok := m[is.KeyName]

	if !ok {
		return itm, false, errors.New("key not found")
	}

	err = json.Unmarshal(v, &itm.Key)

	if err != nil {
		return itm, false, ErrorNew("key unmurshal fail", err)
	}

	var fi int64
	var fil []int64
	for _, fn := range is.IntVals {
		v, ok := m[fn]
		if ok {
			err = json.Unmarshal(v, &fi)
			if err == nil {
				itm.FieldsInt[fn] = fi
				continue
			}
			err = json.Unmarshal(v, &fil)
			if err == nil {
				itm.FieldsIntA[fn] = fil
				continue
			}

		}
	}

	var fs string
	var fsl []string
	for _, fn := range is.StringVals {
		v, ok := m[fn]
		if ok {
			err = json.Unmarshal(v, &fs)
			if err == nil {
				itm.FieldsString[fn] = fs
				continue
			}
			err = json.Unmarshal(v, &fsl)
			if err == nil {
				itm.FieldsStringA[fn] = fsl
				continue
			}

		}
	}

	return itm, true, err
}

// ItemStringMake generate ItemInt from json { "key": , others fields ...  }
func ItemStringMake(r []byte, is *ItemStruct) (itm ItemStringStruct, ok bool, err error) {

	is = ItemStructDefault(is)
	m := make(map[string]json.RawMessage)

	err = json.Unmarshal(r, &m)

	itm.FieldsInt = make(map[string]int64)
	itm.FieldsString = make(map[string]string)
	itm.FieldsIntA = make(map[string][]int64)
	itm.FieldsStringA = make(map[string][]string)

	if err != nil {
		return itm, false, err
	}

	v, ok := m[is.KeyName]

	if !ok {
		return itm, false, errors.New("key not found")
	}

	err = json.Unmarshal(v, &itm.Key)

	if err != nil {
		return itm, false, ErrorNew("key unmurshal fail", err)
	}

	var fi int64
	var fil []int64
	for _, fn := range is.IntVals {
		v, ok := m[fn]
		if ok {
			err = json.Unmarshal(v, &fi)
			if err == nil {
				itm.FieldsInt[fn] = fi
				continue
			}
			err = json.Unmarshal(v, &fil)
			if err == nil {
				itm.FieldsIntA[fn] = fil
				continue
			}

		}
	}

	var fs string
	var fsl []string
	for _, fn := range is.StringVals {
		v, ok := m[fn]
		if ok {
			err = json.Unmarshal(v, &fs)
			if err == nil {
				itm.FieldsString[fn] = fs
				continue
			}
			err = json.Unmarshal(v, &fsl)
			if err == nil {
				itm.FieldsStringA[fn] = fsl
				continue
			}

		}
	}

	return itm, true, err
}
