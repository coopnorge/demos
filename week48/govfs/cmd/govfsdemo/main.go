package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"modernc.org/mathutil"
)

func init() {
	levelString := os.Getenv("LOGRUS_LEVEL")
	useDefault := true
	if len(levelString) > 0 {
		setLevel, err := strconv.ParseUint(levelString, 10, 64)
		if err == nil {
			maxLevel := uint64(logrus.TraceLevel)
			minLevel := uint64(logrus.PanicLevel)
			var newLevel = mathutil.MaxUint64(mathutil.MinUint64(setLevel, maxLevel), minLevel)
			logrus.SetLevel(logrus.Level(newLevel))
			useDefault = false
		}
	}
	if useDefault {
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	logrus.WithField("args", os.Args).Debugf("entry")

	if err := rootCmd.Execute(); err != nil {
		logrus.Debugf("got err = %v", err)
		os.Exit(1)
	}

}
