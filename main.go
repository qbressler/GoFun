package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type people struct {
	personID  int
	firstName string
	lastName  string
	age       int
	jobs      []jobs
}

type jobs struct {
	title string
}

func main() {
	//var u usermodel.User = usermodel.CreateUser()
	//fmt.Printf("%v", u)
	people := []people{
		{
			personID:  1,
			firstName: "Joe",
			lastName:  "Doe",
			age:       18,
			jobs: []jobs{
				{
					title: "Programmer",
				},
				{
					title: "Sales Clerk",
				},
			},
		},
		{
			personID:  2,
			firstName: "Jane",
			lastName:  "Doe",
			age:       19,
		},
	}
	for _, x := range people {
		fmt.Printf("%d %s\n", x.age, x.firstName)
	}

	fmt.Println(people)
	peopleJSON, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", peopleJSON)
}
