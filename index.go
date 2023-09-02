package module

import "github.com/ituoga/salt"

func onIndex(router *salt.Router) {
	router.Handle(topicIndex, func(ctx *salt.Context) {
		if ctx.Can("view-own") {
			ctx.Response().Reply("index own")
			return
		}
		ctx.Response().Reply("index")
	})
}
