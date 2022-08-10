package alert

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dyvmsapi"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"

	"strings"
)

func PostAliyunMessage(Messages, PhoneNumbers string) string {
	open := global.GVA_CONFIG.Alert.AliyunMessage.Open
	if open != "1" {
		global.GVA_LOG.Info("阿里云短信接口未配置未开启状态,请先配置aliyunmessage为1")
		return "阿里云短信接口未配置未开启状态,请先配置aliyunmessage为1"
	}
	AccessKeyId := global.GVA_CONFIG.Alert.AliyunMessage.AccessKeyID
	AccessSecret := global.GVA_CONFIG.Alert.AliyunMessage.AccessKeySecret
	SignName := global.GVA_CONFIG.Alert.AliyunMessage.SignName
	Template := global.GVA_CONFIG.Alert.AliyunMessage.TemplateCode
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", AccessKeyId, AccessSecret)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = PhoneNumbers
	request.SignName = SignName
	request.TemplateCode = Template
	request.TemplateParam = `{"code":"` + Messages + `"}`
	response, err := client.SendSms(request)

	if err != nil {
		global.GVA_LOG.Error("[aliyunmessage]", zap.Error(err))
	}
	return response.Message
}
func PostAliyunPhone(Messages string, PhoneNumbers string) string {
	open := global.GVA_CONFIG.Alert.AliyunPhone.Open
	if open != "1" {
		global.GVA_LOG.Info("阿里云电话接口未配置未开启状态,请先配置aliyunphone为1")
		return "阿里云电话接口未配置未开启状态,请先配置aliyunphone为1"
	}
	AccessKeyId := global.GVA_CONFIG.Alert.AliyunPhone.AccessKeyID
	AccessSecret := global.GVA_CONFIG.Alert.AliyunPhone.AccessKeySecret
	CalledShowNumber := global.GVA_CONFIG.Alert.AliyunPhone.ShowNumber
	TtsCode := global.GVA_CONFIG.Alert.AliyunPhone.TemplateCode

	mobiles := strings.Split(PhoneNumbers, ",")
	for _, m := range mobiles {
		client, err := dyvmsapi.NewClientWithAccessKey("cn-hangzhou", AccessKeyId, AccessSecret)
		request := dyvmsapi.CreateSingleCallByTtsRequest()
		request.Scheme = "https"
		request.CalledShowNumber = CalledShowNumber
		request.CalledNumber = m
		request.TtsCode = TtsCode
		request.TtsParam = `{"code":"` + Messages + `"}`
		request.PlayTimes = requests.NewInteger(2)

		response, err := client.SingleCallByTts(request)
		if err != nil {
			global.GVA_LOG.Error("[aliyunphone]", zap.Error(err))
		}
		global.GVA_LOG.Info("[aliyunphone]", zap.Any("response", response))
	}
	return PhoneNumbers + "Called Over."
}
