package processor

import (
	"github.com/c2fo/vfs/v6"
	"github.com/sirupsen/logrus"
)

func closeFile(file vfs.File) {
	if file != nil {
		err := file.Close()
		if err != nil {
			logrus.WithFields(logrus.Fields{"URI": file.URI(), "err": err}).Errorf("error when closing file")
		}
	}
}
