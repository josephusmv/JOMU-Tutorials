package log_test

import (
	"log"
	"os"
	"testing"
)

func createDefaultLogger() (*log.Logger, error) {
	const cStrDefaultLogFile = "testlogger.log"
	if _, err := os.Stat(cStrDefaultLogFile); err != nil {
		if os.IsNotExist(err) {
			os.Create(cStrDefaultLogFile)
		} else {
			return nil, err
		}
	}

	var f *os.File
	f, err := os.OpenFile(cStrDefaultLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return log.New(f, "testlogger", log.LstdFlags|log.Lshortfile), nil
}

func TestLogs(t *testing.T) {
	logger, err := createDefaultLogger()
	if err != nil {
		panic(err)
	}
	logger.Println("Logger works fine")
}
