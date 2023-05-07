package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/nums")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//prepare statement for inserting data
	stmtDel, err := db.Prepare("DELETE FROM sqarenum;") // ? =placeholder
	if err != nil {
		panic(err.Error()) //proper error handling instead of panic in your app

	}
	defer stmtDel.Close()

	//prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO sqarenum VALUES ( ?, ?)") // ?= placeholder
	if err != nil {
		panic(err.Error())
	}
	defer stmtDel.Close()

	//prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT squareNumber FROM sqarenum WHERE number = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	_, err = stmtDel.Exec()
	if err != nil {
		panic(err.Error())
	}

	//insert square numbers 0-24 in the db
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i, (i * i)) //insert tuples (i, i^2)
		if err != nil {
			panic(err.Error())
		}

	}
	var sqareNum int //we "scan" the result in here
	//Query the square-Number of 13
	err = stmtOut.QueryRow(13).Scan(&sqareNum) //where number = 13
	if err != nil {
		panic(err.Error())

	}
	fmt.Printf("The square number of 13 is: %d \n", sqareNum)

	err = stmtOut.QueryRow(1).Scan(&sqareNum)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 1 is: %d \n", sqareNum)
}
