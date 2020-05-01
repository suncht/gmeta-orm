package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		DisableColors:    false,
	})
}

type OrmLogger struct {
}

func (logger *OrmLogger) Print(values ...interface{}) {
	if len(values) == 3 {
		logger.printError(values)
	} else if len(values) > 5 {
		logger.printNormal(values)
	}
}

func (logger *OrmLogger) printError(values ...interface{}) {
	var (
		level  = values[0]
		source = values[1]
		cause  = values[2]
	)

	logrus.WithFields(logrus.Fields{
		"type":  level,
		"cause": fmt.Sprint(cause),
	}).Debug("SQL执行情况： ", source)
}

func (logger *OrmLogger) printNormal(values ...interface{}) {
	var (
		level   = values[0]
		source  = values[1]
		eclipse = values[2]
		sql     = values[3]
		params  = values[4]
	)

	logrus.WithFields(logrus.Fields{
		"type":    level,
		"sql":     sql,
		"eclipse": eclipse,
		"param":   fmt.Sprint(params),
	}).Debug("SQL执行情况： ", source)
}
