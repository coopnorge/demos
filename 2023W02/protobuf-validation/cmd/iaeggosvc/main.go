package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"modernc.org/mathutil"
)

type extraFieldsHook struct {
	Fields logrus.Fields
}

func (h extraFieldsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h extraFieldsHook) Fire(e *logrus.Entry) error {
	for key, value := range h.Fields {
		e.Data[key] = value
	}
	return nil
}

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
	var filename string
	_, _filename, _, ok := runtime.Caller(0)
	if ok {
		filename = _filename
	} else {
		filename = "/dev/null"
	}
	filedir := filepath.Dir(filepath.Dir(filepath.Dir(filename)))

	callerPrettyfier := func(frame *runtime.Frame) (string, string) {
		relpath, err := filepath.Rel(filedir, frame.File)
		if err != nil {
			return frame.Function, fmt.Sprintf("%s:%d", frame.File, frame.Line)
		}
		return frame.Function, fmt.Sprintf("%s:%d", relpath, frame.Line)
	}

	logrusConsole := false
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: callerPrettyfier,
	})
	if os.Getenv("LOGRUS_CONSOLE") != "" {
		value := os.Getenv("LOGRUS_CONSOLE")
		err := json.Unmarshal([]byte(value), &logrusConsole)
		if err != nil {
			logrus.Errorf("ignoring unparsable LOGRUS_CONSOLE value %#v: %s", value, err)
		}
	}
	if logrusConsole {
		logrus.SetFormatter(&logrus.TextFormatter{
			CallerPrettyfier: callerPrettyfier,
		})
	}
	logrus.SetReportCaller(true)

	logrus.AddHook(extraFieldsHook{Fields: logrus.Fields{
		"pid":      os.Getpid(),
		"instance": uuid.New().String(),
	}})
}

func commandPath(cmd *cobra.Command) string {
	path := []string{}
	var cmdPtr = cmd
	for cmdPtr != nil {
		path = append([]string{cmdPtr.Use}, path...)
		cmdPtr = cmdPtr.Parent()
	}
	return strings.Join(path, "/")
}

func main() {
	logrus.WithField("args", os.Args).Debugf("entry")

	for _, envFileName := range []string{".env", "default.env"} {
		log := logrus.WithField("envFileName", envFileName)
		log.Debug("loading")
		err := godotenv.Load(envFileName)
		if err != nil {
			log = log.WithField("err", err)
			log.Debugf("godotenv error")
			var patherr *fs.PathError
			if !errors.As(err, &patherr) {
				log.Fatalf("Error loading .env files: %v", err)
				panic(err)
			}
		} else {
			log.Debugf("godotenv loaded")
		}
	}

	ctx := context.Background()
	if err := func() error {
		return rootCmd.ExecuteContext(ctx)
	}(); err != nil {
		logrus.WithField("err", err).Error("failed")
		os.Exit(1)
	}
}
