package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[KeyType]valueType)

	// using a Map Literal
	// mapVariable = map[keyType]valueType{
	//	key1: value1,
	//	key2: value2,
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap) // map[] - an empty map

	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)         //map[code:18 key1:9]
	fmt.Println(myMap["key1"]) // 9
	fmt.Println(myMap["key"])  // 0 - if key don't exist it returns a default value for valueType

	myMap["code"] = 21
	fmt.Println(myMap) //map[code:21 key1:9]

	// delete(myMap, "key1") // delete a specific key - value pair
	// fmt.Println(myMap)    //map[code:21]

	myMap["key2"] = 10
	myMap["key3"] = 11
	myMap["key4"] = 12
	fmt.Println(myMap) //map[code:21 key1:9 key2:10 key3:11 key4:12]

	// clear(myMap)       // remove all elements of a map
	// fmt.Println(myMap) // map[]

	value, unknownValue := myMap["key1"]
	fmt.Println(value, unknownValue) // 9 true

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap3 and myMap2 are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
		// a 1
		// b 2
	}

	_, ok := myMap["key1"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("no value exists with key1")
	}

	var myMap4 map[string]string
	if myMap4 != nil {
		fmt.Println("The map is initialized to nil value.")
	} else {
		fmt.Println("The map is not initialized to nil value.")
	}

	val := myMap4["key"]
	fmt.Println(val) // "" - default value for string

	//myMap4["key"] = "1212" // This provoke a panic, aasignment to entry in nil map
	//fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "1212"
	fmt.Println(myMap4) // map[key:1212]

	// nested maps
	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4 // map[map1:map[key:1212]]
	fmt.Println(myMap5)
}
