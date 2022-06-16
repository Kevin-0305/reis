package es

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/es"
	esReq "github.com/flipped-aurora/gin-vue-admin/server/model/es/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ProjectGroupService struct {
}

// CreateProjectGroup 创建ProjectGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (pgService *ProjectGroupService) CreateProjectGroup(pg es.ProjectGroup) (err error) {
	err = global.GVA_DB.Create(&pg).Error
	return err
}

// DeleteProjectGroup 删除ProjectGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (pgService *ProjectGroupService) DeleteProjectGroup(pg es.ProjectGroup) (err error) {
	err = global.GVA_DB.Delete(&pg).Error
	return err
}

// DeleteProjectGroupByIds 批量删除ProjectGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (pgService *ProjectGroupService) DeleteProjectGroupByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]es.ProjectGroup{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProjectGroup 更新ProjectGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (pgService *ProjectGroupService) UpdateProjectGroup(pg es.ProjectGroup) (err error) {
	err = global.GVA_DB.Save(&pg).Error
	return err
}

// GetProjectGroup 根据id获取ProjectGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (pgService *ProjectGroupService) GetProjectGroup(id uint) (pg es.ProjectGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&pg).Error
	return
}

// GetProjectGroupInfoList 分页获取ProjectGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (pgService *ProjectGroupService) GetProjectGroupInfoList(info esReq.ProjectGroupSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&es.ProjectGroup{})
	var pgs []es.ProjectGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&pgs).Error
	return pgs, total, err
}

// GetProjectGroupInfoList 分页获取ProjectGroup记录
// Author [piexlmax](https://github.com/piexlmax)
func (pgService *ProjectGroupService) GetProjectGroupTreeList(info esReq.ProjectGroupSearch) (list interface{}, total int64, err error) {
	// limit := info.PageSize
	// offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&es.ProjectGroup{})
	var pgs []es.ProjectGroup
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Find(&pgs).Error
	for k, v := range pgs {
		if v.ParentId == 0 {
			pgService.ChangeProjectGroupToTree(&pgs[k], &pgs)
		}
	}
	pgs = utils.SliceFilter(pgs, func(pg es.ProjectGroup) bool { return pg.ParentId == 0 })
	total = int64(len(pgs))
	// if offset+limit > len(pgs) {
	// 	pgs = pgs[offset:]
	// } else {
	// 	pgs = pgs[offset : offset+limit]
	// }
	return pgs, total, err
}

//将GroupList转换为GroupTree
func (pgService *ProjectGroupService) ChangeProjectGroupToTree(projectGroup *es.ProjectGroup, projectGroupList *[]es.ProjectGroup) {
	for _, v := range *projectGroupList {
		if v.ParentId == projectGroup.ID {
			projectGroup.Children = append(projectGroup.Children, v)
			pgService.ChangeProjectGroupToTree(&projectGroup.Children[len(projectGroup.Children)-1], projectGroupList)
		}
	}
	return
}
