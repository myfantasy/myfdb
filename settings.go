package main

import "path/filepath"

func settingsLoad() {
	apiSettings = &APISettings{
		Addr:                 ":7171",
		OutputInternalErrors: true,
	}

	storageSettings = &StorageSettings{
		StructFile: filepath.FromSlash("data/struct.json"),
		TableSDir:  filepath.FromSlash("data/S/"),
		TableIDir:  filepath.FromSlash("data/I/"),
	}

}

var apiSettings *APISettings

var storageSettings *StorageSettings

// APISettings - struct contains settings for run api
type APISettings struct {
	// Addr - listen and serve address
	Addr string

	// OutputInternalErrors - output in responce internal errors
	OutputInternalErrors bool
}

// StorageSettings - struct contains settings for storage
type StorageSettings struct {
	StructFile string
	TableSDir  string
	TableIDir  string
}
