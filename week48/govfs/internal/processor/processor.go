package processor

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"

	"github.com/c2fo/vfs/v6/vfssimple"
	"github.com/sirupsen/logrus"
)

var (
	markerPattern = regexp.MustCompile("^(.*)[.]marker-([0-9]+)$")
)

type Processor struct {
	URL  string
	Mark string
}

type void struct{}
type stringSet map[string]void

func (p *Processor) Run() (int, error) {
	log := logrus.WithFields(logrus.Fields{"URL": p.URL, "Mark": p.Mark})
	log.Infof("entry")

	loc, err := vfssimple.NewLocation(p.URL)
	if err != nil {
		return 0, err
	}

	logrus.Debugf("listing files")
	files, err := loc.List()

	fileSet := stringSet{}
	fileMarkers := map[string]stringSet{}

	for fileIndex, fileName := range files {
		flog := log.WithFields(logrus.Fields{"fileName": fileName, "fileIndex": fileIndex})
		flog.Debugf("checking")
		match := markerPattern.FindStringSubmatch(fileName)
		var baseName string
		var foundMark *string = nil
		if match == nil {
			baseName = fileName
			fileSet[baseName] = void{}
		} else {
			baseName = match[1]
			foundMark = &match[2]
		}
		var fileMarkersSet stringSet
		var ok bool
		if fileMarkersSet, ok = fileMarkers[baseName]; !ok {
			fileMarkers[baseName] = stringSet{}
			fileMarkersSet = fileMarkers[baseName]
		}
		if foundMark != nil {
			fileMarkersSet[*foundMark] = void{}
		}
	}

	count := 0
	for fileName, _ := range fileSet {
		markers := fileMarkers[fileName]

		flog := log.WithFields(logrus.Fields{"fileName": fileName, "markers": markers})
		flog.Debugf("checking")

		if _, marked := markers[p.Mark]; !marked {
			flog.Debugf("marking")
		} else {
			flog.Infof("already marked")
			continue
		}

		// Open the file
		file, err := loc.NewFile(fileName)
		if err != nil {
			return count, err
		}
		defer closeFile(file)

		// Read and hash the file
		fileHasher := sha256.New()
		if _, err := io.Copy(fileHasher, file); err != nil {
			return count, err
		}
		fileHash := hex.EncodeToString(fileHasher.Sum(nil))
		fileMark := fmt.Sprintf("%s.marker-%s", fileName, p.Mark)

		// Write a marker file with the hash content
		markFile, err := loc.NewFile(fileMark)
		if err != nil {
			return count, err
		}
		defer closeFile(markFile)
		_, err = markFile.Write([]byte(fmt.Sprintf("%s\n", fileHash)))
		if err != nil {
			return count, err
		}

		flog.Infof("marked")
		count++
	}

	return count, err
}
