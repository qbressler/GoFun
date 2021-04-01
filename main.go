package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type People struct {
	PersonId  int
	FirstName string
	LastName  string
	Age       int
	Jobs      []Jobs
}

type Jobs struct {
	Title string
}

func main() {
	people := []People{
		{
			PersonId:  1,
			FirstName: "Joe",
			LastName:  "Doe",
			Age:       18,
			Jobs: []Jobs{
				{
					Title: "Programmer",
				},
				{
					"Sales Clerk",
				},
			},
		},
		{
			PersonId:  2,
			FirstName: "Jane",
			LastName:  "Doe",
			Age:       19,
		},
	}
	for _, x := range people {
		fmt.Printf("%d %s\n", x.PersonId, x.FirstName)
	}

	peopleJson, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", peopleJson)
}
