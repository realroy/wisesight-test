package handler

import (
	"api/model"

	"github.com/valyala/fasthttp"
)

func GetNumberOfDailyMessage(ctx *fasthttp.RequestCtx) {
	count, err := model.GetMessageCount()

	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	payload := map[string]interface{}{
		"count": count,
	}

	RespondJSON(ctx, payload)
}

func GetAccountsByMessages(ctx *fasthttp.RequestCtx) {
	results, err := model.GetAccountsByMessages()

	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	payload := map[string]interface{}{
		"top_accounts_by_messages": results,
	}

	RespondJSON(ctx, payload)
}

func GetMessagesByEngagements(ctx *fasthttp.RequestCtx) {
	results, err := model.GetMessageByEngagements()

	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	payload := map[string]interface{}{
		"top_messages_by_engagements": results,
	}

	RespondJSON(ctx, payload)
}

func GetWords(ctx *fasthttp.RequestCtx) {
	results, err := model.GetWords()

	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	payload := map[string]interface{}{
		"words": results,
	}

	RespondJSON(ctx, payload)
}

func GetHashtags(ctx *fasthttp.RequestCtx) {
	results, err := model.GetHashtags()

	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	payload := map[string]interface{}{
		"hashtags": results,
	}

	RespondJSON(ctx, payload)
}
