package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	mfd "github.com/myfantasy/myfdbstorage"
	"github.com/valyala/fasthttp"
)

// SetValueISimple - set value struct
type SetValueISimple struct {
	Key      int64  `json:"key,required"`
	Value    string `json:"value,required"`
	IsBase64 bool   `json:"is_base64,omitempty"`
}

func setIValue(ctx *fasthttp.RequestCtx, method string, url string) {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( Expected url format /get/table/key")
		return

	}

	var sv SetValueISimple

	err := json.Unmarshal(ctx.PostBody(), &sv)
	if err != nil {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( json requared")
		return
	}

	name := parts[2]

	st, ok, err := storage.IGetTable(name)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}
	if !ok {
		ctx.Response.SetStatusCode(404)
		fmt.Fprint(ctx, "TBL not found")
		return
	}

	var bval []byte

	if sv.IsBase64 {
		bval, err = base64.StdEncoding.DecodeString(sv.Value)
		if err != nil {
			ctx.Response.SetStatusCode(400)
			fmt.Fprint(ctx, "Bad request: base64 decode error ", err)
			return
		}
	} else {
		bval = []byte(sv.Value)
	}

	err = st.Set(mfd.IItemSimple{Key: sv.Key, BData: bval})
	if err != nil {
		internalServiceError(ctx, err)
		return
	}

	okHandler(ctx)
}

// SetValueSSimple - set value struct
type SetValueSSimple struct {
	Key      string `json:"key,required"`
	Value    string `json:"value,required"`
	IsBase64 bool   `json:"is_base64,omitempty"`
}

func setSValue(ctx *fasthttp.RequestCtx, method string, url string) {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( Expected url format /get/table/key")
		return

	}

	var sv SetValueSSimple

	err := json.Unmarshal(ctx.PostBody(), &sv)
	if err != nil {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( json requared")
		return
	}

	name := parts[2]

	st, ok, err := storage.SGetTable(name)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}
	if !ok {
		ctx.Response.SetStatusCode(404)
		fmt.Fprint(ctx, "TBL not found")
		return
	}

	var bval []byte

	if sv.IsBase64 {
		bval, err = base64.StdEncoding.DecodeString(sv.Value)
		if err != nil {
			ctx.Response.SetStatusCode(400)
			fmt.Fprint(ctx, "Bad request: base64 decode error ", err)
			return
		}
	} else {
		bval = []byte(sv.Value)
	}

	err = st.Set(mfd.SItemSimple{Key: sv.Key, BData: bval})
	if err != nil {
		internalServiceError(ctx, err)
		return
	}

	okHandler(ctx)
}

// KeyValueISimple - del value struct
type KeyValueISimple struct {
	Key int64 `json:"key,required"`
}

func delIValue(ctx *fasthttp.RequestCtx, method string, url string) {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( Expected url format /get/table/key")
		return

	}

	var sv KeyValueISimple

	err := json.Unmarshal(ctx.PostBody(), &sv)
	if err != nil {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( json requared")
		return
	}

	name := parts[2]

	st, ok, err := storage.IGetTable(name)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}
	if !ok {
		ctx.Response.SetStatusCode(404)
		fmt.Fprint(ctx, "TBL not found")
		return
	}

	ok, err = st.Del(sv.Key)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}

	okAndItemHandler(ctx, ok)
}

// KeyValueSSimple - del value struct
type KeyValueSSimple struct {
	Key string `json:"key,required"`
}

func delSValue(ctx *fasthttp.RequestCtx, method string, url string) {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( Expected url format /get/table/key")
		return

	}

	var sv KeyValueSSimple

	err := json.Unmarshal(ctx.PostBody(), &sv)
	if err != nil {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( json requared")
		return
	}

	name := parts[2]

	st, ok, err := storage.SGetTable(name)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}
	if !ok {
		ctx.Response.SetStatusCode(404)
		fmt.Fprint(ctx, "TBL not found")
		return
	}

	ok, err = st.Del(sv.Key)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}

	okAndItemHandler(ctx, ok)
}

func getIValue(ctx *fasthttp.RequestCtx, method string, url string) {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( Expected url format /get/table/key")
		return

	}

	var sv KeyValueISimple

	err := json.Unmarshal(ctx.PostBody(), &sv)
	if err != nil {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( json requared")
		return
	}

	name := parts[2]

	st, ok, err := storage.IGetTable(name)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}
	if !ok {
		ctx.Response.SetStatusCode(404)
		fmt.Fprint(ctx, "TBL not found")
		return
	}

	itm, ok, err := st.Get(sv.Key)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}
	if !ok {
		ctx.Response.SetStatusCode(404)
		fmt.Fprint(ctx, "ITM not found")
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.SetBody(itm.Data())
}

func getSValue(ctx *fasthttp.RequestCtx, method string, url string) {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( Expected url format /get/table/key")
		return

	}

	var sv KeyValueSSimple

	err := json.Unmarshal(ctx.PostBody(), &sv)
	if err != nil {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( json requared")
		return
	}

	name := parts[2]

	st, ok, err := storage.SGetTable(name)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}
	if !ok {
		ctx.Response.SetStatusCode(404)
		fmt.Fprint(ctx, "TBL not found")
		return
	}

	itm, ok, err := st.Get(sv.Key)
	if err != nil {
		internalServiceError(ctx, err)
		return
	}
	if !ok {
		ctx.Response.SetStatusCode(404)
		fmt.Fprint(ctx, "ITM not found")
		return
	}

	ctx.Response.SetStatusCode(200)
	ctx.SetBody(itm.Data())
}
