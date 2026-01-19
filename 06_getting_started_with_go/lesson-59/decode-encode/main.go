package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"is_active"`
}

func main() {
	personDec := Person{
		Name:   "Alice",
		Age:    25,
		Active: true,
	}

	// Decode

	f, err := os.Create("output.json")
	if err != nil {
		fmt.Println("error creating file: ", err)
		panic(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(personDec)
	if err != nil {
		fmt.Println("error encoding person: ", err)
		panic(err)
	}

	// Encode

	var personEnc Person
	f, err = os.Open("output.json")
	if err != nil {
		fmt.Println("error reading file: ", err)
		panic(err)
	}
	defer f.Close()

	dencoder := json.NewDecoder(f)
	err = dencoder.Decode(&personEnc)
	if err != nil {
		fmt.Println("error encoding person: ", err)
		panic(err)
	}

	fmt.Println(personEnc) // {Alice 25 true}
}
