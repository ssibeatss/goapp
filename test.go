package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var x int
	db, err := sql.Open("mysql", "root:passwd@tcp(0.0.0.0:3306)/user")
	for {
		fmt.Println("Enter a number from 1 to 10")
		fmt.Scanln(&x)
		if x > 0 && x < 11 {
			break
		}
	}
}
