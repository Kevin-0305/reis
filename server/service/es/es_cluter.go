package es

import (
	"fmt"
	"sync"

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
	escService.RefreshCluterGroup(esc.ID, esc.GroupIds)
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
	escService.RefreshCluterGroup(esc.ID, esc.GroupIds)
	return err
}

// GetEsCluter 根据id获取EsCluter记录
// Author [piexlmax](https://github.com/piexlmax)
func (escService *EsCluterService) GetEsCluter(id uint) (esc es.EsCluter, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Group").First(&esc).Error
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
	if info.Version != "" {
		db = db.Where("version > ?", info.Version)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Group").Find(&escs).Error

	var wg sync.WaitGroup
	for k, _ := range escs {
		wg.Add(1)
		go func(esc *es.EsCluter) {
			defer wg.Done()
			info, err := esc.GetInfo()
			if err != nil {
				fmt.Println(err)
				return
			} else {
				esc.Status = info["status"].(string)
				esc.NodesNumber = int(info["number_of_data_nodes"].(float64))
			}
		}(&escs[k])
	}
	wg.Wait()
	return escs, total, err
}

func (escService *EsCluterService) CheckEsCluterStatus(esc es.EsCluter) (status int, err error) {
	status = esc.CheckState()
	return
}

// 刷新ES集群所属的分组
func (escService *EsCluterService) RefreshCluterGroup(cluterId uint, groupIds []uint) (err error) {
	db := global.GVA_DB.Model(&es.EsCluterGroup{})
	err = db.Where("es_cluter_id = ?", cluterId).Delete(&es.EsCluterGroup{}).Error
	if err != nil {
		return
	}
	if len(groupIds) > 0 {
		escGroups := []es.EsCluterGroup{}
		for _, groupId := range groupIds {
			escGroup := es.EsCluterGroup{
				EsCluterId:     cluterId,
				ProjectGroupId: groupId,
			}
			escGroups = append(escGroups, escGroup)
		}
		err = db.Create(&escGroups).Error
		if err != nil {
			return
		}
	}
	return
}
