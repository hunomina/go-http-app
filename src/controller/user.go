package controller

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	. "model"
	"net/http"
)

type UserController struct {
}

func (c UserController) GetAllUsers(w http.ResponseWriter, req *http.Request) {

	sqlConnection, err := sql.Open("mysql", "root:chemene123@tcp(db)/go_http_app")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer sqlConnection.Close()

	err = sqlConnection.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	sqlQuery := "SELECT * FROM `user`;"

	rows, err := sqlConnection.Query(sqlQuery)
	if err != nil {
		fmt.Println("Sql Query : " + err.Error())
		return
	}
	defer rows.Close()

	var result []User
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.Age)
		if err != nil {
			fmt.Println("Fetch user : " + err.Error())
		}
		println(u.Id, u.Name, u.Email, u.Password, u.Age)
		result = append(result, u)
	}
}
