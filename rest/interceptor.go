package rest

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	log "mailer/logger"
	"mailer/model"
	"net/http"
	"os"
	"strings"
)

type apiKey struct {
	Key string `json:"api_key"`
}

const authFile string = "./auth.csv"

// AuthMiddleware - interceptor for users authentication
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.RequestURI, "documentation") {
			next.ServeHTTP(w, r)
		} else {
			apiKey := r.Header.Get("Authorization")
			if !checkAPIKey(apiKey, r) {
				// Cannot extract token
				WriteResponse(w, http.StatusForbidden, model.GlobalResponse{Errors: "Forbidden"})
			} else {
				// Proceed
				next.ServeHTTP(w, r)
			}
		}
	})
}

// CommonMiddleware - common functions for default request processing
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Got request at " + r.RequestURI)
		defer log.Info("Sent response from " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func checkAPIKey(key string, r *http.Request) bool {
	result := false

	csvFile, err := os.Open(authFile)
	if err != nil {
		log.Error("Cannot open auth file")
		os.Exit(1)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Error("Cannot read auth file")
		os.Exit(1)
	}

	h := sha256.New()
	h.Write([]byte(key))
	keyHash := hex.EncodeToString(h.Sum(nil))

	for _, row := range csvData {
		dbKey := row[1]
		if dbKey == keyHash {
			result = true
			// Save username for future usage
			r.Header.Set("UserName", row[0])
			break
		}
	}

	return result
}
