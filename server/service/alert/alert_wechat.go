package alert

import (
	"bytes"
	"encoding/json"

	"io/ioutil"
	"net/http"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

type Mark struct {
	Content string `json:"content"`
}
type WXMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown Mark   `json:"markdown"`
}

func PostToWeChat(text, WXurl, atuserid string) string {
	open := global.GVA_CONFIG.Alert.WeChat.Open
	if open != "1" {
		global.GVA_LOG.Info("[wechat] 企业微信发送功能已经关闭 is closed")
		return "微信接口未配置未开启状态,请先配置wechat为1"
	}

	SendContent := text
	if atuserid != "" {
		userid := strings.Split(atuserid, ",")
		idtext := ""
		for _, id := range userid {
			idtext += "<@" + id + ">"
		}
		SendContent += idtext
	}
	u := WXMessage{
		Msgtype:  "markdown",
		Markdown: Mark{Content: SendContent},
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	client := &http.Client{}
	res, err := client.Post(WXurl, "application/json", b)
	if err != nil {
		global.GVA_LOG.Error("[wechat] 发送到企业微信错误", zap.Error(err))
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		global.GVA_LOG.Error("[wechat] 发送到企业微信错误 error", zap.Error(err))
	}
	return string(result)
}
