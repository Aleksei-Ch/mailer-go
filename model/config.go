package model

// Config model
type Config struct {
	Log struct {
		LogDir          string `json:"logDir"`
		LogLevel        string `json:"logLevel"`
		LogRotationDays int64  `json:"logRotationDays"`
	} `json:"log"`
	Mail struct {
		LocalServer string `json:"localServer"`
		Server      string `json:"server"`
		Port        uint16 `json:"port"`
		User        string `json:"user"`
		Password    string `json:"password"`
	} `json:"mail"`
	Sms struct {
		From     string `json:"from"`
		User     string `json:"login"`
		Password string `json:"password"`
	} `json:"sms"`
	Proxy struct {
		Enabled bool   `json:"enabled"`
		Host    string `json:"host"`
		Port    int32  `json:"port"`
	} `json:"proxy"`
}
