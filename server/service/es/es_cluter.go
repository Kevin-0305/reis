package es

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/es"
	esReq "github.com/flipped-aurora/gin-vue-admin/server/model/es/request"
)

type EsCluterService struct {
}

// CreateEsCluter 创建EsCluter记录
// Author [piexlmax](https://github.com/piexlmax)
func (escService *EsCluterService) CreateEsCluter(esc es.EsCluter) (err error) {
	err = global.GVA_DB.Create(&esc).Error
	return err
}

// DeleteEsCluter 删除EsCluter记录
// Author [piexlmax](https://github.com/piexlmax)
func (escService *EsCluterService) DeleteEsCluter(esc es.EsCluter) (err error) {
	err = global.GVA_DB.Delete(&esc).Error
	return err
}

// DeleteEsCluterByIds 批量删除EsCluter记录
// Author [piexlmax](https://github.com/piexlmax)
func (escService *EsCluterService) DeleteEsCluterByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]es.EsCluter{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateEsCluter 更新EsCluter记录
// Author [piexlmax](https://github.com/piexlmax)
func (escService *EsCluterService) UpdateEsCluter(esc es.EsCluter) (err error) {
	err = global.GVA_DB.Save(&esc).Error
	return err
}

// GetEsCluter 根据id获取EsCluter记录
// Author [piexlmax](https://github.com/piexlmax)
func (escService *EsCluterService) GetEsCluter(id uint) (esc es.EsCluter, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&esc).Error
	return
}

// GetEsCluterInfoList 分页获取EsCluter记录
// Author [piexlmax](https://github.com/piexlmax)
func (escService *EsCluterService) GetEsCluterInfoList(info esReq.EsCluterSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&es.EsCluter{})
	var escs []es.EsCluter
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CluterName != "" {
		db = db.Where("cluter_name LIKE ?", "%"+info.CluterName+"%")
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.Version != "" {
		db = db.Where("version > ?", info.Version)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&escs).Error
	return escs, total, err
}

func (escService *EsCluterService) CheckEsCluterStatus(esc es.EsCluter) (status int, err error) {
	status = esc.CheckState()
	return
}
