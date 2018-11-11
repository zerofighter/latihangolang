package database

import (
	"database/sql"
  "fmt"
  "../common"
	_ "github.com/lib/pq"
)

var DBCon *sql.DB

func InitDb(){
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        common.AppConfig.DBHost, common.AppConfig.DBPort,
        common.AppConfig.DBUser, common.AppConfig.DBPwd, common.AppConfig.Database)

    DBCon, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected!")
}
