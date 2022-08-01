package main

import (
	"database/sql"
	"fmt"
	"golang-api/api"
	apiV1 "golang-api/api/v1"
	"golang-api/repository"
	"golang-api/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func init() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?readTimeout=%ds&writeTimeout=%ds", "dev", "Tel4vn.com@2022", "180.93.175.225", 3306, "dev", 10, 30)
	sqldb, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	if err := db.Ping(); err != nil {
		panic(err)
	}
	repository.Db = db
}

func main() {
	server := api.NewServer()

	repository.UserRepo = repository.NewUserRepo()
	apiV1.APIUser(server.Engine, service.NewUser())
	server.Run()
}
