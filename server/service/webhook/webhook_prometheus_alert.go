package webhook

import (
	"encoding/json"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"

	webhookReq "github.com/flipped-aurora/gin-vue-admin/server/model/webhook/request"
	alertService "github.com/flipped-aurora/gin-vue-admin/server/service/alert"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/es"
	esModel "github.com/flipped-aurora/gin-vue-admin/server/utils/es/model"
	elasticV7 "github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type PrometheusAlertService struct {
}

func (s *PrometheusAlertService) PrometheusAlert(PrometheusAlertRequest webhookReq.PrometheusAlertRequest) error {
	title := "Prometheus告警消息"
	content := ""
	endTime := time.Now()

	for _, alert := range PrometheusAlertRequest.Alerts {
		if alert.Status == "firing" {
			endTime = time.Now()
		} else if alert.Status == "resolved" {
			endTime = alert.EndsAt
		} else {
			endTime = time.Now()
		}
		content += "告警名称：" + alert.Labels["alertname"] + "\n"
		content += "告警等级：" + alert.Labels["severity"] + "\n"
		content += "告警描述：" + alert.Annotations.Description + "\n"
		content += "告警详细：" + alert.Annotations.Summary + "\n"
		content += "告警开始时间：" + alert.StartsAt.Add(28800e9).Format("2006-01-02 15:04:05.000") + "\n"
		content += "告警结束时间：" + alert.EndsAt.Add(28800e9).Format("2006-01-02 15:04:05.000") + "\n"
		content += "告警源：" + alert.Labels["instance"] + "\n"
		content += "告警状态：" + alert.Status + "\n"
		content += "异常持续时间：" + endTime.Sub(alert.StartsAt).String() + "\n"
		content += "\n"
	}
	alertService.PostToFS(title, content, global.GVA_CONFIG.Alert.Fs.WebHookUrl, "")
	esConnect := esModel.EsConnect{
		Ip:      global.GVA_CONFIG.Alert.Elastic.Address,
		User:    global.GVA_CONFIG.Alert.Elastic.Account,
		Pwd:     global.GVA_CONFIG.Alert.Elastic.Password,
		Version: global.GVA_CONFIG.Alert.Elastic.Version,
	}

	insertJson, _ := json.Marshal(PrometheusAlertRequest)
	insertMap := make(map[string]interface{})
	_ = json.Unmarshal(insertJson, &insertMap)
	esInsert := esModel.EsDocUpdateByID{
		Index: "prometheus_alert",
		Type:  "_doc",
		JSON:  insertMap,
	}
	esClientService, err := es.NewEsService(&esConnect)
	if err != nil {
		global.GVA_LOG.Error("esClientService", zap.Error(err))
		return err
	}
	res, err := esClientService.EsDocInsert(&esInsert)
	if err != nil {
		global.GVA_LOG.Error("esClientService", zap.Error(err))
		return err
	}
	if esConnect.Version == 7 {
		esResponse := res.(*elasticV7.IndexResponse)
		if esResponse.Status != 0 {
			global.GVA_LOG.Error("esClientService", zap.Error(err))
			return err
		}
	}

	return nil
}
