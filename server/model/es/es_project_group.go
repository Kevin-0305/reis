// 自动生成模板ProjectGroup
package es

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ProjectGroup 结构体
// 如果含有time.Time 请自行import time包
type ProjectGroup struct {
	global.GVA_MODEL
	Name     string         `json:"name" form:"name" gorm:"column:name;comment:组名;size:20;"`
	Level    *int           `json:"level" form:"level" gorm:"column:level;comment:分组等级;"`
	ParentId uint           `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父组ID;"`
	Head     string         `json:"head" form:"head" gorm:"column:head;comment:负责人;"`
	Children []ProjectGroup `json:"children" form:"children" gorm:"-"`
	Cluter   []EsCluter     `json:"authorities" gorm:"many2many:es_cluter_group;"`
}

// TableName ProjectGroup 表名
func (ProjectGroup) TableName() string {
	return "project_group"
}
