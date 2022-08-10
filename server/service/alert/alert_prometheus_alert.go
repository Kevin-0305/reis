package alert

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"

	"github.com/flipped-aurora/gin-vue-admin/server/model/alert/request"
	webhookReq "github.com/flipped-aurora/gin-vue-admin/server/model/webhook/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/es"
	esModel "github.com/flipped-aurora/gin-vue-admin/server/utils/es/model"
	"github.com/olivere/elastic"
	elasticV7 "github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

func (s *AlertService) GetPrometheusAlertHistory(alertHistoryRequest request.GetPrometheusAlertHistoryRequest) (interface{}, int64, error) {
	esConnect := esModel.EsConnect{
		Ip:      global.GVA_CONFIG.Alert.Elastic.Address,
		User:    global.GVA_CONFIG.Alert.Elastic.Account,
		Pwd:     global.GVA_CONFIG.Alert.Elastic.Password,
		Version: global.GVA_CONFIG.Alert.Elastic.Version,
	}
	esClientService, err := es.NewEsService(&esConnect)
	if err != nil {
		global.GVA_LOG.Error("esClientService", zap.Error(err))
		return nil, 0, err
	}
	esDocSearch := esModel.EsDocSearch{
		IndexName: "prometheus_alert",
		Type:      "_doc",
		Page:      alertHistoryRequest.PageInfo.Page,
		Limit:     alertHistoryRequest.PageInfo.PageSize,
		Query:     elastic.NewBoolQuery(),
		Sort:      elastic.NewFieldSort("alertTime"),
	}

	fmt.Println(alertHistoryRequest)
	nilTime := time.Time{}
	if alertHistoryRequest.AlertName != "" {
		esDocSearch.Query = esDocSearch.Query.Must(elastic.NewTermQuery("commonLabels.alertName", alertHistoryRequest.AlertName))
	} else if alertHistoryRequest.Status != "" {
		esDocSearch.Query = esDocSearch.Query.Must(elastic.NewTermQuery("status", alertHistoryRequest.Status))
	} else if alertHistoryRequest.Severity != "" {
		esDocSearch.Query = esDocSearch.Query.Must(elastic.NewTermQuery("commonLabels.severity", alertHistoryRequest.Severity))
	} else if alertHistoryRequest.Job != "" {
		esDocSearch.Query = esDocSearch.Query.Must(elastic.NewTermQuery("commonLabels.job", alertHistoryRequest.Job))
	} else if alertHistoryRequest.Instance != "" {
		esDocSearch.Query = esDocSearch.Query.Must(elastic.NewWildcardQuery("commonLabels.instance", alertHistoryRequest.Instance+"*"))
	} else if alertHistoryRequest.AlertTimeGte != nilTime && alertHistoryRequest.AlertTimeLte != nilTime {
		esDocSearch.Query = esDocSearch.Query.Must(elastic.NewRangeQuery("alertTime").Gte(alertHistoryRequest.AlertTimeGte).Lte(alertHistoryRequest.AlertTimeLte))
	} else if alertHistoryRequest.AlertTimeGte != nilTime && alertHistoryRequest.AlertTimeLte == nilTime {
		esDocSearch.Query = esDocSearch.Query.Must(elastic.NewRangeQuery("alertTime").Gte(alertHistoryRequest.AlertTimeGte))
	} else if alertHistoryRequest.AlertTimeGte == nilTime && alertHistoryRequest.AlertTimeLte != nilTime {
		esDocSearch.Query = esDocSearch.Query.Must(elastic.NewRangeQuery("alertTime").Lte(alertHistoryRequest.AlertTimeLte))
	}

	if alertHistoryRequest.PageInfo.Keyword == "ASC" {
		esDocSearch.Sort = esDocSearch.Sort.Asc()
	} else if alertHistoryRequest.PageInfo.Keyword == "DESC" {
		esDocSearch.Sort = esDocSearch.Sort.Desc()
	}

	res, count, err := esClientService.EsDocSearch(&esDocSearch)
	if err != nil {
		global.GVA_LOG.Error("esClientService", zap.Error(err))
		return nil, 0, err
	}

	var alertMessage webhookReq.PrometheusAlertRequest
	var alerts []webhookReq.Alerts
	if esConnect.Version == 7 {
		esResponse := res.(*elasticV7.SearchResult)
		if esResponse.Status != 0 {
			global.GVA_LOG.Error("esClientService", zap.Error(err))
			return nil, 0, err
		}
		for _, hit := range esResponse.Hits.Hits {
			err := json.Unmarshal(hit.Source, &alertMessage)
			if err != nil {
				global.GVA_LOG.Error("esClientService", zap.Error(err))
				return nil, 0, err
			}
			for _, alert := range alertMessage.Alerts {
				alerts = append(alerts, alert)
			}
		}
	}

	return alerts, count, err
}
