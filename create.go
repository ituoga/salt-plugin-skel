package module

import "github.com/ituoga/salt"

func onCreate(router *salt.Router) {
	router.Handle(topicCreate, func(ctx *salt.Context) {
		ctx.Response().Reply("create")
	})
}
