package alert

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type AlertService struct {
}

func (s *AlertService) AlertTest(alertType string) (string, error) {
	ret := ""
	if alertType == "fsapp" {
		fstext := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = PostToFeiShuApp("PrometheusAlert", fstext, global.GVA_CONFIG.Alert.FsApp.AcceptIds)
	} else if alertType == "fs" {
		fstext := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = PostToFS("PrometheusAlert", fstext, global.GVA_CONFIG.Alert.Fs.WebHookUrl, "")
	} else if alertType == "wechat" {
		text := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = PostToWeChat("PrometheusAlert", global.GVA_CONFIG.Alert.Fs.WebHookUrl, text)
	} else if alertType == "dingding" {
		text := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = PostToDingDing("PrometheusAlert", global.GVA_CONFIG.Alert.DingDing.WebHookUrl, text, "")
	} else if alertType == "work-wechat" {
		text := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = SendWorkWechat(global.GVA_CONFIG.Alert.WorkWeChat.ToUser, global.GVA_CONFIG.Alert.WorkWeChat.ToParty, global.GVA_CONFIG.Alert.WorkWeChat.ToTag, text)
	} else if alertType == "tencent-message" {
		text := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = PostTencentMessage(text, "")
	} else if alertType == "tencent-phone" {
		text := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = PostTencentPhone(text, "")
	} else if alertType == "huawei-message" {
		text := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = PostHuaWeiMessage(text, "")
	} else if alertType == "aliyun-message" {
		text := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = PostAliyunMessage(text, "")
	} else if alertType == "aliyun-phone" {
		text := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = PostAliyunPhone(text, "")
	} else if alertType == "email" {
		text := "[告警消息]\n\n" + "测试告警\n\n" + "告警级别：测试\n\nPrometheusAlert\n\n" + "![PrometheusAlert](" + "11111" + ")"
		ret = SendEmail(text, "")
	} else {
		return "", fmt.Errorf("alert type not found")
	}
	return ret, nil

}

func (s *AlertService) GetAlertTypes() (result []map[string]string, err error) {
	fmt.Println(global.GVA_CONFIG.Alert)
	if global.GVA_CONFIG.Alert.FsApp.Open == "1" {
		result = append(result, map[string]string{"id": "1", "name": "fsapp", "title": "飞书机器人告警"})
	}
	if global.GVA_CONFIG.Alert.Fs.Open == "1" {
		result = append(result, map[string]string{"id": "2", "name": "fs", "title": "飞书告警"})
	}
	if global.GVA_CONFIG.Alert.WeChat.Open == "1" {
		result = append(result, map[string]string{"id": "3", "name": "wechat", "title": "微信告警"})
	}
	if global.GVA_CONFIG.Alert.DingDing.Open == "1" {
		result = append(result, map[string]string{"id": "4", "name": "dingding", "title": "钉钉告警"})
	}
	if global.GVA_CONFIG.Alert.WorkWeChat.Open == "1" {
		result = append(result, map[string]string{"id": "5", "name": "work-wechat", "title": "企业微信告警"})
	}
	if global.GVA_CONFIG.Alert.TencentMessage.Open == "1" {
		result = append(result, map[string]string{"id": "6", "name": "tencent-message", "title": "腾讯云消息告警"})
	}
	if global.GVA_CONFIG.Alert.TencentPhone.Open == "1" {
		result = append(result, map[string]string{"id": "7", "name": "tencent-phone", "title": "腾讯云电话告警"})
	}
	if global.GVA_CONFIG.Alert.HuaWeiMessage.Open == "1" {
		result = append(result, map[string]string{"id": "8", "name": "huwei-message", "title": "华为云消息告警"})
	}
	if global.GVA_CONFIG.Alert.AliyunMessage.Open == "1" {
		result = append(result, map[string]string{"id": "9", "name": "aliyun-message", "title": "阿里云消息告警"})
	}
	if global.GVA_CONFIG.Alert.AliyunPhone.Open == "1" {
		result = append(result, map[string]string{"id": "10", "name": "aliyun-phone", "title": "阿里云电话告警"})
	}
	if global.GVA_CONFIG.Alert.AlertEmail.Open == "1" {
		result = append(result, map[string]string{"id": "11", "name": "alert-email", "title": "邮件告警"})
	}
	return result, nil
}
