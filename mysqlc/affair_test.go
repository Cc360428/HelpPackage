package mysqlc

import (
	"database/sql"
	"testing"
)

func TestAffair(t *testing.T) {

	type args struct {
		db *sql.DB
		// txFunc func(*sql.Tx) error
	}
	var dbclient args

	err := Affair(dbclient.db, func(tx *sql.Tx) error {
		// 这里写 sql 逻辑
		// 有错误返回即可
		return nil
	})
	if err != nil {
		t.Log("Affair update 错误", err.Error())
	}
}
