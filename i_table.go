package main

// Table - data table
type Table interface {
	GetStruct() (name string, t string, s map[string]interface{})
	FlushAndClose() error
}
