package rest

import (
	"encoding/json"
	e "mailer/email"
	log "mailer/logger"
	"mailer/model"
	s "mailer/sms"
	"net/http"

	"github.com/gorilla/mux"
)

// SendEmail godoc
// @Summary Send an email
// @Description Send an email
// @Tags email
// @Accept  json
// @Produce  json
// @Param email body model.Email true "Email message"
// @Success 200 {object} model.GlobalResponse
// @Failure 403 {object} model.GlobalResponse
// @Failure 500 {object} model.GlobalResponse
// @Router /email [post]
// @Security ApiKeyAuth
func SendEmail(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("UserName")
	resp := model.GlobalResponse{}

	var email model.Email
	if errors := json.NewDecoder(r.Body).Decode(&email); errors != nil {
		resp.Errors = "Bad request"
		WriteResponse(w, http.StatusBadRequest, resp)
	} else if errors := email.Validate(); len(errors) > 0 {
		resp.Errors = errors
		WriteResponse(w, http.StatusBadRequest, resp)
	} else {
		e.SendEmail(&email)
		log.Debugf("User %s has sent an email %s", user, email.Title)
		resp.Payload = "Message sent"
		WriteResponse(w, http.StatusOK, resp)
	}
}

// SendSms godoc
// @Summary Send an sms
// @Description Send an sms. You will get sms id in response `payload`
// @Tags sms
// @Accept  json
// @Produce  json
// @Param sms body model.Sms true "Sms message"
// @Success 200 {object} model.GlobalResponse
// @Failure 403 {object} model.GlobalResponse
// @Failure 500 {object} model.GlobalResponse
// @Router /sms [post]
// @Security ApiKeyAuth
func SendSms(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("UserName")
	resp := model.GlobalResponse{}

	var sms model.Sms
	if errors := json.NewDecoder(r.Body).Decode(&sms); errors != nil {
		resp.Errors = "Bad request"
		WriteResponse(w, http.StatusBadRequest, resp)
	} else if errors := sms.Validate(); len(errors) > 0 {
		resp.Errors = errors
		WriteResponse(w, http.StatusBadRequest, resp)
	} else {
		id, err := s.SendSms(&sms)
		if err != nil {
			resp.Errors = err
			WriteResponse(w, http.StatusBadRequest, resp)
		} else {
			log.Debugf("User %s has sent an sms with id %v", user, id)
			resp.Payload = id
			WriteResponse(w, http.StatusOK, resp)
		}
	}
}

// GetSmsStatus godoc
// @Summary Get sms status
// @Description Get sms status. Result will be placed in response `payload`
// @Tags sms
// @Accept  json
// @Produce  json
// @Param id path string true "ID of the sms"
// @Success 200 {object} model.GlobalResponse
// @Failure 403 {object} model.GlobalResponse
// @Failure 500 {object} model.GlobalResponse
// @Router /smsStatus/{id} [get]
// @Security ApiKeyAuth
func GetSmsStatus(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("UserName")
	resp := model.GlobalResponse{}

	params := mux.Vars(r)
	smsID := params["id"]

	if smsID == "" {
		resp.Errors = "Incorrect sms id"
		WriteResponse(w, http.StatusBadRequest, resp)
	} else {
		status, err := s.GetStatus(smsID)
		if err != nil {
			resp.Errors = err
			WriteResponse(w, http.StatusBadRequest, resp)
		} else {
			log.Debugf("User %s has checked sms status of %v", user, status)
			resp.Payload = status
			WriteResponse(w, http.StatusOK, resp)
		}
	}
}
