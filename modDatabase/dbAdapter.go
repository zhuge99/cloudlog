package modDatabase

import (
	"errors"
	"fmt"
	"strconv"
)

type IDBOperator interface {
	Initialize(flag, dburl string) error
	CheckTableExists(tableName string) (bool, error)
	Execsql(sql string) error
	GetFlag() string
}

type CDbAdapter struct {
	dbList []IDBOperator
}

const pqsqlCreateTableLog = `
CREATE TABLE digCloudLog (
	id SERIAL PRIMARY KEY,
	logtime timestamp NOT NULL DEFAULT NOW(),
	loglevel int NOT NULL,
	logcontent text
 );
`

var g_singleDBAdapter *CDbAdapter = &CDbAdapter{}

func getDBAdapter() *CDbAdapter {
	return g_singleDBAdapter
}
func (instSelf *CDbAdapter) addPostgresql(flag, url string) error {
	dbInst := newDbPostgres()
	if dbInst == nil {
		return errors.New("newDbPostgres failed")
	}
	err := dbInst.Initialize(flag, url)
	if err != nil {
		return err
	}

	err = instSelf.initLogTable(dbInst, pqsqlCreateTableLog)
	if err != nil {
		return err
	}

	instSelf.dbList = append(instSelf.dbList, dbInst)

	return nil
}

func (instSelf *CDbAdapter) initLogTable(dbo IDBOperator, createsql string) error {
	exists, err := dbo.CheckTableExists("digCloudLog")
	if err != nil {
		return errors.New("error when check table exists " + err.Error())
	}
	if exists {
		return nil
	}

	err = dbo.Execsql(createsql)
	if err != nil {
		return errors.New("error when create log table " + err.Error())
	}
	if exists {
		return nil
	}

	// create table;

	return nil
}
func (instSelf *CDbAdapter) addLog(level int, log string) error {
	errorInfo := ""
	sql := "insert into digCloudLog(loglevel,logcontent) values(" + strconv.Itoa(level) + ",'" + log + "');"
	fmt.Println(sql)
	for _, v := range instSelf.dbList {
		err := v.Execsql(sql)
		if err != nil {
			errorInfo = errorInfo + "dblog: [" + v.GetFlag() + "] add log error, level: " + strconv.Itoa(level) + ", error: " + err.Error() + "\n"
			fmt.Println(errorInfo)
		}
	}
	if errorInfo != "" {
		return errors.New(errorInfo)
	}
	return nil
}
