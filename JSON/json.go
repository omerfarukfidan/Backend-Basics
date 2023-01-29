package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

func main() {

	myJson := `
[
	{
		"first_name": "Omer",
		"second_name": "Fidan"
	}
]`
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error while unmarshalling process!", err)
	}

	log.Printf("unmasrshalled: %v", unmarshalled)

}
