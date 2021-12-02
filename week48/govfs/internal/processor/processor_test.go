package processor

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"testing"

	"github.com/c2fo/vfs/v6"
	"github.com/c2fo/vfs/v6/backend"
	"github.com/c2fo/vfs/v6/backend/gs"
	"github.com/c2fo/vfs/v6/vfssimple"
	"github.com/fsouza/fake-gcs-server/fakestorage"
	"github.com/otiai10/copy"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func getTestDataPath(pathSegments ...string) (result string) {
	_, filename, _, _ := runtime.Caller(0)
	modRoot := path.Join(path.Dir(filename), "..", "..")
	testDataPath := path.Join(append([]string{modRoot, "test", "data"}, pathSegments...)...)
	return testDataPath
}

func readTestDataFile(pathSegments ...string) (result string) {
	_, filename, _, _ := runtime.Caller(0)
	modRoot := path.Join(path.Dir(filename), "..", "..")
	testDataFilePath := path.Join(append([]string{modRoot, "test", "data"}, pathSegments...)...)
	content, err := ioutil.ReadFile(testDataFilePath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func createObjectsFromDir(t *testing.T, dir string, patternString string, bucketName string) (objects []fakestorage.Object) {
	pattern := regexp.MustCompile(patternString)
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !pattern.MatchString(path) {
			return nil
		}
		name, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		if name == "." {
			return nil
		}
		content, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		objects = append(objects, fakestorage.Object{ObjectAttrs: fakestorage.ObjectAttrs{
			BucketName:      bucketName,
			Name:            name,
			ContentType:     "text/plain",
			ContentEncoding: "utf8",
		}, Content: content})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return objects
}

func readFile(t *testing.T, path string) (result string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(t)
	}
	return string(content)
}

func locationReadFile(location vfs.Location, baseName string) string {
	file, err := location.NewFile(baseName)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func stripDone(baseName string) string {
	return regexp.MustCompile("[.][^.]+[.]done$").ReplaceAllString(baseName, "")
}

func TestProcessFiles(t *testing.T) {
	sourceBucket := "source-bucket"

	type TestCase struct {
		name                     string
		initFilePattern          string
		fsType                   string
		expectedFileInitialCount null.Int
		expectedFileFinalCount   null.Int
		expectedProcessedCount   int
	}
	type TestCases []TestCase

	testCases := TestCases{
		{
			name:                     "gs",
			fsType:                   "gs",
			initFilePattern:          "^.*/file[0-9]+$",
			expectedFileInitialCount: null.NewInt(3, false),
			expectedFileFinalCount:   null.NewInt(6, true),
			expectedProcessedCount:   3,
		},
		{
			name:                     "os",
			fsType:                   "os",
			initFilePattern:          "^.*/file[0-9]+$",
			expectedFileInitialCount: null.NewInt(3, false),
			expectedFileFinalCount:   null.NewInt(6, true),
			expectedProcessedCount:   3,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			var locURI string
			switch testCase.fsType {
			case "gs":
				objects := createObjectsFromDir(t, getTestDataPath(),
					testCase.initFilePattern, sourceBucket)
				server := fakestorage.NewServer(objects)
				defer server.Stop()

				gcfs := backend.Backend("gs").(*gs.FileSystem)
				backend.Unregister("gs")
				gcsClient := server.Client()
				backend.Register("gs", gcfs.WithClient(gcsClient))

				locURI = fmt.Sprintf("gs://%s/", sourceBucket)
			case "os":
				inDir := path.Join(t.TempDir(), "input")
				opt := copy.Options{
					Skip: func(src string) (bool, error) {
						matched, err := regexp.MatchString(testCase.initFilePattern, src)
						return !matched, err
					},
				}
				err := copy.Copy(getTestDataPath(), inDir, opt)
				if err != nil {
					t.Fatal(err)
				}
				locURI = fmt.Sprintf("file://%s/", inDir)
			default:
				t.Fatalf("unsupported inType %s", testCase.fsType)
			}

			t.Logf("locURI = %s", locURI)

			location, err := vfssimple.NewLocation(locURI)
			if err != nil {
				t.Fatal(err)
			}
			filesInitial, err := location.List()
			logrus.Infof("filesInitial = %s", filesInitial)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.expectedFileInitialCount.Valid {
				assert.Len(t, filesInitial, int(testCase.expectedFileInitialCount.Int64))
			}

			processor := Processor{URL: locURI, Mark: "000"}
			t.Logf("converter = %v", processor)
			processedCount, err := processor.Run()
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedProcessedCount, processedCount)
			filesFinal, err := location.List()
			logrus.Infof("filesFinal = %s", filesFinal)
			if err != nil {
				t.Fatal(err)
			}
			if testCase.expectedFileFinalCount.Valid {
				if testCase.expectedFileFinalCount.Int64 >= 0 {
					assert.Len(t, filesFinal,
						int(testCase.expectedFileFinalCount.Int64))
				} else {
					assert.Len(t, filesFinal,
						len(filesInitial)+int(testCase.expectedFileFinalCount.Int64))
				}
			} else {
				assert.Len(t, filesFinal, len(filesInitial))
			}

		})
	}
}
