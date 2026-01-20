package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type CarType int

const (
	_ CarType = iota
	Hatchback
	Sedan
	SUV
)

var (
	stringToCarType = map[string]CarType{
		"hatchback": Hatchback,
		"sedan":     Sedan,
		"suv":       SUV,
	}

	CarTypeToString = map[CarType]string{
		Hatchback: "hatchback",
		Sedan:     "sedan",
		SUV:       "suv",
	}
)

func (c CarType) String() string {
	return CarTypeToString[c]
}

func (c CarType) MarshalJSON() ([]byte, error) {
	t, ok := CarTypeToString[c]
	if !ok {
		return nil, errors.New("invalid car type provided")
	}

	return []byte(`"` + t + `"`), nil
}

func (c *CarType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, ok := stringToCarType[s]
	if !ok {
		return fmt.Errorf("invalid car type: %s", s)
	}

	*c = t
	return nil
}

func main() {
	c := Hatchback

	// Marshal

	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error marshalling car type: ", err)
		os.Exit(1)
	}

	fmt.Println(string(b))

	// Unmarshal

	var UnmarshalledCarType CarType
	err = json.Unmarshal([]byte(`"suv"`), &UnmarshalledCarType)
	if err != nil {
		fmt.Println("error unmarshalling car type: ", err)
		os.Exit(1)
	}

	fmt.Println(UnmarshalledCarType)
}
