// mysql
package mysql_help

import (
	"database/sql"
	"fmt"
	"log"
)

// Trans 事务
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
			log.Println("Rollback, Errmsg: " + err.Error())
			return
		}
		err = tx.Commit()
	}()
	return txFunc(tx)
}
