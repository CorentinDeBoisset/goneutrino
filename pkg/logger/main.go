package logger

import (
	"io"
	"log"
)

var (
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger

	logLevel int
)

func init() {
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	//     log.Fatal(err)
	// }

	debugLogger = log.New(log.Writer(), "[DEBUG] ", log.LstdFlags|log.Lmsgprefix|log.Llongfile)
	infoLogger = log.New(log.Writer(), "[INFO] ", log.LstdFlags|log.Lmsgprefix)
	warningLogger = log.New(log.Writer(), "[WARNING] ", log.LstdFlags|log.Lmsgprefix)
	errorLogger = log.New(log.Writer(), "[ERROR] ", log.LstdFlags|log.Lmsgprefix)
}

func SetLogLevel(level int) {
	logLevel = level
}

func GetLogLevel() int {
	return logLevel
}

func SetLogOutput(w io.Writer) {
	log.SetOutput(w)

	debugLogger.SetOutput(w)
	infoLogger.SetOutput(w)
	warningLogger.SetOutput(w)
	errorLogger.SetOutput(w)
}

func DebugLog(format string, args ...any) {
	if logLevel >= 2 {
		debugLogger.Printf(format, args...)
	}
}

func InfoLog(format string, args ...any) {
	if logLevel >= 1 {
		infoLogger.Printf(format, args...)
	}
}

func WarningLog(format string, args ...any) {
	warningLogger.Printf(format, args...)
}

func ErrorLog(format string, args ...any) {
	errorLogger.Printf(format, args...)
}
