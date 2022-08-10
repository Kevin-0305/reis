package es

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/es/es6_utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/es/es_optimize"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/es/es_settings"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/es/model"

	jsoniter "github.com/json-iterator/go"
	"github.com/olivere/elastic"
	elasticV7 "github.com/olivere/elastic/v7"
)

type EsServiceV7 struct {
	esClient *elasticV7.Client
}

func (this EsServiceV7) EsDocSearch(esDocSearch *model.EsDocSearch) (interface{}, int64, error) {
	search := this.esClient.Search(esDocSearch.IndexName).Type(esDocSearch.Type).Query(esDocSearch.Query)
	if esDocSearch.Sort != nil {
		search = search.SortBy(esDocSearch.Sort)
	}
	search.From(esDocSearch.Page).Size(esDocSearch.Limit)
	result, err := search.Do(context.Background())
	if err != nil {
		return nil, 0, err
	}
	count := result.Hits.TotalHits.Value
	return result, count, nil
}

func (this EsServiceV7) CrudGetList(crudFilter *model.CrudFilter) (res interface{}, count int64, err error) {
	q, err := es6_utils.GetWhereSql(crudFilter.Relation)
	if err != nil {
		return
	}
	search := this.esClient.Search(crudFilter.IndexName)
	q2 := search.Query(q)
	for _, tmp := range crudFilter.SortList {
		switch tmp.SortRule {
		case "desc":
			q2 = q2.Sort(tmp.Col, false)
		case "asc":
			q2 = q2.Sort(tmp.Col, true)
		}
	}

	result, err := q2.From(int(es6_utils.CreatePage(crudFilter.Page, crudFilter.Limit))).Size(crudFilter.Limit).Do(context.Background())
	if err != nil {
		return
	}
	count = result.Hits.TotalHits.Value
	return result, count, nil
}

func (this EsServiceV7) CreateSnapshot(createSnapshot *model.CreateSnapshot) (res interface{}, err error) {
	snapshotCreateService := this.esClient.
		SnapshotCreate(createSnapshot.RepositoryName, createSnapshot.SnapshotName)

	if createSnapshot.Wait != nil {
		snapshotCreateService.WaitForCompletion(*createSnapshot.Wait)
	}

	settings := model.Json{}

	if len(createSnapshot.IndexList) > 0 {
		settings["indices"] = strings.Join(createSnapshot.IndexList, ",")
	}

	if createSnapshot.IgnoreUnavailable != nil {
		settings["indices"] = *createSnapshot.IgnoreUnavailable
	}

	if createSnapshot.Partial != nil {
		settings["partial"] = *createSnapshot.Partial
	}
	if createSnapshot.IncludeGlobalState != nil {
		settings["include_global_state"] = *createSnapshot.IncludeGlobalState
	}

	res, err = snapshotCreateService.BodyJson(settings).Do(context.Background())

	if err != nil {
		return
	}

	return
}

func (this EsServiceV7) SnapshotList(snapshotList *model.SnapshotList) (res interface{}, err error) {
	if snapshotList.Repository == "" {
		return nil, errors.New("repository is empty")
	}

	res, err = this.esClient.PerformRequest(context.Background(), elasticV7.PerformRequestOptions{
		Method: "GET",
		Path:   fmt.Sprintf("/_cat/snapshots/%s", snapshotList.Repository),
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (this EsServiceV7) SnapshotDelete(snapshotDelete *model.SnapshotDelete) (res interface{}, err error) {
	_, err = this.esClient.
		SnapshotDelete(snapshotDelete.Repository, snapshotDelete.Snapshot).Do(context.Background())
	if err != nil {
		return
	}

	return
}

func (this EsServiceV7) SnapshotDetail(snapshotDetail *model.SnapshotDetail) (res interface{}, err error) {
	res, err = this.esClient.PerformRequest(context.Background(), elasticV7.PerformRequestOptions{
		Method: "GET",
		Path:   fmt.Sprintf("/_snapshot/%s/%s", snapshotDetail.Repository, snapshotDetail.Snapshot),
	})
	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) SnapshotRestore(snapshotRestore *model.SnapshotRestore) (res interface{}, err error) {

	snapshotRestoreService := this.esClient.SnapshotRestore(snapshotRestore.RepositoryName, snapshotRestore.SnapshotName)

	if snapshotRestore.Wait != nil {
		snapshotRestoreService = snapshotRestoreService.WaitForCompletion(*snapshotRestore.Wait)
	}

	if snapshotRestore.IgnoreUnavailable != nil {
		snapshotRestoreService = snapshotRestoreService.IgnoreUnavailable(*snapshotRestore.IgnoreUnavailable)
	}
	if len(snapshotRestore.IndexList) > 0 {
		snapshotRestoreService = snapshotRestoreService.Indices(snapshotRestore.IndexList...)
	}
	if snapshotRestore.Partial != nil {
		snapshotRestoreService = snapshotRestoreService.Partial(*snapshotRestore.Partial)
	}
	if snapshotRestore.IncludeGlobalState != nil {
		snapshotRestoreService = snapshotRestoreService.IncludeGlobalState(*snapshotRestore.IncludeGlobalState)
	}
	if snapshotRestore.RenamePattern != "" {
		snapshotRestoreService = snapshotRestoreService.RenamePattern(snapshotRestore.RenamePattern)
	}
	if snapshotRestore.RenameReplacement != "" {
		snapshotRestoreService = snapshotRestoreService.RenameReplacement(snapshotRestore.RenameReplacement)
	}

	res, err = snapshotRestoreService.Do(context.Background())

	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) SnapshotStatus(snapshotStatus *model.SnapshotStatus) (res interface{}, err error) {
	snapshotRestoreStatus := this.esClient.SnapshotStatus().Repository(snapshotStatus.RepositoryName).Snapshot(snapshotStatus.SnapshotName)

	res, err = snapshotRestoreStatus.Do(context.Background())

	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) Cat(esCat *model.EsCat) (data interface{}, err error) {
	switch esCat.Cat {
	case "CatHealth":
		data, err = this.esClient.CatHealth().Human(true).Do(context.Background())
	case "CatShards":
		data, err = this.esClient.CatShards().Human(true).Do(context.Background())
	case "CatCount":
		data, err = this.esClient.CatCount().Human(true).Do(context.Background())
	case "CatAllocation":
		data, err = this.esClient.CatAllocation().Human(true).Do(context.Background())
	case "CatAliases":
		data, err = this.esClient.CatAliases().Human(true).Do(context.Background())
	case "CatIndices":
		if esCat.IndexBytesFormat != "" {
			data, err = this.esClient.CatIndices().Sort("store.size:desc").Human(true).Bytes(esCat.IndexBytesFormat).Do(context.Background())
		} else {
			data, err = this.esClient.CatIndices().Sort("store.size:desc").Human(true).Do(context.Background())
		}
	case "CatSegments":
		data, err = this.esClient.IndexSegments().Human(true).Do(context.Background())
	case "CatStats":
		data, err = this.esClient.ClusterStats().Human(true).Do(context.Background())
	case "Node":
		parmas := url.Values{}
		parmas.Set("h", "ip,name,heap.percent,heap.current,heap.max,ram.percent,ram.current,ram.max,node.role,master,cpu,load_1m,load_5m,load_15m,disk.used_percent,disk.used,disk.total")
		res, err := this.esClient.PerformRequest(context.Background(), elasticV7.PerformRequestOptions{
			Method: "GET",
			Params: parmas,
			Path:   "/_cat/nodes",
		})
		if err != nil {
			return nil, err
		}
		data = res.Body
	}

	if err != nil {
		return
	}

	return
}

// func (this EsServiceV7) RunDsl(esRest EsRest) (res interface{}, err error) {
// 	esRest.Method = strings.ToUpper(esRest.Method)
// 	if esRest.Method == "GET" {
// 		c, err := jwt.ParseToken(ctx.Get("X-Token"))
// 		if err != nil {
// 			return err
// 		}

// 		gmDslHistoryModel := *model.GmDslHistoryModel{
// 			Uid:    int(c.ID),
// 			Method: esRest.Method,
// 			Path:   esRest.Path,
// 			Body:   esRest.Body,
// 		}

// 		err = gmDslHistory*model.Insert()

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	var performRequestOptions elasticV7.PerformRequestOptions

// 	if len(esRest.Path) > 0 {

// 		if esRest.Path[0:1] != "/" {
// 			esRest.Path = "/" + esRest.Path
// 		}

// 		u, err := url.Parse(esRest.Path)

// 		if err != nil {
// 			return err
// 		}
// 		path := strings.Split(esRest.Path, "?")[0]
// 		if len(strings.Split(esRest.Path, "/")) == 2 || strings.Contains(esRest.Path, "/_cat") {

// 			performRequestOptions = elasticV7.PerformRequestOptions{
// 				Method: esRest.Method,
// 				Path:   path,
// 				Body:   nil,
// 			}
// 			performRequestOptions.Params = u.Query()
// 		} else {
// 			performRequestOptions = elasticV7.PerformRequestOptions{
// 				Method: esRest.Method,
// 				Path:   path,
// 				Body:   esRest.Body,
// 			}
// 			performRequestOptions.Params = u.Query()
// 		}
// 	}

// 	res, err = this.esClient.PerformRequest(context.Background(), performRequestOptions)

// 	if err != nil {
// 		return err
// 	}

// 	if res.StatusCode != 200 && res.StatusCode != 201 {
// 		return this.Output(ctx, util.Map{
// 			"code": res.StatusCode,
// 			"msg":  fmt.Sprintf("请求异常! 错误码 :" + strconv.Itoa(res.StatusCode)),
// 			"data": res.Body,
// 		})
// 	}

// 	return
// }

func (this EsServiceV7) Optimize(esOptimize *model.EsOptimize) (res interface{}, err error) {
	optimize := es_optimize.OptimizeFactory(esOptimize.Command)

	if optimize == nil {
		return nil, errors.New("optimize command not found")

	}
	if esOptimize.IndexName != "" {
		optimize.SetIndexName(esOptimize.IndexName)
	}
	err = optimize.DoV7(this.esClient)
	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) RecoverCanWrite() (res interface{}, err error) {

	res, err = this.esClient.PerformRequest(context.Background(), elasticV7.PerformRequestOptions{
		Method: "PUT",
		Path:   "/_settings",
		Body: model.Json{
			"index": model.Json{
				"blocks": model.Json{
					"read_only_allow_delete": "false",
				},
			},
		},
	})

	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) EsDocDeleteRowByID(esDocDeleteRowByID *model.EsDocDeleteRowByID) (res interface{}, err error) {

	res, err = this.esClient.Delete().Index(esDocDeleteRowByID.IndexName).Id(esDocDeleteRowByID.ID).Do(context.Background())

	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) EsDocUpdateByID(esDocUpdateByID *model.EsDocUpdateByID) (res interface{}, err error) {
	res, err = this.esClient.Update().Index(esDocUpdateByID.Index).Type(esDocUpdateByID.Type).Id(esDocUpdateByID.ID).
		Doc(esDocUpdateByID.JSON).Do(context.Background())
	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) EsDocInsert(esDocUpdateByID *model.EsDocUpdateByID) (res interface{}, err error) {
	res, err = this.esClient.Index().
		Index(esDocUpdateByID.Index).
		Type(esDocUpdateByID.Type).BodyJson(esDocUpdateByID.JSON).Do(context.Background())
	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) EsIndexCreate(esIndexInfo *model.EsIndexInfo) (res interface{}, err error) {
	if esIndexInfo.IndexName == "" {
		return nil, errors.New("index name is empty")
	}

	if esIndexInfo.Types == "update" {
		res, err = this.esClient.IndexPutSettings().Index(esIndexInfo.IndexName).BodyJson(esIndexInfo.Settings).Do(context.TODO())
		if err != nil {
			return
		}

	} else {
		res, err = this.esClient.CreateIndex(esIndexInfo.IndexName).BodyJson(model.Json{
			"settings": esIndexInfo.Settings,
		}).Do(context.Background())
		if err != nil {
			return
		}
	}
	return
}

func (this EsServiceV7) EsIndexDelete(esIndexInfo *model.EsIndexInfo) (res interface{}, err error) {
	if esIndexInfo.IndexName == "" {
		return nil, errors.New("index name is empty")

	}
	_, err = this.esClient.DeleteIndex(strings.Split(esIndexInfo.IndexName, ",")...).Do(context.Background())
	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) EsIndexGetSettings(esIndexInfo *model.EsIndexInfo) (res interface{}, err error) {
	if esIndexInfo.IndexName == "" {
		return nil, errors.New("index name is empty")
	}

	res, err = this.esClient.IndexGetSettings(esIndexInfo.IndexName).Do(context.Background())
	if err != nil {
		return
	}

	return
}

func (this EsServiceV7) EsIndexGetSettingsInfo(esIndexInfo *model.EsIndexInfo) (res interface{}, err error) {
	if esIndexInfo.IndexName == "" {
		return nil, errors.New("index name is empty")
	}

	res, err = this.esClient.IndexGetSettings(esIndexInfo.IndexName).Do(context.Background())
	if err != nil {
		return
	}

	return
}

func (this EsServiceV7) EsIndexGetAlias(esAliasInfo *model.EsAliasInfo) (res []map[string]interface{}, err error) {
	if esAliasInfo.IndexName == "" {
		return nil, errors.New("index name is empty")
	}
	aliasRes, err := this.esClient.Aliases().Index(esAliasInfo.IndexName).Do(context.Background())
	result := aliasRes.Indices[esAliasInfo.IndexName].Aliases
	for _, v := range result {
		m := make(map[string]interface{})
		m["AliasName"] = v.AliasName
		m["IsWriteIndex"] = v.IsWriteIndex
		res = append(res, m)
	}
	return
}

func (this EsServiceV7) EsIndexOperateAlias(esAliasInfo *model.EsAliasInfo) (res interface{}, err error) {
	const Add = 1
	const Delete = 2
	const MoveToAnotherIndex = 3
	const PatchAdd = 4
	switch esAliasInfo.Types {
	case Add:
		if esAliasInfo.IndexName == "" {
			return nil, errors.New("index name is empty")
		}
		res, err = this.esClient.Alias().Add(esAliasInfo.IndexName, esAliasInfo.AliasName).Do(context.Background())
	case Delete:
		if esAliasInfo.IndexName == "" {
			return nil, errors.New("index name is empty")
		}
		res, err = this.esClient.Alias().Remove(esAliasInfo.IndexName, esAliasInfo.AliasName).Do(context.Background())
	case MoveToAnotherIndex:
		res, err = this.esClient.Alias().Action(elastic.NewAliasAddAction(esAliasInfo.AliasName).Index(esAliasInfo.NewIndexList...)).Do(context.Background())
	case PatchAdd:
		if esAliasInfo.IndexName == "" {
			return nil, errors.New("index name is empty")
		}
		wg := sync.WaitGroup{}
		NewAliasNameListLen := len(esAliasInfo.NewAliasNameList)
		if len(esAliasInfo.NewAliasNameList) > 10 {
			err = errors.New("别名列表数量不能大于10")
			break
		} else {
			wg.Add(NewAliasNameListLen)
			for _, aliasName := range esAliasInfo.NewAliasNameList {
				go func(aliasName string) {
					defer wg.Done()
					res, err = this.esClient.Alias().
						Add(esAliasInfo.IndexName, aliasName).
						Do(context.TODO())
				}(aliasName)
			}
			wg.Wait()
		}
	default:
		err = model.ReqParmasValid
	}

	if err != nil {
		return
	}

	return
}

func (this EsServiceV7) EsIndexReindex(esReIndexInfo *model.EsReIndexInfo) (res interface{}, err error) {
	reindex := this.esClient.Reindex()
	urlValues := esReIndexInfo.UrlValues
	if urlValues.WaitForActiveShards != "" {
		reindex = reindex.WaitForActiveShards(urlValues.WaitForActiveShards)
	}
	if urlValues.Slices != 0 {
		reindex = reindex.Slices(urlValues.Slices)
	}
	if urlValues.Refresh != "" {
		reindex = reindex.Refresh(urlValues.Refresh)
	}
	if urlValues.Timeout != "" {
		reindex = reindex.Timeout(urlValues.Refresh)
	}
	if urlValues.RequestsPerSecond != 0 {
		reindex = reindex.RequestsPerSecond(urlValues.RequestsPerSecond)
	}
	if urlValues.WaitForCompletion != nil {
		reindex = reindex.WaitForCompletion(*urlValues.WaitForCompletion)
	}

	res, err = reindex.Body(esReIndexInfo.Body).Do(context.Background())
	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) EsIndexIndexNames() (res []string, err error) {
	catIndicesResponse, err := this.esClient.CatIndices().Human(true).Do(context.Background())
	if err != nil {
		return
	}
	indexNames := []string{}

	for _, catIndices := range catIndicesResponse {
		indexNames = append(indexNames, catIndices.Index)
	}
	res = indexNames
	return
}

func (this EsServiceV7) EsIndexStats(esIndexInfo *model.EsIndexInfo) (res interface{}, err error) {
	if esIndexInfo.IndexName == "" {
		return nil, errors.New("index is empty")

	}

	res, err = this.esClient.IndexStats(esIndexInfo.IndexName).Do(context.Background())
	if err != nil {
		return
	}

	return
}

func (this EsServiceV7) EsIndexCatStatus(esIndexInfo *model.EsIndexInfo) (res interface{}, err error) {
	res, err = this.esClient.PerformRequest(context.Background(), elasticV7.PerformRequestOptions{
		Method: "GET",
		Path:   fmt.Sprintf("/_cat/indices/%s?h=status", esIndexInfo.IndexName),
	})
	if err != nil {
		return
	}

	return
}

func (this EsServiceV7) EsMappingList(esConnect *model.EsMapGetProperties) (res map[string]interface{}, ver int64, err error) {
	if esConnect.IndexName == "" {
		res, err = this.esClient.GetMapping().Do(context.Background())
		if err != nil {
			return
		}
		ver = 7
		return
	} else {
		res, err = this.esClient.GetMapping().Index(esConnect.IndexName).Do(context.Background())
		if err != nil {
			return
		}
		ver = 7
		return
	}
}

func (this EsServiceV7) UpdateMapping(updateMapping *model.UpdateMapping) (res interface{}, err error) {
	res, err = this.esClient.PutMapping().
		Index(updateMapping.IndexName).
		BodyJson(updateMapping.Properties).
		Do(context.Background())
	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) TaskList() (res interface{}, err error) {
	tasksListService := this.esClient.TasksList().Detailed(true)

	tasksListResponse, err := tasksListService.Do(context.Background())
	if err != nil {
		return
	}

	taskListRes := map[string]*elasticV7.TaskInfo{}

	for _, node := range tasksListResponse.Nodes {
		for taskId, taskInfo := range node.Tasks {
			taskListRes[taskId] = taskInfo
		}
	}

	return
}

func (this EsServiceV7) Cancel(cancelTask *model.CancelTask) (res interface{}, err error) {
	res, err = this.esClient.TasksCancel().TaskId(cancelTask.TaskID).Do(context.Background())
	if err != nil {
		return
	}
	return
}

func (this EsServiceV7) SnapshotRepositoryList(esSnapshotInfo *model.EsSnapshotInfo) (res map[string]interface{}, err error) {

	clusterSettings, err := es_settings.NewSettingsByV7(this.esClient)
	if err != nil {
		return
	}
	pathRepo := clusterSettings.GetPathRepo()

	if len(pathRepo) == 0 {
		return
	}

	result, err := this.esClient.SnapshotGetRepository(esSnapshotInfo.SnapshotInfoList...).Do(context.Background())
	if err != nil {
		return
	}

	type tmp struct {
		Name                   string `json:"name"`
		Type                   string `json:"type"`
		Location               string `json:"location"`
		Compress               string `json:"compress"`
		MaxRestoreBytesPerSec  string `json:"max_restore_bytes_per_sec"`
		MaxSnapshotBytesPerSec string `json:"max_snapshot_bytes_per_sec"`
		ChunkSize              string `json:"chunk_size"`
		Readonly               string `json:"readonly"`
	}
	list := []tmp{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	for name, settings := range result {
		var t tmp
		t.Type = settings.Type
		t.Name = name
		b, err := json.Marshal(settings.Settings)
		if err != nil {
			log.Panicln(err)
			continue
		}
		err = json.Unmarshal(b, &t)
		if err != nil {
			log.Panicln(err)
			continue
		}
		list = append(list, t)
	}
	res = map[string]interface{}{
		"list":     list,
		"res":      res,
		"pathRepo": pathRepo,
	}
	return
}

func (this EsServiceV7) SnapshotCreateRepository(snapshotCreateRepository *model.SnapshotCreateRepository) (res interface{}, err error) {

	clusterSettings, err := es_settings.NewSettingsByV7(this.esClient)
	if err != nil {
		return
	}
	pathRepo := clusterSettings.GetPathRepo()
	getAllowedUrls := clusterSettings.GetAllowedUrls()

	settings := make(map[string]interface{})

	if snapshotCreateRepository.Compress != "" {
		compress := snapshotCreateRepository.Compress
		settings["compress"] = compress
	}

	if snapshotCreateRepository.MaxRestoreBytesPerSec != "" {
		settings["max_restore_bytes_per_sec"] = snapshotCreateRepository.MaxRestoreBytesPerSec
	}

	if snapshotCreateRepository.MaxSnapshotBytesPerSec != "" {
		settings["max_snapshot_bytes_per_sec"] = snapshotCreateRepository.MaxSnapshotBytesPerSec
	}

	if snapshotCreateRepository.Readonly != "" {
		settings["readonly"] = snapshotCreateRepository.Readonly
	}

	if snapshotCreateRepository.ChunkSize != "" {
		settings["chunk_size"] = snapshotCreateRepository.ChunkSize
	}

	switch snapshotCreateRepository.Type {
	case "fs":
		if len(pathRepo) == 0 {
			return nil, errors.New("pathRepo is empty")
		}
		settings["location"] = snapshotCreateRepository.Location
	case "url":
		if len(getAllowedUrls) == 0 {
			errors.New("无效的type")
			return nil, errors.New("请先设置 allowed_urls")
		}
		settings["url"] = snapshotCreateRepository.Location
	default:
		return nil, errors.New("无效的type")
	}

	_, err = this.esClient.SnapshotCreateRepository(snapshotCreateRepository.Repository).Type(snapshotCreateRepository.Type).Settings(
		settings,
	).Do(context.Background())
	if err != nil {
		return
	}

	return
}

func (this EsServiceV7) CleanupeRepository(cleanupeRepository *model.CleanupeRepository) (res interface{}, err error) {
	res, err = this.esClient.PerformRequest(context.Background(), elasticV7.PerformRequestOptions{
		Method: "POST",
		Path:   fmt.Sprintf("/_snapshot/%s/_cleanup", cleanupeRepository.Repository),
	})
	if err != nil {
		return
	}

	return
}

func (this EsServiceV7) SnapshotDeleteRepository(repository *model.SnapshotDeleteRepository) (res interface{}, err error) {
	_, err = this.esClient.SnapshotDeleteRepository(repository.Repository).Do(context.Background())
	if err != nil {
		return
	}

	return
}

func NewEsServiceV7(connect *model.EsConnect) (service EsInterface, err error) {
	esClinet, err := model.NewEsClientV7(connect)

	if err != nil {
		return nil, err
	}

	return &EsServiceV7{esClient: esClinet}, nil
}

// func (this EsServiceV7) CrudGetDSL(crudFilter *model.CrudFilter) (res interface{}, err error) {
// 	q, err := es7_utils.GetWhereSql(crudFilter.Relation)
// 	if err != nil {
// 		return err
// 	}

// 	search := elasticV7.NewSearchSource()

// 	q2 := search.Query(q)
// 	for _, tmp := range crudFilter.SortList {
// 		switch tmp.SortRule {
// 		case "desc":
// 			q2 = q2.Sort(tmp.Col, false)
// 		case "asc":
// 			q2 = q2.Sort(tmp.Col, true)
// 		}
// 	}

// 	res, err = q2.From(int(db.CreatePage(crudFilter.Page, crudFilter.Limit))).Size(crudFilter.Limit).Source()
// 	if err != nil {
// 		return err
// 	}
// 	return this.Success(ctx, response.SearchSuccess, util.Map{"list": res})
// }

// func (this EsServiceV7) CrudDownload(filter *model.CrudFilter) (res *elasticV7.SearchResult, err error) {

// 	fields, err := this.esClient.GetMapping().Index(filter.IndexName).Do(context.Background())
// 	if err != nil {
// 		return
// 	}
// 	fieldsArr := []string{"_index", "_type", "_id"}
// 	data, ok := fields[filter.IndexName].(map[string]interface{})
// 	if !ok {
// 		return  nil, errors.New("无效的索引")
// 	}
// 	mappings, ok := data["mappings"].(map[string]interface{})
// 	if !ok {
// 		return nil, errors.New("该索引没有映射结构")
// 	}
// 	properties, ok := mappings["properties"].(map[string]interface{})
// 	if !ok {
// 		return nil, errors.New("该索引没有映射结构")
// 	}
// 	propertiesArr := []string{}
// 	for key := range properties {
// 		propertiesArr = append(propertiesArr, key)
// 	}
// 	sort.Strings(propertiesArr)
// 	fieldsArr = append(fieldsArr, propertiesArr...)
// 	q, err := es7_utils.GetWhereSql(filter.Relation)
// 	if err != nil {
// 		return
// 	}
// 	search := this.esClient.Search(filter.IndexName)
// 	res, err = search.Query(q).Sort("_id", false).Size(8000).Do(context.Background())
// 	if err != nil {
// 		return
// 	}

// 	lastIdArr := res.Hits.Hits[len(res.Hits.Hits)-1].Sort

// 	llist := [][]string{}
// 	var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 	flushHitsDataFn := func(hits []*elasticV7.SearchHit) {
// 		for _, data := range hits {
// 			list := []string{}
// 			list = append(list, data.Index, "_doc", data.Id)
// 			m := map[string]interface{}{}

// 			json.Unmarshal(data.Source, &m)

// 			for _, field := range fieldsArr {
// 				if field == "_index" || field == "_type" || field == "_id" {
// 					continue
// 				}
// 				if value, ok := m[field]; ok {
// 					list = append(list, utils.ToExcelData(value))
// 				} else {
// 					list = append(list, "")
// 				}
// 			}

// 			llist = append(llist, list)
// 		}
// 	}

// 	flushHitsDataFn(res.Hits.Hits)
// 	haveData := true
// 	for haveData {
// 		search := this.esClient.Search(filter.IndexName)
// 		res, err = search.Query(q).Sort("_id", false).Size(8000).SearchAfter(lastIdArr...).Do(context.Background())
// 		if err != nil {
// 			return
// 		}
// 		if len(res.Hits.Hits) == 0 {
// 			break
// 		}

// 		lastIdArr = res.Hits.Hits[len(res.Hits.Hits)-1].Sort
// 		flushHitsDataFn(res.Hits.Hits)
// 	}

// 	return this.DownloadExcel(
// 		"test",
// 		fieldsArr,
// 		llist, ctx)

// }
