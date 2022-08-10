package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type GetPrometheusAlertHistoryRequest struct {
	PageInfo     request.PageInfo `json:"pageInfo" form:"pageInfo"`
	AlertName    string           `json:"alertName" form:"alertName"`
	Status       string           `json:"status" form:"status"`
	AlertTimeGte time.Time        `json:"alertTimeGte" form:"alertTimeGte"`
	AlertTimeLte time.Time        `json:"alertTimeLte" form:"alertTimeLte"`
	Job          string           `json:"job" form:"job"`
	Instance     string           `json:"instance" form:"instance"`
	Severity     string           `json:"severity" form:"severity"`
}
