package es_optimize

import (
	"context"
	"log"

	elasticV6 "github.com/olivere/elastic"
	elasticV7 "github.com/olivere/elastic/v7"

	"go.uber.org/zap"
)

// 合并索引
type Forcemerge struct {
	indexName []string
}

func (this *Forcemerge) SetIndexName(indexName string) {
	this.indexName = append(this.indexName, indexName)
}

func (this *Forcemerge) DoV6(client *elasticV6.Client) (err error) {
	//手动异步
	go func() {

		forcemergeRes, err := client.Forcemerge(this.indexName...).MaxNumSegments(1).Do(context.Background())
		if err != nil {
			log.Println("强制合并索引操作异常", forcemergeRes, err.Error())
		} else {
			log.Println("强制合并索引操作异常", forcemergeRes, err.Error())
		}
	}()
	return
}

func (this *Forcemerge) DoV7(client *elasticV7.Client) (err error) {
	//手动异步
	go func() {

		forcemergeRes, err := client.Forcemerge(this.indexName...).MaxNumSegments(1).Do(context.Background())
		if err != nil {
			log.Println("强制合并索引操作异常", zap.Reflect("forcemergeRes", forcemergeRes), zap.String("err.Error()", err.Error()))
		} else {
			log.Println("强制合并索引操作成功", zap.Reflect("forcemergeRes", forcemergeRes))
		}
	}()
	return
}

func newForcemerge() OptimizeInterface {
	return &Forcemerge{}
}
