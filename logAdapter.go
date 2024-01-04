package cloudlog

import (
	"fmt"

	"github.com/zhuge99/cloudlog/modDatabase"
)

type IDCLogger interface {
	Info(log string)
	Error(log string)
}

type CDCLogAdapter struct {
	instList []IDCLogger
	dblog    bool
}

var g_SingleLogAdapter *CDCLogAdapter = &CDCLogAdapter{}

func GetLogAdapter() *CDCLogAdapter {
	return g_SingleLogAdapter
}
func (g *CDCLogAdapter) Initialize() error {
	g.dblog = false
	return nil
}
func (g *CDCLogAdapter) Info(args ...any) {
	strLog := fmt.Sprint(args...)
	for _, inst := range g.instList {
		inst.Info(strLog)
	}
	if g.dblog {
		err := modDatabase.DB_AddInfo(strLog)
		if err != nil {
			g.Error(err.Error())
		}
	}
}
func (g *CDCLogAdapter) Error(args ...any) {
	strLog := fmt.Sprint(args...)
	for _, inst := range g.instList {
		inst.Error(strLog)
	}
	if g.dblog {
		err := modDatabase.DB_AddError(strLog)
		if err != nil {
			g.Error(err.Error())
		}
	}
}
func (g *CDCLogAdapter) AddStdout() {
	getLogStdout().Initialize()
	g.instList = append(g.instList, getLogStdout())
}
func (g *CDCLogAdapter) AddLocalFile(basePath, infoFileName, errorFileName string) error {
	localLog := newLogLocalFile()
	err := localLog.Initialize(basePath, infoFileName, errorFileName)
	if err != nil {
		return err
	}
	g_defaultFileLog = localLog
	g.instList = append(g.instList, localLog)

	return nil
}
func (g *CDCLogAdapter) AddLogflare(sourceid, apiKey string) error {
	logflare := newCLogFlare()
	err := logflare.Initialize(sourceid, apiKey)
	if err != nil {
		return err
	}

	g.instList = append(g.instList, logflare)
	return nil
}
func (g *CDCLogAdapter) AddDbPostgres(flag, dburl string) error {
	err := modDatabase.DB_AddPostgresql(flag, dburl)
	if err != nil {
		return err
	}
	g.dblog = true
	return nil
}
