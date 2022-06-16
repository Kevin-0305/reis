// 自动生成模板EsCluterGroup
package es

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EsCluterGroup 结构体
// 如果含有time.Time 请自行import time包
type EsCluterGroup struct {
      global.GVA_MODEL
      EsCluterId  *int `json:"esCluterId" form:"esCluterId" gorm:"column:es_cluter_id;comment:Es集群ID;"`
      ProjectGroupId  *int `json:"projectGroupId" form:"projectGroupId" gorm:"column:project_group_id;comment:Es分组ID;"`
}


// TableName EsCluterGroup 表名
func (EsCluterGroup) TableName() string {
  return "es_cluter_group"
}

