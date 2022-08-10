package es

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/es/model"
)

type EsInterface interface {
	SnapshotRepositoryList(esSnapshotInfo *model.EsSnapshotInfo) (res map[string]interface{}, err error)
	SnapshotCreateRepository(snapshotCreateRepository *model.SnapshotCreateRepository) (res interface{}, err error)
	CleanupeRepository(repository *model.CleanupeRepository) (res interface{}, err error)
	SnapshotDeleteRepository(repository *model.SnapshotDeleteRepository) (res interface{}, err error)
	CreateSnapshot(snapshot *model.CreateSnapshot) (res interface{}, err error)
	SnapshotList(list *model.SnapshotList) (res interface{}, err error)
	SnapshotDelete(snapshotDelete *model.SnapshotDelete) (res interface{}, err error)
	SnapshotDetail(detail *model.SnapshotDetail) (res interface{}, err error)
	SnapshotRestore(restore *model.SnapshotRestore) (res interface{}, err error)
	SnapshotStatus(status *model.SnapshotStatus) (res interface{}, err error)
	Cat(rest *model.EsCat) (res interface{}, err error)
	// RunDsl(optimize *model.EsRest) (res interface{}, err error)
	Optimize(optimize *model.EsOptimize) (res interface{}, err error)
	RecoverCanWrite() (res interface{}, err error)
	EsDocDeleteRowByID(id *model.EsDocDeleteRowByID) (res interface{}, err error)
	EsDocUpdateByID(id *model.EsDocUpdateByID) (res interface{}, err error)
	EsDocInsert(id *model.EsDocUpdateByID) (res interface{}, err error)
	EsIndexCreate(info *model.EsIndexInfo) (res interface{}, err error)
	EsIndexDelete(info *model.EsIndexInfo) (res interface{}, err error)
	EsIndexGetSettings(info *model.EsIndexInfo) (res interface{}, err error)
	EsIndexGetSettingsInfo(info *model.EsIndexInfo) (res interface{}, err error)
	EsIndexGetAlias(info *model.EsAliasInfo) (res []map[string]interface{}, err error)
	EsIndexOperateAlias(info *model.EsAliasInfo) (res interface{}, err error)
	EsIndexReindex(info *model.EsReIndexInfo) (res interface{}, err error)
	EsIndexIndexNames() (res []string, err error)
	EsIndexStats(info *model.EsIndexInfo) (res interface{}, err error)
	EsIndexCatStatus(info *model.EsIndexInfo) (res interface{}, err error)
	EsMappingList(properties *model.EsMapGetProperties) (res map[string]interface{}, ver int64, err error)
	UpdateMapping(mapping *model.UpdateMapping) (res interface{}, err error)
	TaskList() (res interface{}, err error)
	Cancel(task *model.CancelTask) (res interface{}, err error)
	CrudGetList(filter *model.CrudFilter) (res interface{}, count int64, err error)
	EsDocSearch(esDocSearch *model.EsDocSearch) (res interface{}, count int64, err error)
	// CrudGetDSL(filter *model.CrudFilter) (res interface{},err error)
	// CrudDownload(filter *model.CrudFilter) (res interface{}, err error)
}

var VerError = errors.New("ES版本暂只支持6,7")

var EsServiceMap = map[int]func(conn *model.EsConnect) (EsInterface, error){
	//6: NewEsServiceV6,
	7: NewEsServiceV7,
	//8: NewEsServiceV8,
}

func NewEsService(conn *model.EsConnect) (EsInterface, error) {
	var found bool
	var fn func(conn *model.EsConnect) (EsInterface, error)
	if fn, found = EsServiceMap[conn.Version]; !found {
		return nil, VerError
	}
	fn = EsServiceMap[conn.Version]
	return fn(conn)
}
