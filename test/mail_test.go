package test

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <tanyuyan19@163.com>"
	e.To = []string{"1505985616@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为：<h1>123456</h1>")
	err := e.SendWithStartTLS("smtp.163.com:25", smtp.PlainAuth("", "tanyuyan19@163.com", define.MailPassWord, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
