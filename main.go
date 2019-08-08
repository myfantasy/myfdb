package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func main() {

	log.Infof("App is starting up: %v", time.Now())

	settingsLoad()

	err := runStorage()

	if err != nil {
		log.Fatalf(err.Error())
	}

	api := &fasthttp.Server{
		Handler: fastHTTPHandler,
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Infof("Listen and serve %s", apiSettings.Addr)
		serverErrors <- api.ListenAndServe(apiSettings.Addr)
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

	log.Infof("App is stopped")
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	method := string(ctx.Method())
	url := string(ctx.URI().Path())

	if strings.Index(url, "/get/") == 0 && method == "GET" {
		getItemHandler(ctx, url)
		return
	}

	if url == "/api-health" && method == "GET" {
		healthHandler(ctx)
		return
	}

	if url == "/crtbli" && method == "POST" {
		createITable(ctx, method, url)
		return
	}

	if strings.Index(url, "/seti/") == 0 && method == "POST" {
		setIValue(ctx, method, url)
		return
	}

	if strings.Index(url, "/sets/") == 0 && method == "POST" {
		setSValue(ctx, method, url)
		return
	}

	if strings.Index(url, "/deli/") == 0 && method == "POST" {
		delIValue(ctx, method, url)
		return
	}

	if strings.Index(url, "/dels/") == 0 && method == "POST" {
		delSValue(ctx, method, url)
		return
	}

	if strings.Index(url, "/geti/") == 0 && method == "POST" {
		getIValue(ctx, method, url)
		return
	}

	if strings.Index(url, "/gets/") == 0 && method == "POST" {
		getSValue(ctx, method, url)
		return
	}

	methodNotAllowedHandler(ctx, method, url)
}

func healthHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(200)
	fmt.Fprint(ctx, "ok")
}

func methodNotAllowedHandler(ctx *fasthttp.RequestCtx, method string, url string) {
	ctx.Response.SetStatusCode(405)
	fmt.Fprint(ctx, "Method Not Allowed", " method:", method, "; url:", url)
}

func getItemHandler(ctx *fasthttp.RequestCtx, url string) {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		ctx.Response.SetStatusCode(400)
		fmt.Fprint(ctx, "Bad request:( Expected url format /get/table/key")
	}
	fmt.Fprint(ctx, parts[1])
}

func internalServiceError(ctx *fasthttp.RequestCtx, err error) {
	ctx.Response.SetStatusCode(503)
	if apiSettings.OutputInternalErrors {
		fmt.Fprint(ctx, err.Error())
	} else {
		fmt.Fprint(ctx, "Internal error.")
	}

	log.Errorln(err)
}

func okHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.SetStatusCode(200)

	fmt.Fprint(ctx, "ok")
}

func okAndItemHandler(ctx *fasthttp.RequestCtx, item interface{}) {
	ctx.Response.SetStatusCode(200)

	fmt.Fprint(ctx, item)
}
