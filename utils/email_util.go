package utils

import (
	"gopkg.in/gomail.v2"
	"miconvert-go/setting"
)

//
// SendMail
//  @Description: 发送邮件，发送文本，无附件
//  @param mailTo 目标邮箱
//  @param subject 主题
//  @param body 正文
//  @return error
//
func SendMail(mailTo []string, subject string, body string) error {
	m := gomail.NewMessage()
	emailConfig := setting.Conf.EmailConfig
	m.SetHeader("From", m.FormatAddress(emailConfig.User, "miconvert"))
	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//发送邮件
	d := gomail.NewDialer(emailConfig.Host,
		emailConfig.Port,
		emailConfig.User,
		emailConfig.Password)
	err := d.DialAndSend(m)
	return err
}
