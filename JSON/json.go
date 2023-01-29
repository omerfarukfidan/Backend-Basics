package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	//Unmarshalling json to struct and printing
	myJson := `
[
	{
		"first_name": "Omer",
		"last_name": "Faruk"
	}
]`
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error while unmarshalling process!", err)
	}

	log.Printf("unmasrshalled: %v", unmarshalled)

	//Write json from struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Ali"
	m1.LastName = "Hakan"

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Murat"
	m2.LastName = "Demir"

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("Error while marshalling process!", err)
	}
	fmt.Println(string(newJson))
}
