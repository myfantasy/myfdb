package istorage

import (
	"encoding/json"
	"sync"

	"github.com/myfantasy/mdp"
	"github.com/myfantasy/myfdb/ftools"
)

// IntMapStorage simple map based storage
type IntMapStorage struct {
	StoragePlace string `json:"-"`
	NeedSave     bool   `json:"-"`

	Data map[int64]mdp.ItemInt `json:"data"`

	MinVal int64 `json:"min_val,omitempty"`
	MaxVal int64 `json:"max_val,omitempty"`

	mx  sync.RWMutex
	mxF sync.Mutex
}

// IntMapStorageCreate create IntMapStorage
func IntMapStorageCreate(storagePlace string) *IntMapStorage {
	return &IntMapStorage{
		Data:         make(map[int64]mdp.ItemInt),
		StoragePlace: storagePlace,
		NeedSave:     true,
	}
}

// LoadFromDisc load data from LoadFromDisc
func (s *IntMapStorage) LoadFromDisc() *mdp.Error {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.mxF.Lock()
	defer s.mxF.Unlock()

	dpath := s.StoragePlace + "data.json"
	d, e, err := ftools.FileLoad(dpath)

	if err != nil {
		return mdp.ErrorNew("IntMapStorage load file "+dpath+"", err)
	}

	if !e {
		return mdp.ErrorS("IntMapStorage Not Found file " + dpath + "")
	}

	err = json.Unmarshal(d, s)
	if err != nil {
		return mdp.ErrorNew("IntMapStorage unmarshal "+dpath+"", err)
	}

	s.NeedSave = false

	return nil
}

// FlushToDisc flush data to LoadFromDisc
func (s *IntMapStorage) FlushToDisc() *mdp.Error {
	s.mx.RLock()
	defer s.mx.RUnlock()
	if !s.NeedSave {
		return nil
	}

	s.mxF.Lock()
	defer s.mxF.Unlock()
	dpath := s.StoragePlace + "data.json"
	b, err := json.Marshal(s)

	if err != nil {
		return mdp.ErrorNew("IntMapStorage write file "+dpath+" fail Marshal JSON ", err)
	}

	err = ftools.FileReplace(dpath, b, 0660)

	if err != nil {
		return mdp.ErrorNew("IntMapStorage json write file "+dpath+"", err)
	}

	s.NeedSave = false

	return nil
}

// GetItem get Item
func (s *IntMapStorage) GetItem(key int64) (r mdp.ItemInt, ok bool, paramsErr *mdp.Error, internalErr *mdp.Error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	r, ok = s.Data[key]

	return r, ok, nil, nil
}

// SetItem set item
func (s *IntMapStorage) SetItem(i mdp.ItemInt) (r mdp.ItemInt, paramsErr *mdp.Error, internalErr *mdp.Error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	if i.IsRemoved {
		delete(s.Data, i.Key)
	} else {
		s.Data[i.Key] = i
	}
	s.NeedSave = true

	return i, nil, nil
}
