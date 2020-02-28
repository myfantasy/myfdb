package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/myfantasy/myfdb/storage"

	"github.com/myfantasy/mdp"

	"github.com/valyala/fasthttp"

	log "github.com/sirupsen/logrus"
)

func fastHTTPHandlerParamsError(ctx *fasthttp.RequestCtx, err error) {

	fmt.Fprint(ctx, err)
	ctx.Response.SetStatusCode(400)
}

func fastHTTPHandlerInternalError(ctx *fasthttp.RequestCtx, err error) {

	fmt.Fprint(ctx, err)
	ctx.Response.SetStatusCode(500)
}

func fastHTTPHandlerMdpStructGet(ctx *fasthttp.RequestCtx, s mdp.StructGet) {

	if s.InternalErr != nil {
		ctx.Response.SetStatusCode(500)
	} else if s.ParamsErr != nil {
		ctx.Response.SetStatusCode(400)
	}

	d, e := json.Marshal(s)

	if e != nil {
		log.Fatal("Marshaling struct error: ", e)
	}

	ctx.SetBody(d)

}

func fastHTTPHandlerMdpStructStorageGet(ctx *fasthttp.RequestCtx, s mdp.StructStorageGet) {

	if s.InternalErr != nil {
		ctx.Response.SetStatusCode(500)
	} else if s.ParamsErr != nil {
		ctx.Response.SetStatusCode(400)
	}

	d, e := json.Marshal(s)

	if e != nil {
		log.Fatal("Marshaling struct error: ", e)
	}

	ctx.SetBody(d)

}

func fastHTTPHandlerMdpItemsGet(ctx *fasthttp.RequestCtx, s mdp.ItemsGet) {

	if s.InternalErr != nil {
		ctx.Response.SetStatusCode(500)
	} else if s.ParamsErr != nil {
		ctx.Response.SetStatusCode(400)
	}

	d, e := json.Marshal(s)

	if e != nil {
		log.Fatal("Marshaling struct error: ", e)
	}

	ctx.SetBody(d)

}

func fastHTTPHandlerOk(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Ok")
	ctx.Response.SetStatusCode(200)
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {

	url := string(ctx.URI().Path())
	log.Debugln(url)

	if url == "/"+mdp.URLGet {

		b := ctx.PostBody()
		r := DBItemQueryGet(b)

		fastHTTPHandlerMdpItemsGet(ctx, r)

		log.Debugln(string(b))
		log.Debugln(r)

		return
	}

	if url == "/"+mdp.URLSet {

		b := ctx.PostBody()
		r := DBItemQuerySet(b)

		fastHTTPHandlerMdpItemsGet(ctx, r)

		log.Debugln(string(b))
		log.Debugln(r)

		return
	}

	if url == "/"+mdp.URLStructGet {

		b := ctx.PostBody()
		r := DBStructQueryGet(b)

		fastHTTPHandlerMdpStructGet(ctx, r)

		log.Debugln(string(b))
		log.Debugln(r)

		return
	}

	if url == "/"+mdp.URLStructSet {

		b := ctx.PostBody()
		r := DBStructQuerySet(b)

		fastHTTPHandlerMdpStructGet(ctx, r)

		log.Debugln(string(b))
		log.Debugln(r)

		return
	}

	if url == "/"+mdp.URLStructStorageGet {

		b := ctx.PostBody()
		r := DBStructStorageQueryGet(b)

		fastHTTPHandlerMdpStructStorageGet(ctx, r)

		log.Debugln(string(b))
		log.Debugln(r)

		return
	}

	if url == "/"+mdp.URLStructStorageSet {

		b := ctx.PostBody()
		r := DBStructStorageQuerySet(b)

		fastHTTPHandlerMdpStructStorageGet(ctx, r)

		log.Debugln(string(b))
		log.Debugln(r)

		return
	}

	fmt.Fprint(ctx, "Not found")
	ctx.Response.SetStatusCode(404)
}

func logsOut(err error) {
	log.Errorln(err)
}

var db *storage.DB

func main() {

	initConf()

	log.Debug(config)

	stopSignal := make(chan bool, 1)

	dbO, e := storage.DBLoadFromWriteStruct(config.DBFolder, config.DBFlushTimeout, stopSignal)

	if e != nil {
		log.Fatal(e)
	}

	db = dbO

	log.Debug("Cluster: ", db.ServerName)

	api := &fasthttp.Server{
		Handler: fastHTTPHandler,
	}

	serverErrors := make(chan error, 1)
	go func() {
		ls := ":" + strconv.Itoa(config.Port)
		log.Info("Listen and serve " + ls)
		serverErrors <- api.ListenAndServe(ls)
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("Can`t start server; %v", err)
	case <-stopSignal:
		log.Info("Stop signal recived. Start shutdown...")
		go func() {
			if err := api.Shutdown(); err != nil {
				log.Infof("S: Graceful shutdown did not complete in 5s : %v", err)
			}
		}()
	case <-osSignals:
		log.Info("Start shutdown...")
		go func() {
			if err := api.Shutdown(); err != nil {
				log.Infof("OS: Graceful shutdown did not complete in 5s : %v", err)
			}
		}()
	}

	log.Info("FlushAll")
	db.FlushAll()

	log.Info("Good Bye")
}
