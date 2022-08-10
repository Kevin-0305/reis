package webhook

import (
	"fmt"
	"io/ioutil"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebHookRobotApi struct {
}

func (s *WebHookRobotApi) FeiShuApp(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		global.GVA_LOG.Error("err", zap.Error(err))
	}
	fmt.Println(string(body))
	c.Writer.Write(body)
}
