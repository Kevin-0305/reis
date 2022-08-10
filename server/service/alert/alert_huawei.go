package alert

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

//华为云短信子程序
func PostHuaWeiMessage(Messages string, PhoneNumbers string) string {
	open := global.GVA_CONFIG.Alert.HuaWeiMessage.Open
	if open != "1" {
		global.GVA_LOG.Info("[HuaWei] 华为云短信发送功能已经关闭 is closed")
		return "华为云短信接口未配置未开启状态,请先配置HuaWei为1"
	}
	hwappkey := global.GVA_CONFIG.Alert.HuaWeiMessage.AppKey
	hwappsecret := global.GVA_CONFIG.Alert.HuaWeiMessage.AppSecret
	hwappurl := global.GVA_CONFIG.Alert.HuaWeiMessage.AppUrl
	hwtplid := global.GVA_CONFIG.Alert.HuaWeiMessage.TemplateID
	hwsign := global.GVA_CONFIG.Alert.HuaWeiMessage.Signature
	sender := global.GVA_CONFIG.Alert.HuaWeiMessage.Sender
	//mobile格式:"15888888888,16666666666"
	//生成header
	now := time.Now().Format("2006-01-02T15:04:05Z")
	nonce := "7226249334"
	digest := utils.GetSha256Code(nonce + now + hwappsecret)
	digestBase64 := base64.StdEncoding.EncodeToString([]byte(digest))
	xheader := `"UsernameToken Username="` + hwappkey + `",PasswordDigest="` + digestBase64 + `",Nonce="` + nonce + `",Created="` + now + `"`
	//生成form数据
	FormData := strings.NewReader(url.Values{"from": {sender}, "to": {PhoneNumbers}, "templateId": {hwtplid}, "templateParas": {"[\"" + Messages + "\"]"}, "signature": {hwsign}, "statusCallback": {""}, "extend": {""}}.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("POST", hwappurl+"/sms/batchSendSms/v1", FormData)
	req.Header.Set("Authorization", `WSSE realm="SDP",profile="UsernameToken",type="Appkey"`)
	req.Header.Set("X-WSSE", xheader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		global.GVA_LOG.Error("[HuaWei] 发送到华为云短信错误", zap.Error(err))
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.GVA_LOG.Error("[HuaWei] 读取华为云短信返回数据错误", zap.Error(err))
	}
	return string(result)
}
