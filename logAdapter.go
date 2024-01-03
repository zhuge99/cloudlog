package cloudlog

import "fmt"

type IDCLogger interface {
	//Initialize() error
	Info(log string)
	Error(log string)
}

type CDCLogAdapter struct {
	instList []IDCLogger
}

var g_SingleLogAdapter *CDCLogAdapter = &CDCLogAdapter{}

func GetLogAdapter() *CDCLogAdapter {
	return g_SingleLogAdapter
}
func (g *CDCLogAdapter) Initialize() error {
	return nil
}
func (g *CDCLogAdapter) Info(args ...any) {
	strLog := fmt.Sprintln(args...)
	for _, inst := range g.instList {
		inst.Info(strLog)
	}
}
func (g *CDCLogAdapter) Error(args ...any) {
	strLog := fmt.Sprintln(args...)
	for _, inst := range g.instList {
		inst.Error(strLog)
	}
}
func (g *CDCLogAdapter) AddStdout() {
	getLogStdout().Initialize()
	g.instList = append(g.instList, getLogStdout())
}
func (g *CDCLogAdapter) AddLocalFile(basePath, infoFileName, errorFileName string) error {
	localLog := &CLogLocalFile{}
	err := localLog.Initialize(basePath, infoFileName, errorFileName)
	if err != nil {
		return err
	}
	g.instList = append(g.instList, &CLogLocalFile{})

	return nil
}
