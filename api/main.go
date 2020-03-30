package main

import (
	"api/config"
	"api/route"

	"log"

	"github.com/AubSs/fasthttplogger"
	"github.com/valyala/fasthttp"
)

func main() {
	router := route.Router()
	port := config.Port()

	router = fasthttplogger.CombinedColored(router)

	log.Fatal(fasthttp.ListenAndServe(port, router))
}
