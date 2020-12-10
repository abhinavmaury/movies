package common

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	maxOpenConnection    = 10
	maxConnectionTimeout = time.Hour * 1
	maxIdleConnection    = 5
)

func InitMySQL() (*sql.DB, error) {
	client, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci", "root", "root1234", "full_db_mysql", "3306", "movies"))
	client.SetMaxOpenConns(maxOpenConnection)
	client.SetMaxIdleConns(maxIdleConnection)
	client.SetConnMaxLifetime(maxConnectionTimeout)
	return client, err
}
