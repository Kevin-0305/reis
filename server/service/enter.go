package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/alert"
	"github.com/flipped-aurora/gin-vue-admin/server/service/es"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/prom"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/webhook"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	EsServiceGroup      es.ServiceGroup
	PromServiceGroup    prom.ServiceGroup
	AlertServiceGroup   alert.ServiceGroup
	WebHookServiceGroup webhook.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
