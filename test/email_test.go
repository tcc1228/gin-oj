package test

import (
	"crypto/tls"
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <t1269324450@163.com>"
	e.To = []string{"1269324450@qq.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码：<b>123456</b>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "t1269324450@163.com", "BBPHSMDGDDUXCFRY", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Error(err)
	}
}
