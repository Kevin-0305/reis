package alert

import (
	"crypto/tls"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/go-gomail/gomail"
	"go.uber.org/zap"
)

// SendEmail
func SendEmail(EmailBody, Emails string) string {
	open := global.GVA_CONFIG.Alert.AlertEmail.Open
	if open != "1" {
		global.GVA_LOG.Info("[email] 邮件发送功能已经关闭 is closed")
		return "邮件发送功能已经关闭,请先配置email为1"
	}
	serverHost := global.GVA_CONFIG.Alert.AlertEmail.Host
	serverPort := global.GVA_CONFIG.Alert.AlertEmail.Port
	fromEmail := global.GVA_CONFIG.Alert.AlertEmail.FromEmail
	Passwd := global.GVA_CONFIG.Alert.AlertEmail.Password
	EmailTitle := global.GVA_CONFIG.Alert.AlertEmail.Title
	SendToEmails := []string{}
	m := gomail.NewMessage()
	if len(Emails) == 0 {
		return "收件人不能为空"
	}
	for _, Email := range strings.Split(Emails, ",") {
		SendToEmails = append(SendToEmails, strings.TrimSpace(Email))
	}
	// 收件人,...代表打散列表填充不定参数
	m.SetHeader("To", SendToEmails...)
	// 发件人
	m.SetAddressHeader("From", fromEmail, EmailTitle)
	// 主题
	m.SetHeader("Subject", EmailTitle)
	// 正文
	m.SetBody("text/html", EmailBody)
	d := gomail.NewDialer(serverHost, serverPort, fromEmail, Passwd)
	//忽略证书
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	//d.SSL=true
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		global.GVA_LOG.Error("[email] 发送邮件错误", zap.Error(err))
		return "发送邮件错误"
	}
	return "email send ok to " + Emails
}
