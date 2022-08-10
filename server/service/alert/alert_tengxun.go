package alert

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

//腾讯短信接口消息格式
type Mobiles struct {
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

type TXmessage struct {
	Ext    string    `json:"ext"`
	Extend string    `json:"extend"`
	Params []string  `json:"params"`
	Sig    string    `json:"sig"`
	Sign   string    `json:"sign"`
	Tel    []Mobiles `json:"tel"`
	Time   int       `json:"time"`
	Tpl_id int       `json:"tpl_id"`
}

//腾讯短信子程序
func PostTencentMessage(Messages string, PhoneNumbers string) string {
	open := global.GVA_CONFIG.Alert.TencentMessage.Open
	if open != "1" {
		global.GVA_LOG.Info("腾讯短信接口未配置未开启状态,请先配置tencentmessage为1")
		return "腾讯短信接口未配置未开启状态,请先配置tencentmessage为1"
	}
	strAppKey := global.GVA_CONFIG.Alert.TencentMessage.AppKey
	tpl_id := global.GVA_CONFIG.Alert.TencentMessage.TemplateID
	sdkappid := global.GVA_CONFIG.Alert.TencentMessage.SdkAppID
	sign := global.GVA_CONFIG.Alert.TencentMessage.Signature
	//腾讯短信接口算法部分
	//mobile格式:"15888888888,16666666666"
	TXmobile := Mobiles{}
	TXmobiles := []Mobiles{}
	mobiles := strings.Split(PhoneNumbers, ",")
	for _, m := range mobiles {
		TXmobile.Mobile = m
		TXmobile.Nationcode = "86"
		TXmobiles = append(TXmobiles, TXmobile)
	}
	strRand := "7226249334"
	strTime := strconv.FormatInt(time.Now().Unix(), 10)
	intTime, _ := strconv.Atoi(strTime)
	sig := utils.GetSha256Code("appkey=" + strAppKey + "&random=" + strRand + "&time=" + strTime + "&mobile=" + PhoneNumbers)
	TXurl := "https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2?sdkappid=" + sdkappid + "&random=" + strRand
	u := TXmessage{
		Ext:    "",
		Extend: "",
		Params: []string{Messages},
		Sig:    sig,
		Sign:   sign,
		Tel:    TXmobiles,
		Time:   intTime,
		Tpl_id: tpl_id,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	//res,err := http.Post(Ddurl, "application/json", b)
	//resp, err := http.PostForm(url,url.Values{"key": {"Value"}, "id": {"123"}})
	client := &http.Client{}
	res, err := client.Post(TXurl, "application/json", b)

	if err != nil {
		global.GVA_LOG.Error("腾讯短信接口发送失败", zap.Error(err))
	}

	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		global.GVA_LOG.Error("腾讯短信接口发送失败", zap.Error(err))
	}
	return string(result)
}

//腾讯语音提醒接口
type TXphonecall struct {
	Ext       string   `json:"ext"`
	Tpl_id    int      `json:"tpl_id"`
	Params    []string `json:"params"`
	Playtimes int      `json:"playtimes"`
	Sig       string   `json:"sig"`
	Tel       Mobiles  `json:"tel"`
	Time      int      `json:"time"`
}

//腾讯语音子程序
func PostTencentPhone(Messages string, PhoneNumbers string) string {
	open := global.GVA_CONFIG.Alert.TencentPhone.Open
	if open != "1" {
		global.GVA_LOG.Info("腾讯语音提醒接口未配置未开启状态,请先配置tencentphone为1")
		return "腾讯语音提醒接口未配置未开启状态,请先配置tencentphone为1"
	}
	strAppKey := global.GVA_CONFIG.Alert.TencentPhone.AppKey
	sdkappid := global.GVA_CONFIG.Alert.TencentPhone.SdkAppID
	tpl_id, _ := strconv.Atoi(global.GVA_CONFIG.Alert.TencentPhone.TemplateID)
	//腾讯短信接口算法部分
	TXmobile := Mobiles{}
	mobiles := strings.Split(PhoneNumbers, ",")
	for _, m := range mobiles {
		TXmobile.Mobile = m
		TXmobile.Nationcode = "86"
		strRand := "7226249334"
		strTime := strconv.FormatInt(time.Now().Unix(), 10)
		intTime, _ := strconv.Atoi(strTime)
		sig := utils.GetSha256Code("appkey=" + strAppKey + "&random=" + strRand + "&time=" + strTime + "&mobile=" + m)
		TXurl := "https://cloud.tim.qq.com/v5/tlsvoicesvr/sendtvoice?sdkappid=" + sdkappid + "&random=" + strRand
		u := TXphonecall{
			Ext:       "",
			Tpl_id:    tpl_id,
			Params:    []string{Messages},
			Playtimes: 2,
			Sig:       sig,
			Tel:       TXmobile,
			Time:      intTime,
		}
		res := PhoneCallPost(TXurl, u)
		global.GVA_LOG.Info("腾讯语音提醒接口发送结果", zap.Any("res", res))
	}
	return PhoneNumbers + " Called Over."
}

func PhoneCallPost(url string, u TXphonecall) string {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Post(url, "application/json", b)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}
	return string(result)
}
