package main

import (
	"fmt"

	"github.com/zhuge99/cloudlog"
)

func main() {

	cloudlog.DCL_addStdout()
	err := cloudlog.DCL_addLocalFileDefault()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = cloudlog.DCL_addLogflare("d0399bff-7fc7-4874-a572-05309021d853", "WoMn49mFQDkh")
	if err != nil {
		fmt.Println(err)
		return
	}
	cloudlog.DCL_Info("hello info")
	cloudlog.DCL_Error("hello error")
}
