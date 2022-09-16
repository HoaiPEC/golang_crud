package dal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func InitializeMySQL() {
	dBConnection, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/crud_golang")
	if err != nil {
		fmt.Println("Connection Failed!!")
	}
	err = dBConnection.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	}
	db = dBConnection
	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)
}

func GetConnection() *sql.DB {
	if db == nil {
		InitializeMySQL()
	}
	return db
}

// CloseStmt after run stmt
func CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		err := stmt.Close()
		if err != nil {
			return
		}
	}
}

// CloseRows when select
func CloseRows(rows *sql.Rows) {
	if rows != nil {
		err := rows.Close()
		if err != nil {
			return
		}
	}
}
