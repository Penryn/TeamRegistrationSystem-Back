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
	e.From = " <"+femail+">"
	e.To = []string{mail}
	e.Subject = "验证码发送"
	e.HTML = []byte(`<div>
	<div>
		尊敬的艾斯比，您好！
	</div>
	<div style="padding: 8px 40px 8px 50px;">
		<p>你本次的验证码为`+code+`,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
	</div>
	<div>
		<p>此邮箱为系统邮箱，请勿回复。</p>
	</div>    
</div>`)
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