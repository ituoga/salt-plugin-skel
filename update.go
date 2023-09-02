package module

import "github.com/ituoga/salt"

func onUpdate(router *salt.Router) {
	router.Handle(topicUpdate, func(ctx *salt.Context) {
		ctx.Response().Reply("Update")
	})
}
