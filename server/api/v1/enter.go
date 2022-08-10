package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/alert"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/es"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/prom"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/webhook"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	EsApiGroup      es.ApiGroup
	PromApiGroup    prom.ApiGroup
	AlertApiGroup   alert.ApiGroup
	WebHookApiGroup webhook.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
