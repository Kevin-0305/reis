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

type DDMessage struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	} `json:"at"`
}

func PostToDingDing(title, text, Ddurl, AtSomeOne string) string {
	open := global.GVA_CONFIG.Alert.DingDing.Open
	if open != "1" {
		global.GVA_LOG.Info("[dingding] 钉钉发送功能已经关闭")
		return "钉钉接口未配置未开启状态,请先配置open-dingding为1"
	}
	atMobile := []string{"15888888888"}
	SendText := text
	if AtSomeOne != "" {
		atMobile = strings.Split(AtSomeOne, ",")
		AtText := ""
		for _, phoneN := range atMobile {
			AtText += " @" + phoneN
		}
		SendText += AtText
	}
	u := DDMessage{
		Msgtype: "markdown",
		Markdown: struct {
			Title string `json:"title"`
			Text  string `json:"text"`
		}{Title: title, Text: SendText},
		At: struct {
			AtMobiles []string `json:"atMobiles"`
			IsAtAll   bool     `json:"isAtAll"`
		}{AtMobiles: atMobile, IsAtAll: false},
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	client := &http.Client{}
	res, err := client.Post(Ddurl, "application/json", b)
	if err != nil {
		global.GVA_LOG.Error("[dingding] 发送到钉钉错误", zap.Error(err))
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		global.GVA_LOG.Error("[dingding] 发送到钉钉错误", zap.Error(err))
	}
	return string(result)
}
