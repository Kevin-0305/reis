// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package alert

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/ysicing/workwxbot"
	"go.uber.org/zap"
)

// SendWorkWechat 发送微信企业应用消息
func SendWorkWechat(touser, toparty, totag, msg string) string {
	open := global.GVA_CONFIG.Alert.WorkWeChat.Open
	if open != "1" {
		global.GVA_LOG.Error("微信企业应用未开启")
		return "微信企业应用未开启,请设置WorkWeChat的open为1"
	}
	cropid := global.GVA_CONFIG.Alert.WorkWeChat.CropID
	agentid := global.GVA_CONFIG.Alert.WorkWeChat.AgentID
	agentsecret := global.GVA_CONFIG.Alert.WorkWeChat.Secret

	//touser := beego.AppConfig.String("WorkWechat_ToUser")
	//toparty := beego.AppConfig.String("WorkWechat_ToParty")
	//totag := beego.AppConfig.String("WorkWechat_ToTag")

	workwxapi := workwxbot.Client{
		CropID:      cropid,
		AgentID:     agentid,
		AgentSecret: agentsecret,
	}
	workwxmsg := workwxbot.Message{
		ToUser:   touser,
		ToParty:  toparty,
		ToTag:    totag,
		MsgType:  "markdown",
		Markdown: workwxbot.Content{Content: msg},
	}
	if err := workwxapi.Send(workwxmsg); err != nil {
		global.GVA_LOG.Error("发送微信企业应用消息失败", zap.Error(err))
		return "发送微信企业应用消息失败"
	}

	return "workwechat send ok"
}
