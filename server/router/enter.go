package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/alert"
	"github.com/flipped-aurora/gin-vue-admin/server/router/es"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/prom"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/webhook"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Es      es.RouterGroup
	Prom    prom.RouterGroup
	Alert   alert.RouterGroup
	WebHook webhook.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
