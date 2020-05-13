package mysql_help

import (
	"database/sql"
	"fmt"
	"github.com/Cc360428/HelpPackage/UtilsHelp/logs"
)
// Trans
func Affair(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
		}
		if err != nil {
			tx.Rollback()
			logs.Info("Rollback, Errmsg: " + err.Error())
			return
		}
		err = tx.Commit()
	}()
	return txFunc(tx)
}
