package route

import (
	h "api/handler"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Router() func(*fasthttp.RequestCtx) {
	r := router.New()

	r.GET("/number_of_daily_message", h.GetNumberOfDailyMessage)
	r.GET("/top_accounts_by_messages", h.GetAccountsByMessages)
	r.GET("/top_messages_by_engagements", h.GetMessagesByEngagements)
	r.GET("/words", h.GetWords)
	r.GET("/hashtags", h.GetHashtags)

	return r.Handler
}
