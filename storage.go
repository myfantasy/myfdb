package main

import (
	"io/ioutil"

	mfd "github.com/myfantasy/myfdbstorage"
	log "github.com/sirupsen/logrus"
)

var storage *mfd.Storage

func runStorage() error {

	mfd.LogFunc = func(err error) {
		log.Errorln(err)
	}

	ok, err := mfd.FileExists(storageSettings.StructFile)
	if err != nil {
		return mfd.ErrorNew("Fail check struct file "+storageSettings.StructFile, err)
	}

	if !ok {
		storage, err = mfd.StorageLoad([]byte("{}"))
		if err != nil {
			return mfd.ErrorNew("Fail load empty storage", err)
		}
		return nil
	}

	b, err := ioutil.ReadFile(storageSettings.StructFile)
	if err != nil {
		return mfd.ErrorNew("Fail load struct file "+storageSettings.StructFile, err)
	}

	storage, err = mfd.StorageLoad(b)
	if err != nil {
		return mfd.ErrorNew("Fail load storage from file "+storageSettings.StructFile, err)
	}

	return nil
}
