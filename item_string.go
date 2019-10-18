package main

// ItemString - one row with key string
type ItemString struct {
	Key           string              `json:"key,omitempty"`
	FieldsInt     map[string]int64    `json:"fi,omitempty"`
	FieldsString  map[string]string   `json:"fs,omitempty"`
	FieldsIntA    map[string][]int64  `json:"fia,omitempty"`
	FieldsStringA map[string][]string `json:"fsa,omitempty"`
	Data          *[]byte             `json:"d,omitempty"`
	Rv            int64               `json:"rv,omitempty"`
	IsRemoved     bool                `json:"rm,omitempty"`
}

// ItemStringStat - one row with key string
type ItemStringStat struct {
	Key       string `json:"key,omitempty"`
	Rv        int64  `json:"rv,omitempty"`
	IsRemoved bool   `json:"rm,omitempty"`
}

// Stat - get stat object
func (i ItemString) Stat() ItemStringStat {
	return ItemStringStat{
		Key:       i.Key,
		Rv:        i.Rv,
		IsRemoved: i.IsRemoved,
	}
}
