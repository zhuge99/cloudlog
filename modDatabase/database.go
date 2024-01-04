package modDatabase

func DB_Initialize() error {
	return nil
}
func DB_AddPostgresql(flag, url string) error {
	return getDBAdapter().addPostgresql(flag, url)
}
func DB_AddInfo(log string) error {
	return getDBAdapter().addLog(1, log)
}
func DB_AddError(log string) error {
	return getDBAdapter().addLog(2, log)
}
