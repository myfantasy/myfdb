package main

import (
	"encoding/json"
	"path/filepath"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

var mx sync.RWMutex

func createITable(ctx *fasthttp.RequestCtx, method string, url string) {
	var ts TableStruct

	if err := json.Unmarshal(ctx.PostBody(), &ts); err != nil {
		ctx.Response.SetStatusCode(400)
		return
	}

	mx.Lock()
	defer mx.Unlock()

	params := make(map[string]interface{})
	if ts.Stored {
		params["dump_path"] = filepath.FromSlash(storageSettings.TableIDir + ts.Name + "/")
		params["dump_timeout"] = time.Second
	}

	err := storage.AddTableI(ts.Name, ts.TableType, params)

	if err != nil {
		internalServiceError(ctx, err)
		return
	}

	err = storage.StorageStructFlush(storageSettings.StructFile)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}

	okHandler(ctx)
}

// TableStruct Table create params
type TableStruct struct {
	Name      string `json:"name,required"`
	TableType string `json:"type,required"`
	Stored    bool   `json:"stored,required"`
}
