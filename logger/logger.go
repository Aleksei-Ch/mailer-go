package logger

import (
	"fmt"
	"log"
	"mailer/config"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	errorLogger   *log.Logger
	warningLogger *log.Logger
	debugLogger   *log.Logger
	infoLogger    *log.Logger
)

// Init logger
func Init() {
	// Reload config
	conf := config.Init().Log

	// Log rotation
	err := filepath.Walk(conf.LogDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil || !strings.Contains(info.Name(), ".log") {
				return err
			}
			fDate, err := time.Parse("2006-01-02", strings.ReplaceAll(info.Name(), ".log", ""))
			if err == nil {
				if diff := time.Now().Sub(fDate); int64(diff) > conf.LogRotationDays*24*int64(time.Hour) {
					os.Remove(path)
				}
			}
			return nil
		})
	if err != nil {
		fmt.Println(err)
	}

	logFile := time.Now().Format("2006-01-02") + ".log"

	file, err := os.OpenFile(conf.LogDir+"/"+logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
	}

	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	warningLogger = nil
	if conf.LogLevel == "DEBUG" || conf.LogLevel == "INFO" || conf.LogLevel == "WARNING" {
		warningLogger = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	}
	debugLogger = nil
	if conf.LogLevel == "DEBUG" || conf.LogLevel == "INFO" {
		debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	}
	infoLogger = nil
	if conf.LogLevel == "INFO" {
		infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lmsgprefix)
	}
}

// Error - prints an error row
func Error(v ...interface{}) {
	Init()
	errorLogger.Println(v...)
}

// Errorf - prints an formatted error row
func Errorf(format string, v ...interface{}) {
	Error(fmt.Sprintf(format, v...))
}

// Warn - prints an error row
func Warn(v ...interface{}) {
	Init()
	if warningLogger != nil {
		warningLogger.Println(v...)
	}
}

// Warnf - prints an error row
func Warnf(format string, v ...interface{}) {
	Warn(fmt.Sprintf(format, v...))
}

// Info - prints an error row
func Info(v ...interface{}) {
	Init()
	if infoLogger != nil {
		infoLogger.Println(v...)
	}
}

// Infof - prints an error row
func Infof(format string, v ...interface{}) {
	Info(fmt.Sprintf(format, v...))
}

// Debug - prints an error row
func Debug(v ...interface{}) {
	Init()
	if debugLogger != nil {
		debugLogger.Println(v...)
	}
}

// Debugf - prints an error row
func Debugf(format string, v ...interface{}) {
	Debug(fmt.Sprintf(format, v...))
}
