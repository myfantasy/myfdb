package ftools

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/myfantasy/mdp"
)

// FileExists check file on disk
func FileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// FileReplace - Create or replace file (remove .old && .new if exists -> create .new -> move current to .old -> move .new to current -> remove .old)
// 0660 - i use this mode
func FileReplace(path string, data []byte, fmode os.FileMode) error {

	path = filepath.FromSlash(path)

	pathOld := path + ".old"
	pathNew := path + ".new"

	e, err := FileExists(pathOld)
	if err != nil {
		return mdp.ErrorNew("FileReplace Check old file "+pathOld, err)
	}
	if e {
		err = os.Remove(pathOld)
		if err != nil {
			return mdp.ErrorNew("FileReplace Remove old file (1) "+pathOld, err)
		}
	}

	e, err = FileExists(pathNew)
	if err != nil {
		return mdp.ErrorNew("FileReplace Check new file "+pathNew, err)
	}
	if e {
		err = os.Remove(pathNew)
		if err != nil {
			return mdp.ErrorNew("FileReplace Remove new file "+pathNew, err)
		}
	}

	err = ioutil.WriteFile(pathNew, data, fmode)
	if err != nil {
		return mdp.ErrorNew("FileReplace Write new file "+pathNew, err)
	}

	e, err = FileExists(path)
	if err != nil {
		return mdp.ErrorNew("FileReplace Check file "+path, err)
	}
	if e {
		err = os.Rename(path, pathOld)
		if err != nil {
			return mdp.ErrorNew("FileReplace move "+path+" to new file "+pathOld, err)
		}
	}

	err = os.Rename(pathNew, path)
	if err != nil {
		return mdp.ErrorNew("FileReplace move new file "+pathNew+" to "+path, err)
	}

	e, err = FileExists(pathOld)
	if err != nil {
		return mdp.ErrorNew("FileReplace Check old file "+pathOld, err)
	}
	if e {
		err = os.Remove(pathOld)
		if err != nil {
			return mdp.ErrorNew("FileReplace Remove old file (2) "+pathOld, err)
		}
	}

	return nil
}

//FileLoad load file path -> .new -> .old
func FileLoad(path string) (data []byte, e bool, err error) {

	path = filepath.FromSlash(path)

	pathOld := path + ".old"
	pathNew := path + ".new"

	e, err = FileExists(path)
	if err != nil {
		return data, e, mdp.ErrorNew("FileLoad Check file "+path, err)
	}
	if e {
		data, err = ioutil.ReadFile(path)
		if err != nil {
			return data, e, mdp.ErrorNew("FileLoad file "+path, err)
		}
		return data, e, nil
	}

	e, err = FileExists(pathNew)
	if err != nil {
		return data, e, mdp.ErrorNew("FileLoad Check new file "+pathNew, err)
	}
	if e {
		data, err = ioutil.ReadFile(pathNew)
		if err != nil {
			return data, e, mdp.ErrorNew("FileLoad new file "+pathNew, err)
		}
		return data, e, nil
	}

	e, err = FileExists(pathOld)
	if err != nil {
		return data, e, mdp.ErrorNew("FileLoad Check old file "+pathOld, err)
	}
	if e {
		data, err = ioutil.ReadFile(pathOld)
		if err != nil {
			return data, e, mdp.ErrorNew("FileLoad old file "+pathOld, err)
		}
		return data, e, nil
	}

	return data, false, nil
}

// MkDirIfNotExists make all directory if not exists
// 0760 - i use this mode
func MkDirIfNotExists(path string, fmode os.FileMode) (err error) {
	path = filepath.FromSlash(path)

	ok, err := FileExists(path)
	if err != nil {
		return mdp.ErrorNew("MkDirIfNotExists Check directory "+path, err)
	}
	if !ok {
		err = os.MkdirAll(path, fmode)
		if err != nil {
			return mdp.ErrorNew("MkDirIfNotExists Mkdir file "+path, err)
		}
	}

	return nil
}
