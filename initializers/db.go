package initializers

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 3306 // Default port
	user     = "dbadmin"
	password = "d7zB4B3dG9R2Ed6x"
	dbname   = "spc_holding"
)

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	fmt.Println(dsn)
	//DB, err := sql.Open("mysql", dsn)
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	if err = DB.Ping(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database Connect")
		// See "Important settings" section.
		DB.SetConnMaxLifetime(time.Minute * 3)
		DB.SetMaxOpenConns(10)
		DB.SetMaxIdleConns(10)
	}

}
