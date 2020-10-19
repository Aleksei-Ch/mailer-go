package email

import (
	"encoding/base64"
	"fmt"
	"mailer/config"
	log "mailer/logger"
	"mailer/model"
	"net/smtp"
	"strings"
)

var delimeter string = "=====CONTENT=SEPARATOR======"

// SendEmail function
func SendEmail(email *model.Email) {
	conf := config.Init().Mail

	// Choose auth method and set it up
	var auth smtp.Auth = nil
	if conf.User != "" {
		auth = smtp.PlainAuth("", conf.User, conf.Password, conf.Server)
	}

	msg := fmt.Sprintf("From: %s\r\n", email.From)
	msg += fmt.Sprintf("To: %s\r\n", email.To)
	if email.CC != "" {
		msg += fmt.Sprintf("Cc: %s\r\n", email.CC)
	}
	if email.BCC != "" {
		msg += fmt.Sprintf("Bcc: %s\r\n", email.BCC)
	}
	msg += fmt.Sprintf("Subject: %s\r\n", email.Title)

	msg += "MIME-Version: 1.0\r\n"
	msg += fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", delimeter)

	msg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
	msg += "Content-Type: text/plain; charset=\"UTF-8\"\r\n"
	msg += "Content-Transfer-Encoding: base64\r\n"
	msg += "\r\n" + base64.StdEncoding.EncodeToString([]byte(email.Body))

	if len(email.Attachments) > 0 {
		for name, body := range email.Attachments {
			msg += fmt.Sprintf("\r\n--%s\r\n", delimeter)
			msg += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
			msg += "Content-Transfer-Encoding: base64\r\n"
			msg += "Content-Disposition: attachment;filename=\"" + name + "\"\r\n"
			msg += "\r\n" + body
		}
	}

	// Send async
	go func() {
		if auth == nil {
			// Send without authentication
			client, err := smtp.Dial(fmt.Sprintf("%s:%v", conf.Server, conf.Port))
			if err != nil {
				log.Error(err)
				return
			}
			defer client.Quit()

			if err := client.Hello(conf.LocalServer); err != nil {
				log.Error(err)
				return
			}
			if err = client.Mail(email.From); err != nil {
				log.Error(err)
				return
			}
			for _, addr := range strings.Split(email.To, ";") {
				if err = client.Rcpt(addr); err != nil {
					log.Error(err)
					return
				}
			}

			writer, err := client.Data()
			if err != nil {
				log.Error(err)
				return
			}
			_, err = writer.Write([]byte(msg))
			if err != nil {
				log.Error(err)
				return
			}
			err = writer.Close()
			if err != nil {
				log.Error(err)
				return
			}
		} else {
			// Send with authentication
			err := smtp.SendMail(fmt.Sprintf("%s:%v", conf.Server, conf.Port), auth, conf.User, strings.Split(email.To, ";"), []byte(msg))
			if err != nil {
				log.Error(err)
			}
		}
	}()
}
