package utils

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

type plainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

func (f *plainFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	if pc, file, line, ok := runtime.Caller(7); ok {
		funcName := runtime.FuncForPC(pc).Name()

		fil := fmt.Sprintf("%s : %v", path.Base(file), line)
		fun := fmt.Sprintf("%s", path.Base(funcName))
		return []byte(fmt.Sprintf("[%s]   %s   %-70s | %-20s | %s \n", f.LevelDesc[entry.Level], timestamp, entry.Message, fil, fun)), nil
	}
	return nil, nil
}

func InitializeLogger() {
	os.MkdirAll(os.Getenv("SERVER_LOGS_PATH"), os.ModePerm)
	file, err := os.OpenFile(os.Getenv("SERVER_LOGS_PATH"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file in path: "+os.Getenv("SERVER_LOGS_PATH"), " error: ", err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.SetReportCaller(true)
	plainFormatter := new(plainFormatter)
	plainFormatter.TimestampFormat = "2006-01-02 15:04:05"
	plainFormatter.LevelDesc = []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG"}
	log.SetFormatter(plainFormatter)

	// log.SetLevel(log.ErrorLevel)

	log.Info("Log formatter initialized correctly")
}
