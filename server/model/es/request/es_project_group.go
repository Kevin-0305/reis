package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/es"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProjectGroupSearch struct{
    es.ProjectGroup
    request.PageInfo
}
