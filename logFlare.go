package cloudlog

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type CLogFlare struct {
	logFlareSourceID string
	logFlareApiKey   string
	url              string
	infoTemplate     string
	errorTemplate    string
	httpRequest      *http.Request
	httpClient       *http.Client
}

var baseUrl = "https://api.logflare.app/logs/json?source="

// log template
// first line: level, info / error
// second line: log, text
var log_template = `{
		"%s": "%s",
		"%s": "%s"
		}`

var keyLevel = "level"
var valueLevelInfo = "info"
var keyLog = "log"
var valueLevelError = "error"

func newCLogFlare() *CLogFlare {
	return &CLogFlare{}
}

func (g *CLogFlare) Initialize(sourceid, apiKey string) error {
	if sourceid == "" {
		return errors.New("logfalre sourceid is empty")
	}
	if apiKey == "" {
		return errors.New("logfalre apiKey is empty")
	}
	g.logFlareSourceID = sourceid
	g.logFlareApiKey = apiKey
	g.url = baseUrl + g.logFlareSourceID
	g.infoTemplate = fmt.Sprintf(log_template, keyLevel, valueLevelInfo, keyLog, "%s")
	g.errorTemplate = fmt.Sprintf(log_template, keyLevel, valueLevelError, keyLog, "%s")
	req, err := http.NewRequest("POST", g.url, nil)
	if err != nil {
		return err
	}
	g.httpRequest = req
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("X-API-KEY", g.logFlareApiKey)
	g.httpClient = &http.Client{}

	return nil
}
func (g *CLogFlare) Info(log string) {
	logtext := fmt.Sprintf(g.infoTemplate, log)
	g.httpRequest.Body = io.NopCloser(bytes.NewBuffer([]byte(logtext)))
	res, err := g.httpClient.Do(g.httpRequest)
	if err != nil {
		errInfo := fmt.Sprintf("http client error when send log to logflare: %s", err.Error())
		getLogStdout().Error(errInfo)
		g_defaultFileLog.Error(errInfo)
		return
	}
	res.Body.Close()
}
func (g *CLogFlare) Error(log string) {
	logtext := fmt.Sprintf(g.errorTemplate, log)
	g.httpRequest.Body = io.NopCloser(bytes.NewBuffer([]byte(logtext)))
	res, err := g.httpClient.Do(g.httpRequest)
	if err != nil {
		errInfo := fmt.Sprintf("http client error when send log to logflare: %s", err.Error())
		getLogStdout().Error(errInfo)
		g_defaultFileLog.Error(errInfo)
		return
	}
	res.Body.Close()
}
