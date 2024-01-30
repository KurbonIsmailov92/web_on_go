package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	fmt.Println("We got connection to DB")

	// ustanovka dannikh
	// insert, err := db.Query("INSERT INTO users (name, age) VALUES('Max', 33)")
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()

	//Selecting data
	res, err := db.Query("SELECT name, age FROM users")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var User User
		err = res.Scan(&User.Name, &User.Age)
		if err != nil {
			panic(err)
		}

		fmt.Printf(fmt.Sprintf("User: %s with age %d \n", User.Name, User.Age))
	}

}
