package main

import (
	"mailer/rest"
	"net/http"

	log "mailer/logger"

	_ "mailer/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Mailer Rest API
// @version 1.0
// @description This is a mailer server that provides an API for sending emails and SMS.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /api/v1
func main() {
	log.Info("Mail server starting...")
	r := mux.NewRouter()

	// Common interceptor
	r.Use(rest.CommonMiddleware)
	// Authorization interceptor
	r.Use(rest.AuthMiddleware)

	r.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	// Mappings of routes
	r.HandleFunc("/api/v1/email", rest.SendEmail).Methods("POST")
	r.HandleFunc("/api/v1/sms", rest.SendSms).Methods("POST")
	r.HandleFunc("/api/v1/smsStatus/{id}", rest.GetSmsStatus).Methods("GET")

	log.Error(http.ListenAndServe(":8080", r))
}
