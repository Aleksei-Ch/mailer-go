package model

import (
	"regexp"
	"strings"
	"time"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var phoneRegexp = regexp.MustCompile("^\\d{10,11}$")

// Email model
type Email struct {
	From        string            `json:"from" example:"user@server.ru"`
	To          string            `json:"to" example:"recipient1@server.ru;recipient2@server.ru"`
	CC          string            `json:"cc" example:"copy_user@server.ru;copy_user2@server.ru"`
	BCC         string            `json:"bcc" example:"secret_user@server.ru;secret_user2@server.ru"`
	Title       string            `json:"title" example:"Hi there!"`
	Body        string            `json:"body" example:"This is a body of the message.\n\nBest regards."`
	Attachments map[string]string `json:"attachments"`
}

// Validate an email
func (email *Email) Validate() map[string]string {
	errors := make(map[string]string)

	if !emailRegexp.Match([]byte(email.From)) {
		errors["From"] = "Not a valid email address"
	}
	email.To = strings.ReplaceAll(email.To, ",", ";")
	if email.To == "" || !checkEmailList(email.To) {
		errors["To"] = "Not a valid email list. Use list of emails separated by ; or ,"
	}
	email.CC = strings.ReplaceAll(email.CC, ",", ";")
	if email.CC != "" && !checkEmailList(email.CC) {
		errors["CC"] = "Not a valid email list. Use list of emails separated by ; or ,"
	}
	email.BCC = strings.ReplaceAll(email.BCC, ",", ";")
	if email.BCC != "" && !checkEmailList(email.BCC) {
		errors["BCC"] = "Not a valid email list. Use list of emails separated by ; or ,"
	}
	if email.Title == "" {
		errors["Title"] = "Field must not be empty"
	}
	if email.Body == "" {
		errors["Body"] = "Field must not be empty"
	}

	return errors
}

func checkEmailList(value string) bool {
	result := true
	if value != "" {
		for _, email := range strings.Split(strings.ReplaceAll(value, ",", ";"), ";") {
			if !emailRegexp.Match([]byte(email)) {
				result = false
				break
			}
		}
	}
	return result
}

// Sms model
type Sms struct {
	To   string `json:"to" example:"79770010101"`
	Body string `json:"body" example:"This is a body of the message.\n\nBest regards."`
}

// Validate an email
func (sms *Sms) Validate() map[string]string {
	errors := make(map[string]string)

	sms.To = strings.ReplaceAll(sms.To, "+7", "7")
	sms.To = strings.ReplaceAll(sms.To, " ", "")
	sms.To = strings.ReplaceAll(sms.To, "(", "")
	sms.To = strings.ReplaceAll(sms.To, ")", "")
	sms.To = strings.ReplaceAll(sms.To, "-", "")

	if sms.To == "" || !phoneRegexp.Match([]byte(sms.To)) {
		errors["To"] = "Not a valid phone number. Available patterns: 71112223344 or +71112223344"
	}
	if sms.Body == "" {
		errors["Body"] = "Field must not be empty"
	}

	return errors
}

// GlobalResponse struct
type GlobalResponse struct {
	Payload   interface{} `json:"payload,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
}
