package initialize

import (
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func Prom() {
	APIs := make(map[string]v1.API)
	for _, value := range global.GVA_CONFIG.Prom.Services {
		config := api.Config{
			Address: value.Address,
		}
		client, err := api.NewClient(config)
		if err != nil {
			log.Println("api.NewClient error : ", err)
		}
		api := v1.NewAPI(client)
		if err != nil {
			log.Println("client.NewAPI error : ", err)
		}
		APIs[value.Name] = api
	}

	global.GVA_PromAPIs = APIs
}
