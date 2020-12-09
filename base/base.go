package base

import (
	"database/sql"
	"movies/common"
)

var (
	DB *sql.DB
)

func InitMain() {
	var err error
	DB, err = common.InitMySQL()
	if err != nil {
		panic(err)
	}
}
