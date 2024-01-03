package cloudlog

// dig cloud log
func DCL_Info(args ...any) {
	GetLogAdapter().Info(args...)
}
func DCL_error(log string, args ...any) {
	GetLogAdapter().Error(args...)
}
func DCL_addStdout() {
	GetLogAdapter().AddStdout()
}
func DCL_addLocalFile(basePath, infoFileName, errorFileName string) error {
	GetLogAdapter().AddLocalFile(basePath, infoFileName, errorFileName)
}
