package database

import (
	"backend-task-eventori/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	config := config.GetConfig()

	connectionString := config.DB_USERNAME+":"+config.DB_PASSWORD+"@/"+config.DB_NAME+"?parseTime=true&loc=Asia%2FJakarta"

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	} 
	
	return db
}
