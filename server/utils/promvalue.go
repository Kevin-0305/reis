package utils

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

func QueryVectorValue(serviceName string, pql string) (model.Vector, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	value, _, err := global.GVA_PromAPIs[serviceName].Query(ctx, pql, time.Now())
	if err != nil {
		//global.GVA_LOG.Error("query error", zap.Error(err))
		return nil, err
	}
	if value.Type() != model.ValVector {
		//global.GVA_LOG.Error("query error", zap.Error(err))
		return nil, fmt.Errorf("value is not vector")
	}
	v, ok := value.(model.Vector)
	if !ok {
		return nil, fmt.Errorf("value is not vector")
	}
	return v, nil
}

func QueryMetricValue(serviceName string, pql string, startTime time.Time, endTime time.Time, step time.Duration) (model.Matrix, error) {
	r := v1.Range{
		Start: startTime,
		End:   endTime,
		Step:  step,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	value, _, err := global.GVA_PromAPIs[serviceName].QueryRange(ctx, pql, r)
	if err != nil {
		//global.GVA_LOG.Error("query error", zap.Error(err))
		return nil, err
	}
	if value.Type() != model.ValMatrix {
		//global.GVA_LOG.Error("query error", zap.Error(err))
		return nil, fmt.Errorf("value is not Metric")
	}
	v, ok := value.(model.Matrix)
	if !ok {
		return nil, fmt.Errorf("value is not Metric")
	}
	return v, nil
}
func QueryHostValue(serviceName string, pql string, instance string) (float64, error) {
	pql = fmt.Sprintf("%s{instance=\"%s\"}", pql, instance)
	v, err := QueryVectorValue(serviceName, pql)
	if err != nil {
		return 0, err
	}
	vv, err := ParseVectorValue(v)
	if err != nil {
		return 0, err
	}
	result := vv[instance]
	return result, nil
}

func ParseVectorValue(value model.Vector) (map[string]float64, error) {
	result := make(map[string]float64)
	for _, v := range value {
		vv, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(v.Value)), 64)
		result[string(v.Metric["instance"])] = vv
	}
	return result, nil
}

// func ParseMetricValue(value model.Matrix)(map[string]float64,error){

// }

func ParseMultiVectorValue(value model.Vector, key string) (map[string]map[string]float64, error) {
	result := make(map[string]map[string]float64)
	labelName := model.LabelName(key)
	for _, v := range value {
		instance := string(v.Metric["instance"])
		if _, ok := result[instance]; !ok {
			result[instance] = make(map[string]float64)
		}
		result[instance][string(v.Metric[labelName])] = float64(v.Value)
	}
	return result, nil
}

func ParseLabelValue(value model.Vector, key string) (map[string]string, error) {
	result := make(map[string]string)
	labelName := model.LabelName(key)
	for _, v := range value {
		result[string(v.Metric["instance"])] = string(v.Metric[labelName])
	}
	return result, nil
}

func ParseMultiLabelValue(value model.Vector, keys ...string) (map[string]map[string]string, error) {
	result := make(map[string]map[string]string)
	var labelName model.LabelName
	for _, v := range value {
		for _, key := range keys {
			if _, ok := result[string(v.Metric["instance"])]; !ok {
				result[string(v.Metric["instance"])] = make(map[string]string)
			}
			labelName = model.LabelName(key)
			result[string(v.Metric["instance"])][key] = string(v.Metric[labelName])
		}
	}
	return result, nil
}

func MetricToMap(metric model.Metric) map[string]string {
	result := make(map[string]string)
	for k, v := range metric {
		if k != "__name__" {
			result[string(k)] = string(v)
		}
	}
	return result
}

// func ParseMatrixValue(value model.Matrix) (map[string]float64, error) {
// 	result := make(map[string]float64)
// 	for _, v := range value {
// 		for _, j := range v.Values {
// 			result[string(j.Metric["instance"])] = float64(j.Value)
// 		}
// 	}
// 	return result, nil
// }
