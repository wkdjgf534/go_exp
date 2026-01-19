package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string   `json:"full_name"`           // Customize field name
	Age        int      `json:"years_old,omitempty"` // Exclude if zero
	Occupation string   `json:"-"`                   // Ignore this field
	Languages  []string `json:"spoken_languages"`
}

func main() {
	jsonData := `{"full_name":"Jane Doe", "years_old":25, "spoken_languages":["French", "Spanish"]}`

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("error unmarshalling JSON:", err)
		panic(err)
	}

	// Print the unmarshalled data
	fmt.Println("Name:", person.Name)
	fmt.Println("Age", person.Age)
	fmt.Println("Languages:", person.Languages)

	// Accessing ignored fields (not recommended)
	fmt.Println("Occupation:", person.Occupation) // Will print an empty string
}
