package cloudlog

import (
	"fmt"
	"time"
)

type CLogStdout struct {
}

var g_singleLogStdout *CLogStdout = &CLogStdout{}

func getLogStdout() *CLogStdout {
	return g_singleLogStdout
}

func (g *CLogStdout) Initialize() error {
	return nil
}
func (g *CLogStdout) Info(log string) {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"), "[I] ", log)
}
func (g *CLogStdout) Error(log string) {
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"), "[X] ", log)
}
