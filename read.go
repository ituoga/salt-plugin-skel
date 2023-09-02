package module

import "github.com/ituoga/salt"

func onRead(router *salt.Router) {
	router.Handle(topicRead, func(ctx *salt.Context) {
		ctx.Response().Reply("Read")
	})
}
