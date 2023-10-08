package emailService

import (
	"TeamRegistrationSystem-Back/config/config"
	"crypto/tls"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

func MailSendCode(mail, code string) error {
	femail:=config.Config.GetString("email.fromemail")
	mpass:=config.Config.GetString("email.mailpassword")
	e := email.NewEmail()
	e.From = "艾斯比 <"+femail+">"
	e.To = []string{mail}
	e.Subject = "验证码发送"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", femail, mpass, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}
const (
    CodeLength = 6
    CodeChars  = "1234567890"
)

func RandCode() string {
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))

    code := make([]byte, CodeLength)
    for i := 0; i < CodeLength; i++ {
        code[i] = CodeChars[rng.Intn(len(CodeChars))]
    }

    return string(code)
}