package main

import "fmt"

// 1. Defining a struct.
// 2. Initializing a struct.
// 3. Read and writing struct fields.
// 4. Anonymous struct.
func main() {
	type student struct {
		firstName string
		lastName  string
		age       int
		subjects  []string
	}

	var student1 student
	// +v prints filed names and values
	fmt.Printf("%+v\n", student1) // {firstName: lastName: age:0 subjects:[]}%

	// Initializing a struct
	student1 = student{"code", "learn", 10, []string{"maths", "science"}} // in this case, you have to put all data
	fmt.Printf("%+v\n", student1)                                         // {firstName:code lastName:learn age:10 subjects:[maths science]}

	student2 := student{
		firstName: "foo",
		lastName:  "bar",
		age:       15,
	}

	fmt.Printf("%+v\n", student2)                               // {firstName:foo lastName:bar age:15 subjects:[]}
	fmt.Println("First name of student2: ", student2.firstName) // foo

	student2.subjects = append(student2.subjects, "arts")
	fmt.Printf("%+v\n", student2) // {firstName:foo lastName:bar age:15 subjects:[arts]}

	// Anonymous struct
	guardian := struct {
		firstName string
		lastName  string
	}{
		firstName: "Alex",
		lastName:  "Theo",
	}

	fmt.Println(guardian)
}
