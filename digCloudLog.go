package cloudlog

// dig cloud log
func DCL_Info(args ...any) {
	GetLogAdapter().Info(args...)
}
func DCL_Error(args ...any) {
	GetLogAdapter().Error(args...)
}
func DCL_addStdout() {
	GetLogAdapter().AddStdout()
}
func DCL_addLocalFileDefault() error {
	return GetLogAdapter().AddLocalFile("", "", "")
}
func DCL_addLocalFile(basePath, infoFileName, errorFileName string) error {
	return GetLogAdapter().AddLocalFile(basePath, infoFileName, errorFileName)
}
func DCL_addLogflare(sourceid, apiKey string) error {
	return GetLogAdapter().AddLogflare(sourceid, apiKey)
}

/// db log
func DCL_addPostgresql(flag, dburl string) error {
  return modDatabase.GetDBAdapter().AddPostgresql(dburl)
}
