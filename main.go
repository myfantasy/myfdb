package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/valyala/fasthttp"

	log "github.com/sirupsen/logrus"
)

func fastHTTPHandlerParamsError(ctx *fasthttp.RequestCtx, err error) {

	fmt.Fprint(ctx, err)
	ctx.Response.SetStatusCode(406)
}

func fastHTTPHandlerInternalError(ctx *fasthttp.RequestCtx, err error) {

	fmt.Fprint(ctx, err)
	ctx.Response.SetStatusCode(503)
}

func fastHTTPHandlerOk(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Ok")
	ctx.Response.SetStatusCode(200)
}

func fastHTTPHandlerBodyJSON(ctx *fasthttp.RequestCtx, b []byte, err error) {

	if err != nil {
		fastHTTPHandlerInternalError(ctx, err)
	} else {
		ctx.SetContentType("application/json")
		ctx.SetBody(b)
		ctx.Response.SetStatusCode(200)
	}
}

func fastHTTPHandlerDblError(ctx *fasthttp.RequestCtx, parErr error, intErr error) {

	if parErr != nil {
		fastHTTPHandlerParamsError(ctx, parErr)
	} else if intErr != nil {
		fastHTTPHandlerInternalError(ctx, intErr)
	} else {
		fastHTTPHandlerOk(ctx)
	}
}

func fastHTTPHandlerDblErrorBodyJSON(ctx *fasthttp.RequestCtx, b []byte, parErr error, intErr error) {

	if intErr != nil {
		if len(b) == 0 {
			fastHTTPHandlerInternalError(ctx, intErr)
		} else {
			ctx.SetContentType("application/json")
			ctx.SetBody(b)
			ctx.Response.SetStatusCode(500)
		}
	} else if parErr != nil {
		if len(b) == 0 {
			fastHTTPHandlerParamsError(ctx, parErr)
		} else {
			ctx.SetContentType("application/json")
			ctx.SetBody(b)
			ctx.Response.SetStatusCode(400)
		}

	} else {
		ctx.SetContentType("application/json")
		ctx.SetBody(b)
		ctx.Response.SetStatusCode(200)
	}
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {

	token := ""
	ctx.Request.Header.VisitAll(func(key, value []byte) {
		if string(key) == "Token" {
			token = string(value)
		}
	})

	if !dataBase.CheckToken(token) {
		fmt.Fprint(ctx, "401 Unauthorized")
		ctx.Response.SetStatusCode(401)
		return
	}

	url := string(ctx.URI().Path())

	if url == "/i/set/" {
		b, parErr, intErr := dataBase.SetItemIntoTableIntFromJSON(ctx.PostBody())
		fastHTTPHandlerDblErrorBodyJSON(ctx, b, parErr, intErr)

		return
	}
	if url == "/s/set/" {
		b, parErr, intErr := dataBase.SetItemIntoTableStringFromJSON(ctx.PostBody())
		fastHTTPHandlerDblErrorBodyJSON(ctx, b, parErr, intErr)

		return
	}

	if url == "/i/get/" {
		b, parErr, intErr := dataBase.GetItemFromTableIntFromJSON(ctx.PostBody())
		fastHTTPHandlerDblErrorBodyJSON(ctx, b, parErr, intErr)

		return
	}
	if url == "/s/get/" {
		b, parErr, intErr := dataBase.GetItemFromTableStringFromJSON(ctx.PostBody())
		fastHTTPHandlerDblErrorBodyJSON(ctx, b, parErr, intErr)

		return
	}

	if url == "/i/iiget/" {
		b, parErr, intErr := dataBase.GetItemFromTableIntFromByIntIndexJSON(ctx.PostBody())
		fastHTTPHandlerDblErrorBodyJSON(ctx, b, parErr, intErr)

		return
	}
	if url == "/s/iiget/" {
		b, parErr, intErr := dataBase.GetItemFromTableStringFromByIntIndexJSON(ctx.PostBody())
		fastHTTPHandlerDblErrorBodyJSON(ctx, b, parErr, intErr)

		return
	}
	if url == "/i/siget/" {
		b, parErr, intErr := dataBase.GetItemFromTableIntFromByStringIndexJSON(ctx.PostBody())
		fastHTTPHandlerDblErrorBodyJSON(ctx, b, parErr, intErr)

		return
	}
	if url == "/s/siget/" {
		b, parErr, intErr := dataBase.GetItemFromTableIntFromByStringIndexJSON(ctx.PostBody())
		fastHTTPHandlerDblErrorBodyJSON(ctx, b, parErr, intErr)

		return
	}

	if url == "/c/struct/" {
		b, err := dataBase.StructGet()
		fastHTTPHandlerBodyJSON(ctx, b, err)

		return
	}

	if url == "/i/create_table/" {
		parErr, intErr := dataBase.CreateTableIntFromJSON(ctx.PostBody())
		fastHTTPHandlerDblError(ctx, parErr, intErr)

		return
	}
	if url == "/s/create_table/" {
		parErr, intErr := dataBase.CreateTableStringFromJSON(ctx.PostBody())
		fastHTTPHandlerDblError(ctx, parErr, intErr)

		return
	}

	if url == "/i/create_index/" {
		parErr, intErr := dataBase.CreateIndexOnTableIntFromJSON(ctx.PostBody())
		fastHTTPHandlerDblError(ctx, parErr, intErr)

		return
	}
	if url == "/s/create_index/" {
		parErr, intErr := dataBase.CreateIndexOnTableStringFromJSON(ctx.PostBody())
		fastHTTPHandlerDblError(ctx, parErr, intErr)

		return
	}

	if url == "/sec/token_add/" {
		parErr, intErr := dataBase.AddTokenFromJSON(ctx.PostBody())
		fastHTTPHandlerDblError(ctx, parErr, intErr)

		return
	}
	if url == "/sec/token_remove/" {
		parErr, intErr := dataBase.RMTokenFromJSON(ctx.PostBody())
		fastHTTPHandlerDblError(ctx, parErr, intErr)

		return
	}

	fmt.Fprint(ctx, "Not found")
	ctx.Response.SetStatusCode(404)
}

func logsOut(err error) {
	log.Errorln(err)
}

var (
	dataBase *DB
)

func main() {

	initConf()

	db, err := CreateDB(config.DBFolder, config.DBFlushTimeout, logsOut)
	dataBase = db

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(config)

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

	case <-osSignals:
		log.Infof("Start shutdown...")
		go func() {
			if err := api.Shutdown(); err != nil {
				log.Infof("Graceful shutdown did not complete in 5s : %v", err)
			}
		}()
	}

	fmt.Println("Good Bye")
}
