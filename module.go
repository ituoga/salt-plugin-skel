package module

import (
	"github.com/ituoga/salt"
)

func OnInit(router *salt.Router) {
	onCreate(router)
	onRead(router)
	onUpdate(router)
	onDelete(router)
	onIndex(router)
}
