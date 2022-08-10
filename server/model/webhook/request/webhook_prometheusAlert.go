package request

import "time"

type PrometheusAlertRequest struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	Alerts            []Alerts          `json:"alerts"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations CommonAnnotations `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	AlertTime         time.Time         `json:"alertTime"`
}

// type Labels struct {
// 	Alertname string `json:"alertname"`
// 	Instance  string `json:"instance"`
// 	Job       string `json:"job"`
// }
type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}
type Alerts struct {
	Status       string            `json:"status"`
	Labels       map[string]string `json:"labels"`
	Annotations  Annotations       `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`
	EndsAt       time.Time         `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Fingerprint  string            `json:"fingerprint"`
}

// type GroupLabels struct {
// 	Alertname string `json:"alertname"`
// }
// type CommonLabels struct {
// 	Alertname string `json:"alertname"`
// 	Instance  string `json:"instance"`
// 	Job       string `json:"job"`
// }
type CommonAnnotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}
