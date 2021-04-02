package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//  Employees this is the employee struct...very simple data
type Employees struct {
	id        int
	firstname string
	lastname  string
}

func main() {
	db, err := sql.Open("mysql", "quintin:******@/corenet")
	if err != nil {
		fmt.Printf("%v", err)
	}

	rows, err := db.Query("select firstname, lastname from Employees")
	if err != nil {
		log.Fatal(err)
	}

	employeesSlice := []Employees{}
	for rows.Next() {
		e := Employees{}
		err = rows.Scan(&e.firstname, &e.lastname)
		employeesSlice = append(employeesSlice, e)
	}

	fmt.Println(employeesSlice)
}
