package prom

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type PromHostService struct {
}

func (phService *PromHostService) GetHostRealInfo(serviceName string, instance string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	gpuTemperature, err := utils.QueryHostValue(serviceName, "winbase_gpuScrape_GPUTemperature", instance)
	if err != nil {
		return nil, err
	}
	result["gpuTemperature"] = gpuTemperature
	cpuTemperature, err := utils.QueryHostValue(serviceName, "winbase_cpuScrape_CPUTemperature", instance)
	if err != nil {
		return nil, err
	}
	result["cpuTemperature"] = cpuTemperature

	timeNow := time.Now()
	min := timeNow.Minute()
	sec := timeNow.Second()
	endTime := timeNow.Add(-time.Duration(time.Duration(min) * time.Minute)).Add(-time.Duration(time.Duration(sec) * time.Second))
	startTime := endTime.Add(-time.Duration(1 * 24 * time.Hour))
	step := time.Duration(1 * 1 * time.Hour)

	labels := map[string]string{"instance": instance}
	// diskLabels := map[string]string{"instance": instance, "device": "sda"}
	cpuUsageData, err := queryValue(serviceName, labels, "winbase_cpuScrape_Load1Hour", startTime, endTime, step)
	gpuUsageData, err := queryValue(serviceName, labels, "winbase_gpuScrape_Load1Hour", startTime, endTime, step)
	memoryUsageData, err := queryValue(serviceName, labels, "winbase_memoryScrape_UsePercent", startTime, endTime, step)
	diskUsageData, err := utils.QueryMetricValue(serviceName, "winbase_diskScrape_disk{type=\"Used\",instance=\""+instance+"\"}", startTime, endTime, step)
	diskFreeData, err := utils.QueryMetricValue(serviceName, "winbase_diskScrape_disk{type=\"Free\",instance=\""+instance+"\"}", startTime, endTime, step)

	cpuUsageRealVector, _ := utils.QueryVectorValue(serviceName, "winbase_cpuScrape_Load0")
	gpuUsageRealVector, _ := utils.QueryVectorValue(serviceName, "winbase_gpuScrape_GPUUsePercent")
	memoryUsageRealVector, _ := utils.QueryVectorValue(serviceName, "winbase_memoryScrape_UsePercent")

	cpuUsageReal, _ := utils.ParseVectorValue(cpuUsageRealVector)
	gpuUsageReal, _ := utils.ParseVectorValue(gpuUsageRealVector)
	memoryUsageReal, _ := utils.ParseVectorValue(memoryUsageRealVector)

	timeList := make([]string, 0)
	cpuUsageList := make([]float64, 0)
	gpuUsageList := make([]float64, 0)
	memoryUsageList := make([]float64, 0)

	diskPathList := make([]string, 0)
	diskUsageList := make([]float64, 0)
	diskFreeList := make([]float64, 0)

	for i := 1; i <= 24; i++ {
		hour := (timeNow.Hour() + i) % 24
		day := timeNow.Add(time.Duration((timeNow.Hour()+i)/24-1) * 24 * time.Hour).Day()
		t := strconv.Itoa(day) + " " + strconv.Itoa(hour) + ":00"
		timeList = append(timeList, t)

		if _, ok := cpuUsageData[hour]; ok {
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", cpuUsageData[hour]), 64)
			cpuUsageList = append(cpuUsageList, value)
		} else {
			cpuUsageList = append(cpuUsageList, 0)
		}

		if _, ok := gpuUsageData[hour]; ok {
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", gpuUsageData[hour]*100), 64)
			gpuUsageList = append(gpuUsageList, value)
		} else {
			gpuUsageList = append(gpuUsageList, 0)
		}

		if _, ok := memoryUsageData[hour]; ok {
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", memoryUsageData[hour]), 64)
			memoryUsageList = append(memoryUsageList, value)
		} else {
			memoryUsageList = append(memoryUsageList, 0)
		}
	}

	for _, v := range diskUsageData {
		path := string(v.Metric["path"])
		diskPathList = append(diskPathList, path)
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(v.Values[0].Value)/1024/1024/1024/1024), 64)
		diskUsageList = append(diskUsageList, value)
	}
	for _, v := range diskFreeData {
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(v.Values[0].Value)/1024/1024/1024/1024), 64)
		diskFreeList = append(diskFreeList, value)
	}

	result["time"] = timeList
	result["cpuUse"] = cpuUsageList
	result["gpuUse"] = gpuUsageList
	result["memoryUse"] = memoryUsageList
	result["diskPath"] = diskPathList
	result["diskUse"] = diskUsageList
	result["diskFree"] = diskFreeList
	result["cpuUseReal"] = cpuUsageReal[instance]
	result["gpuUseReal"] = gpuUsageReal[instance]
	result["memoryUseReal"] = memoryUsageReal[instance]
	return result, nil
}

func queryValue(serviceName string, labels map[string]string, pql string, startTime time.Time, endTime time.Time, step time.Duration) (map[int]float64, error) {

	//pql = pql + "{" + "instance=\"" + instance + "\"" + "}"
	// labelsBytes, err := json.Marshal(labels)
	// dataString := string(labelsBytes)
	// pql = pql + dataString
	// pql = strings.Replace(pql, ":", "=", -1)
	labelsString := ""
	for k, v := range labels {
		labelsString = labelsString + k + "=\"" + v + "\","
	}
	labelsString = "{" + strings.TrimRight(labelsString, ",") + "}"
	pql = pql + labelsString
	fmt.Println(pql)
	usageMatrix, err := utils.QueryMetricValue(serviceName, pql, startTime, endTime, step)
	if err != nil {
		return nil, err
	}
	usageData := make(map[int]float64, 0)
	for _, v := range usageMatrix {
		for _, vv := range v.Values {
			tm := time.Unix(int64(vv.Timestamp/1000), 0)
			usageData[tm.Hour()] = float64(vv.Value)
		}
	}
	return usageData, nil
}

func (phService *PromHostService) GetActiveHostList(serviceName string) ([]map[string]string, error) {
	timeEnd := time.Now()
	timeStart := timeEnd.Add(-time.Duration(7 * 24 * time.Hour))
	var r v1.Range
	r.Start = timeStart
	r.End = timeEnd
	r.Step = time.Duration(1 * 1 * time.Hour)
	value, _, err := global.GVA_PromAPIs[serviceName].QueryRange(context.Background(), "winbase_hostScrape_host", r)
	if err != nil {
		return nil, err
	}
	if value.Type() != model.ValMatrix {
		return nil, err
	}
	v, _ := value.(model.Matrix)
	var result []map[string]string
	for _, vv := range v {
		result = append(result, utils.MetricToMap(vv.Metric))
	}
	return result, nil
}
