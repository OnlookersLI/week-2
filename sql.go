package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

const (
	sqlUser = "root"
	sqlPwd  = "root"
)

// root:123456%^@tcp(39.107.87.114:3306)/book_manager?charset=utf8
func initSql() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", sqlUser+":"+sqlPwd+"@tcp(127.0.0.1:3306)/user?charset=utf8")
	if err != nil {
		err = errors.Wrap(err, "sql Open fail")
		return
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		err = errors.Wrap(err, "sql link fail")
		return
	}

	return
}

type users struct {
	Id       int64
	name    string
	password string
}

func queryAc(db *sql.DB, sqlStr string) (err error) {
	fmt.Println("sqlStr:", sqlStr)
	rows, err := db.Query(sqlStr)
	if err != nil {
		return errors.Wrap(err, "db query fail")
	}

	for rows.Next() {
		var ac users
		err = rows.Scan(&ac.Id, &ac.name, &ac.password)
		if err != nil {
			switch {
			case err == sql.ErrNoRows:
				err = errors.Wrap(err, "ErrNoRows")
			default:
				err = errors.Wrap(err, "scan error")
			}
			return
		}
		fmt.Println(ac)
	}
	fmt.Println("queryAc finished")
	return
}
