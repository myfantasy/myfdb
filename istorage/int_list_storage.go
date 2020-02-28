package istorage

// import (
// 	"sort"
// 	"strconv"
// 	"sync"

// 	"github.com/myfantasy/mdp"
// )

// // IntListBlock data block info of ItemInt
// type IntListBlock struct {
// 	ID int64 `json:"block_id,omitempty"`

// 	Count  int   `json:"cnt,omitempty"`
// 	MinVal int64 `json:"min_val,omitempty"`
// 	MaxVal int64 `json:"max_val,omitempty"`

// 	Items []mdp.ItemInt `json:"-"`

// 	mx sync.RWMutex
// }

// // Less ilb less then ilb2
// func (ilb *IntListBlock) Less(ilb2 *IntListBlock) bool {

// 	return ilb != nil && ilb2 != nil && ilb.MinVal < ilb2.MinVal

// }

// // IntListBlockCreate create new IntListBlock
// func IntListBlockCreate(itm mdp.ItemInt) *IntListBlock {
// 	ilb := &IntListBlock{
// 		ID:     RvGet(),
// 		Count:  1,
// 		MinVal: itm.Key,
// 		MaxVal: itm.Key,
// 		Items:  []mdp.ItemInt{itm},
// 	}

// 	return ilb
// }

// // GetInternal get item
// func (ilb *IntListBlock) GetInternal(key int64) (ok bool, itmOut mdp.ItemInt, idx int) {
// 	if ilb.MinVal > key || ilb.MaxVal < key || ilb.Count == 0 {
// 		return false, itmOut, idx
// 	}

// 	if ilb.Count == 1 {
// 		if ilb.Items[0].Key == key {
// 			return true, ilb.Items[0], 0
// 		}
// 		return false, itmOut, idx
// 	}

// 	mediana := ilb.Count / 2
// 	append := 0
// 	for {
// 		if ilb.Items[mediana].Key == key {
// 			return true, ilb.Items[mediana], mediana
// 		}

// 		if ilb.Items[mediana].Key < key {
// 			if mediana == ilb.Count-1 {
// 				return false, itmOut, mediana
// 			}
// 			append = (ilb.Count - mediana) / 2
// 			if append == 0 {
// 				append = 1
// 			}
// 			mediana += append
// 		}

// 		if ilb.Items[mediana].Key > key {
// 			if mediana == 0 {
// 				return false, itmOut, mediana
// 			}
// 			append = (mediana) / 2
// 			if append == 0 {
// 				append = 1
// 			}
// 			mediana -= append
// 		}
// 	}

// }

// // AddIntoList add into list
// func (ilb *IntListBlock) AddIntoList(itm mdp.ItemInt, idx int) {
// 	ilb.Items = append(ilb.Items, itm)
// 	sort.Slice(ilb.Items, func(i, j int) bool { return ilb.Items[i].Less(ilb.Items[j]) })
// }

// // Add add item
// func (ilb *IntListBlock) Add(itm mdp.ItemInt) (isUpd bool, isEqual bool, itmOut mdp.ItemInt) {
// 	ilb.mx.Lock()
// 	defer ilb.mx.Unlock()

// 	ok, itmOut, idx := ilb.GetInternal(itm.Key)
// 	if !ok {
// 		ilb.AddIntoList(itm, idx)
// 		return true, false, itm
// 	}

// 	if itmOut.Rv == itm.Rv {
// 		return false, true, itmOut
// 	}

// 	if itmOut.Rv > itm.Rv {
// 		return false, false, itmOut
// 	}

// 	ilb.Items[idx] = itm
// 	return true, false, itm

// }

// // IntListStorage data storage of ItemInt
// type IntListStorage struct {
// 	Count  int   `json:"cnt,omitempty"`
// 	MinVal int64 `json:"min_val,omitempty"`
// 	MaxVal int64 `json:"max_val,omitempty"`

// 	BlocksIds        []int64           `json:"blocks,omitempty"`
// 	InternalStorages []*IntListStorage `json:"internal_storages,omitempty"`

// 	Blocks map[int64]*IntListBlock `json:"-"`

// 	mx sync.RWMutex
// }

// // GetBlockByOrderIndexInternal get block by index and fatal when not exists or out of range
// func (ils *IntListStorage) GetBlockByOrderIndexInternal(idx int) (blockOut *IntListBlock) {
// 	v, ok := ils.Blocks[ils.BlocksIds[idx]]
// 	if ok {
// 		return v
// 	}

// 	panic(mdp.ErrorS("Error: object must be. IntListStorage.GetBlockByOrderIndexInternal(" + strconv.Itoa(idx) + ")"))
// }

// // GetInternal get block
// func (ils *IntListStorage) GetInternal(key int64) (ok bool, blockOut *IntListBlock, idx int) {
// 	if ils.Count == 0 {
// 		return false, blockOut, idx
// 	}

// 	cnt := len(ils.BlocksIds)

// 	if cnt == 1 {
// 		return true, ils.GetBlockByOrderIndexInternal(0), 0
// 	}

// 	mediana := cnt / 2
// 	append := 0
// 	for {
// 		bl := ils.GetBlockByOrderIndexInternal(mediana)
// 		bl.mx.RLock()

// 		bl.mx.RUnlock()

// 		if ilb.Items[mediana].Key == key {
// 			return true, ilb.Items[mediana], mediana
// 		}

// 		if ilb.Items[mediana].Key < key {
// 			if mediana == ilb.Count-1 {
// 				return false, itmOut, mediana
// 			}
// 			append = (ilb.Count - mediana) / 2
// 			if append == 0 {
// 				append = 1
// 			}
// 			mediana += append
// 		}

// 		if ilb.Items[mediana].Key > key {
// 			if mediana == 0 {
// 				return false, itmOut, mediana
// 			}
// 			append = (mediana) / 2
// 			if append == 0 {
// 				append = 1
// 			}
// 			mediana -= append
// 		}
// 	}

// }

// // AddInternal item
// func (ils *IntListStorage) AddInternal(itm mdp.ItemInt) {

// 	var isFirst bool

// 	if ils.Count == 0 {
// 		isFirst = true
// 	}
// 	if isFirst {
// 		ils.Count = 1
// 		ils.MinVal = itm.Key
// 		ils.MaxVal = itm.Key

// 		ilb := IntListBlockCreate(itm)
// 		ils.BlocksIds = []int64{ilb.ID}
// 		ils.Blocks = map[int64]*IntListBlock{ilb.ID: ilb}
// 		return
// 	}

// }
