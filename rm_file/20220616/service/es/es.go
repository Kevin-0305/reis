package es

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/es"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    esReq "github.com/flipped-aurora/gin-vue-admin/server/model/es/request"
)

type EsCluterGroupService struct {
}

// CreateEsCluterGroup 创建EsCluterGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (ecgService *EsCluterGroupService) CreateEsCluterGroup(ecg es.EsCluterGroup) (err error) {
	err = global.GVA_DB.Create(&ecg).Error
	return err
}

// DeleteEsCluterGroup 删除EsCluterGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (ecgService *EsCluterGroupService)DeleteEsCluterGroup(ecg es.EsCluterGroup) (err error) {
	err = global.GVA_DB.Delete(&ecg).Error
	return err
}

// DeleteEsCluterGroupByIds 批量删除EsCluterGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (ecgService *EsCluterGroupService)DeleteEsCluterGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]es.EsCluterGroup{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEsCluterGroup 更新EsCluterGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (ecgService *EsCluterGroupService)UpdateEsCluterGroup(ecg es.EsCluterGroup) (err error) {
	err = global.GVA_DB.Save(&ecg).Error
	return err
}

// GetEsCluterGroup 根据id获取EsCluterGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (ecgService *EsCluterGroupService)GetEsCluterGroup(id uint) (ecg es.EsCluterGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ecg).Error
	return
}

// GetEsCluterGroupInfoList 分页获取EsCluterGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (ecgService *EsCluterGroupService)GetEsCluterGroupInfoList(info esReq.EsCluterGroupSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&es.EsCluterGroup{})
    var ecgs []es.EsCluterGroup
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&ecgs).Error
	return  ecgs, total, err
}
