package main

// ItemInt - one row with key string
type ItemInt struct {
	Key           int64               `json:"key,omitempty"`
	FieldsInt     map[string]int64    `json:"fi,omitempty"`
	FieldsString  map[string]string   `json:"fs,omitempty"`
	FieldsIntA    map[string][]int64  `json:"fia,omitempty"`
	FieldsStringA map[string][]string `json:"fsa,omitempty"`
	Data          *[]byte             `json:"d,omitempty"`
	Rv            int64               `json:"rv,omitempty"`
	IsRemoved     bool                `json:"rm,omitempty"`
}

// ItemIntStat - one row with key string
type ItemIntStat struct {
	Key       int64 `json:"key,omitempty"`
	Rv        int64 `json:"rv,omitempty"`
	IsRemoved bool  `json:"rm,omitempty"`
}

// Stat - get stat object
func (i ItemInt) Stat() ItemIntStat {
	return ItemIntStat{
		Key:       i.Key,
		Rv:        i.Rv,
		IsRemoved: i.IsRemoved,
	}
}
