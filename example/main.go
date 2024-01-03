package main

import (
	"fmt"
	"github.com/zhuge99/cloudlog"
)

func main() {
	cloudlog.DCL_addStdout()
	cloudlog.DCL_addLocalFileDefault()
	cloudlog.DCL_addLogflare("", "")
	cloudlog.DCL_Info("hello info")
	cloudlog.DCL_Error("hello error")
}
