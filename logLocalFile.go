package cloudlog

import (
	"internal/goarch"
	"log"
	"os"
	"path/filepath"
	"time"
)

type CLogLocalFile struct {
	basePath string
	logInfo  *log.Logger
	logError *log.Logger
}

var g_defaultFileLog = &CLogLocalFile{logInfo: &log.Logger{}, logError: &log.Logger{}}

func newLogLocalFile() *CLogLocalFile {
	return &CLogLocalFile{}
}
func (g *CLogLocalFile) Initialize(basePath, infoFileName, errorFileName string) error {

	return g.initlogFile(basePath, infoFileName, errorFileName)
}
func (g *CLogLocalFile) initlogFile(basePath, infoFileName, errorFileName string) error {
	g.basePath = basePath
	if basePath == "" {
		ex, err := os.Executable()
		if err != nil {
			return err
		}
		exPath := filepath.Dir(ex)
		g.basePath = filepath.Join(exPath, "logs")
	}
	if infoFileName == "" {
		infoFileName = "info.log"
	}
	if errorFileName == "" {
		errorFileName = "error.log"
	}

	logInst, err := g.opeLogFile(infoFileName)
	if err != nil {
		return err
	}
	g.logInfo = logInst
	logInst, err = g.opeLogFile(errorFileName)
	if err != nil {
		return err
	}
	g.logError = logInst

	return nil
}
func (g *CLogLocalFile) opeLogFile(path string) (*log.Logger, error) {
	logfilePath := filepath.Join(g.basePath, path)
	fileHandler, err := os.OpenFile(logfilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		getLogStdout().Error(err.Error())
		return nil, err
	}
	logInst := log.New(fileHandler, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	logInst.Println(time.Now(), ": local file log start")

	return logInst, nil
}
func (g *CLogLocalFile) Info(log string) {
	g.logInfo.Println(log)
}
func (g *CLogLocalFile) Error(log string) {
	g.logError.Println(log)
}
