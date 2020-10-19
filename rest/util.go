package rest

import (
	"encoding/json"
	log "mailer/logger"
	"mailer/model"
	"net/http"
	"time"
)

// WriteResponse - common func for response writing
func WriteResponse(w http.ResponseWriter, httpStatus int, data model.GlobalResponse) {
	data.Timestamp = time.Now()
	js, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Error("Could not marshal http write data: ", data)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(js)
	return
}
