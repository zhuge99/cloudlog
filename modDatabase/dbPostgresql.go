package modDatabase

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

type CDbPostgres struct {
	flag        string
	dburl       string
	pgxConn     *pgx.Conn
	ctx         context.Context
	ignoreLimit int
	ignoreTimes int
}

func newDbPostgres() *CDbPostgres {
	inst := &CDbPostgres{}
	return inst
}
func (instSelf *CDbPostgres) Initialize(flag, dburl string) error {
	instSelf.dburl = dburl
	instSelf.flag = flag
	instSelf.ignoreLimit = 0
	instSelf.ignoreTimes = 0
	err := instSelf.connect()
	if err != nil {
		return err
	}
	return nil
}
func (instSelf *CDbPostgres) GetFlag() string {
	return instSelf.flag
}
func (instSelf *CDbPostgres) connect() error {
	instSelf.ctx = context.Background()
	config, err := pgx.ParseConfig(instSelf.dburl)
	if err != nil {
		return err
	}
	conn, err := pgx.ConnectConfig(instSelf.ctx, config)
	if err != nil {
		return err
	}
	instSelf.pgxConn = conn
	return nil
}
func (instSelf *CDbPostgres) CheckTableExists(tableName string) (bool, error) {
	var n int64
	err := instSelf.pgxConn.QueryRow(instSelf.ctx, "select 1 from information_schema.tables where table_name = $1", tableName).Scan(&n)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return true, nil
		}
		return false, err
	}

	return true, nil
}
func (instSelf *CDbPostgres) Execsql(sql string) error {
	if instSelf.ignoreTimes < instSelf.ignoreLimit {
		instSelf.ignoreTimes++
		return nil
	}
	_, err := instSelf.pgxConn.Exec(instSelf.ctx, sql)
	if err != nil {
		instSelf.ignoreLimit++
	} else {
		instSelf.ignoreLimit = 0
		instSelf.ignoreTimes = 0
	}
	return err
}
