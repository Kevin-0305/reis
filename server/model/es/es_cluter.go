// 自动生成模板EsCluter
package es

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EsCluter 结构体
// 如果含有time.Time 请自行import time包
type EsCluter struct {
	global.GVA_MODEL
	CluterName  string `json:"cluterName" form:"cluterName" gorm:"column:cluter_name;comment:集群名称;size:50;"`
	Status      *int   `json:"status" form:"status" gorm:"column:status;comment:集群状态;size:1;"`
	Version     string `json:"version" form:"version" gorm:"column:version;comment:集群ES版本;size:20;"`
	Address     string `json:"address" form:"address" gorm:"column:address;comment:集群地址;size:30;"`
	Monitor     *bool  `json:"monitor" form:"monitor" gorm:"column:monitor;comment:是否启用监控;size:1;"`
	TLS         *bool  `json:"TLS" form:"TLS" gorm:"column:TLS;comment:TLS;"`
	Auth        *bool  `json:"auth" form:"auth" gorm:"column:auth;comment:身份验证;"`
	UserName    string `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;size:20;"`
	Password    string `json:"password" form:"password" gorm:"column:password;comment:密码;size:50;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述;size:200;"`
}

// TableName EsCluter 表名
func (EsCluter) TableName() string {
	return "es_cluter"
}

func (es *EsCluter) CheckState() int {
	url := "http://" + es.Address + "/_cluster/health"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(es.Address, es.Password)
	resp, err := client.Do(req)
	if err != nil {
		return 0
	}
	result, _ := ioutil.ReadAll(resp.Body)
	resultMap := make(map[string]interface{})
	_ = json.Unmarshal(result, &resultMap)

	defer resp.Body.Close()
	if resultMap["status"].(string) == "red" {
		return 1
	} else if resultMap["status"].(string) == "yellow" {
		return 2
	} else if resultMap["status"].(string) == "green" {
		return 3
	}
	return 0
}
