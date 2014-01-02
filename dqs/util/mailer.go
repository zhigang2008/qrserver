package util

import (
	"encoding/base64"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/jpoehls/gophermail"
	"net/smtp"
)

var (
	b64 = base64.StdEncoding
)

type Mailer struct {
	Host     string
	Port     string
	auth     smtp.Auth
	Mail     string
	User     string
	Password string
}

//初始化mailer
func InitMailer(host, port, mail string, auth bool, user, password string) *Mailer {
	mailer := Mailer{
		Host:     host,
		Port:     port,
		Mail:     mail,
		User:     user,
		Password: password}

	if auth {
		mailer.auth = smtp.PlainAuth(
			"",
			mailer.User,
			mailer.Password,
			mailer.Host)
	} else {
		mailer.auth = smtp.PlainAuth(
			"",
			"",
			"",
			mailer.Host)
	}
	return &mailer
}

//发送邮件
func (m *Mailer) SendMail(to []string, subject, body string) error {

	header := make(map[string]string)
	header["From"] = m.Mail
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"
	toaddr := ""
	for _, v := range to {
		toaddr += v + ";"
	}
	header["To"] = toaddr
	header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte(subject)))

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + b64.EncodeToString([]byte(body))

	err := smtp.SendMail(
		m.Host+":"+m.Port,
		m.auth,
		m.Mail,
		to,
		[]byte(message),
	)
	if err != nil {
		log.Errorf("发送失败:%s", err.Error())
		return err

	}
	return nil
}

//发送邮件
func SendMulityMail(host, port string, auth bool, user, password string, msg *gophermail.Message) error {
	addr := fmt.Sprintf("%s:%s", host, port)
	var mailAuth smtp.Auth
	if auth {
		mailAuth = smtp.PlainAuth(
			"",
			user,
			password,
			host)
	} else {
		mailAuth = smtp.PlainAuth(
			"",
			"",
			"",
			host)
	}
	return gophermail.SendMail(addr, mailAuth, msg)
}
