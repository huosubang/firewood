package firewood

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func getLogLevel(level string) logrus.Level {
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

func InitLog(cnf *ServerConf) error {
	Log.Level = getLogLevel(Conf.LogLevel)
	Log.Formatter = &logrus.TextFormatter{}

	Log.Out = os.Stdout

	if Conf.AccessLog != ""{
		file, err := os.OpenFile(Conf.AccessLog, os.O_CREATE|os.O_RDWR, 0666)
		if err == nil {
			Log.Out = file
		} else {
			Log.Infof("Failed to log to file %s, err:%s, using default stderr", Conf.AccessLog, err.Error())
		}
	}

	return nil
}
