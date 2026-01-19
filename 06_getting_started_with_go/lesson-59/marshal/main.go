package main

import (
	"encoding/json"
	"fmt"
)

//{
//    "user_id": 12345,
//    "name": "User A",
//    "age": 35,
//    "password": "my-password",
//    "roles": ["admin", "collaborator"]
//}

type User struct {
	ID          int      `json:"user_id"`
	Name        string   `json:"name,omitempty"`
	Age         int      `json:"age"`
	Password    string   `json:"-"`     // ignore password
	Permissions []string `json:"roles"` // rename "Permissions" to "roles"
}

func main() {
	u := User{
		ID:          1,
		Name:        "User One",
		Age:         20,
		Password:    "my-password",
		Permissions: []string{"admin", "group-member"},
	}

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error marshaling JSON: ", err)
		panic(err)
	}

	fmt.Println(string(b)) // {"user_id":1,"name":"User One","age":20,"roles":["admin","group-member"]}
}
