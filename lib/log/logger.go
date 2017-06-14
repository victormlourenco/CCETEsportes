package log

import (
	"fmt"
	"os"

	"CCETEsportes/lib/log/config"

	"github.com/Sirupsen/logrus"
)

const (
	LoggerConsoleConst string = "console"
	LoggerFileConst    string = "file"
)

type Content struct {
	Message          string
	AdditionalFields interface{}
}

var loggers map[string]*logrus.Logger

// var consoleLogger = logrus.New()
// var fileLogger = logrus.New()

func init() {
	loggers = make(map[string]*logrus.Logger)
	instantiateConsoleLogger()
	//instantiateFileLogger()
}

func instantiateConsoleLogger() {
	consoleLogger := logrus.New()
	consoleLogger.Formatter = &logrus.TextFormatter{FullTimestamp: true}
	consoleLogger.Level = logrus.InfoLevel
	loggers[LoggerConsoleConst] = consoleLogger
}

func instantiateFileLogger() {
	filepath := config.GetString("FilePath") + config.GetString("FileName")
	fileFlags := os.O_CREATE | os.O_WRONLY | os.O_APPEND

	outputFile, err := os.OpenFile(filepath, fileFlags, 0644)
	if err != nil {
		fmt.Println("[ERROR] Unable to open log file")
	}

	fileLogger := logrus.New()
	fileLogger.Out = outputFile
	fileLogger.Formatter = &logrus.TextFormatter{FullTimestamp: true, ForceColors: true}
	fileLogger.Level = logrus.InfoLevel
	loggers[LoggerFileConst] = fileLogger
}

func Debug(content Content) {
	if content.AdditionalFields != nil {
		logWithFields(logrus.DebugLevel, content)
	} else {
		log(logrus.DebugLevel, content.Message)
	}
}

func Info(content Content) {
	if content.AdditionalFields != nil {
		logWithFields(logrus.InfoLevel, content)
	} else {
		log(logrus.InfoLevel, content.Message)
	}
}

func Warn(content Content) {
	if content.AdditionalFields != nil {
		logWithFields(logrus.WarnLevel, content)
	} else {
		log(logrus.WarnLevel, content.Message)
	}
}

func Error(content Content) {
	if content.AdditionalFields != nil {
		logWithFields(logrus.ErrorLevel, content)
	} else {
		log(logrus.ErrorLevel, content.Message)
	}
}

func Fatal(content Content) {
	if content.AdditionalFields != nil {
		logWithFields(logrus.FatalLevel, content)
	} else {
		log(logrus.FatalLevel, content.Message)
	}
}

func Panic(content Content) {
	if content.AdditionalFields != nil {
		logWithFields(logrus.PanicLevel, content)
	} else {
		log(logrus.PanicLevel, content.Message)
	}
}

func log(level logrus.Level, message string) {
	switch level {
	case logrus.DebugLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.Debug(message)
			}
		}
	case logrus.InfoLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.Info(message)
			}
		}
	case logrus.WarnLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.Warn(message)
			}
		}
	case logrus.ErrorLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.Error(message)
			}
		}
	case logrus.FatalLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.Fatal(message)
			}
		}
	case logrus.PanicLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.Panic(message)
			}
		}
	}
}

// func log(level logrus.Level, message string) {
// 	switch level {
// 	case logrus.DebugLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.Debug(message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.Debug(message)
// 		}
// 	case logrus.InfoLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.Info(message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.Info(message)
// 		}
// 	case logrus.WarnLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.Warn(message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.Warn(message)
// 		}
// 	case logrus.ErrorLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.Error(message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.Error(message)
// 		}
// 	case logrus.FatalLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.Fatal(message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.Fatal(message)
// 		}
// 	case logrus.PanicLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.Panic(message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.Panic(message)
// 		}
// 	}
// }

func logWithFields(level logrus.Level, content Content) {
	switch level {
	case logrus.DebugLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.WithFields(logrus.Fields{
					"AdditionalFields": content.AdditionalFields,
				}).Debug(content.Message)
			}
		}
	case logrus.InfoLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.WithFields(logrus.Fields{
					"AdditionalFields": content.AdditionalFields,
				}).Info(content.Message)
			}
		}
	case logrus.WarnLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.WithFields(logrus.Fields{
					"AdditionalFields": content.AdditionalFields,
				}).Warn(content.Message)
			}
		}
	case logrus.ErrorLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.WithFields(logrus.Fields{
					"AdditionalFields": content.AdditionalFields,
				}).Error(content.Message)
			}
		}
	case logrus.FatalLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.WithFields(logrus.Fields{
					"AdditionalFields": content.AdditionalFields,
				}).Fatal(content.Message)
			}
		}
	case logrus.PanicLevel:
		for _, logger := range loggers {
			if logger != nil {
				logger.WithFields(logrus.Fields{
					"AdditionalFields": content.AdditionalFields,
				}).Panic(content.Message)
			}
		}
	}
}

// func logWithFields(level logrus.Level, content Content) {
// 	switch level {
// 	case logrus.DebugLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Debug(content.Message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Debug(content.Message)
// 		}
// 	case logrus.InfoLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Info(content.Message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Info(content.Message)
// 		}
// 	case logrus.WarnLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Warn(content.Message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Warn(content.Message)
// 		}
// 	case logrus.ErrorLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Error(content.Message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Error(content.Message)
// 		}
// 	case logrus.FatalLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Fatal(content.Message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Fatal(content.Message)
// 		}
// 	case logrus.PanicLevel:
// 		if consoleLogger != nil {
// 			consoleLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Panic(content.Message)
// 		}
// 		if fileLogger != nil {
// 			fileLogger.WithFields(logrus.Fields{
// 				"AdditionalFields": content.AdditionalFields,
// 			}).Panic(content.Message)
// 		}
// 	}
// }
