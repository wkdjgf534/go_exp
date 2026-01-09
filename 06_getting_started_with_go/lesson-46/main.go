package main

import "fmt"

// Engine represents a car engine
type Engine struct {
	Model      string
	HorsePower int
}

// GPS represents a GPS device
type GPS struct {
	Model string
}

// Car represents a car with an engine and GPS
type Car struct {
	Model string
	Engine
	GPS
}

func main() {
	myCar := Car{
		Model: "Toyota",
		Engine: Engine{
			Model:      "Mazda",
			HorsePower: 200,
		},
		GPS: GPS{Model: "Garmin"},
	}

	fmt.Println("Car model:", myCar.Model)
	fmt.Println("Engine model:", myCar.Engine.Model) // here is a name conflicts have too put struct name too
	fmt.Println("GPS model:", myCar.GPS.Model)       // the same situation here
	fmt.Println("Car horsepower:", myCar.HorsePower) // you can use it directly
}
