package db

import (
	"database/sql"
	"log"
)

type blogItem struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func getconn() (*sql.DB, error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root123"
	dbName := "blog"
	conn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		log.Println("Something Went Wrong")
		panic(err.Error())
	}
	log.Println("DataBase is Connected")

	return conn, nil

}
