package sms

import (
	"errors"
	"mailer/config"
	log "mailer/logger"
	"mailer/model"
	"strconv"
	"time"

	"mailer/soap"
)

var (
	c    *soap.Client
	s    SmsServiceSoap
	conf model.Config
)

func init() {
	conf = config.Init()
	c = soap.NewClient("https://ws-wsdl.devinotele.com/smsservice.asmx",
		soap.WithTimeout(time.Minute),
	)
	s = NewSmsServiceSoap(c)
}

// SendSms via sms provider
func SendSms(sms *model.Sms) (string, error) {
	session, err := s.GetSessionID(&GetSessionID{Login: conf.Sms.User, Password: conf.Sms.Password})
	if err != nil {
		log.Error(err)
		return "-1", err
	}
	response, err := s.SendMessageByTimeZone(&SendMessageByTimeZone{
		SessionID:          session.GetSessionIDResult,
		SourceAddress:      conf.Sms.From,
		DestinationAddress: sms.To,
		Data:               sms.Body,
		SendDate:           time.Now(),
		Validity:           1,
	})
	if err != nil {
		log.Error(err)
		return "-1", err
	}
	if response.SendMessageByTimeZoneResult == nil && len(response.SendMessageByTimeZoneResult.Astring) == 0 {
		log.Error("Response is empty")
		return "-1", errors.New("Response is empty")
	}

	return response.SendMessageByTimeZoneResult.Astring[0], nil
}

// GetStatus of the sms via sms provider
func GetStatus(id string) (string, error) {
	session, err := s.GetSessionID(&GetSessionID{Login: conf.Sms.User, Password: conf.Sms.Password})
	if err != nil {
		log.Error(err)
		return "-1", err
	}
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Error(err)
		return "-1", err
	}
	response, err := s.GetMessageState(&GetMessageState{
		SessionID: session.GetSessionIDResult,
		MessageID: idNum,
	})
	if err != nil {
		log.Error(err)
		return "-1", err
	}
	if response.GetMessageStateResult == nil || response.GetMessageStateResult.StateDescription == "" {
		log.Error("Response is empty")
		return "-1", errors.New("Response is empty")
	}

	return response.GetMessageStateResult.StateDescription, nil
}
