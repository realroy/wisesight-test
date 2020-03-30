package handler

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func RespondJSON(ctx *fasthttp.RequestCtx, jsonPayload map[string]interface{}) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.SetStatusCode(fasthttp.StatusOK)

	result, err := json.Marshal(jsonPayload)

	if err != nil {
		ctx.Error("Can't marshal json", fasthttp.StatusInternalServerError)
		return
	}

	ctx.Write(result)
}
