package module

import "github.com/ituoga/salt"

func onDelete(router *salt.Router) {
	router.Handle(topicDelete, func(ctx *salt.Context) {
		ctx.Response().Reply("Delete")
	})
}
