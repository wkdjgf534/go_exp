package main

import "fmt"

type multiplier func(int) int

type operation func(int, int) int

func main() {
	result1 := apply(double, 5)
	fmt.Println(result1)

	result2 := apply(triple, 5)
	fmt.Println(result2)

	fiveTimes := func(x int) int {
		return 5 * x
	}
	result3 := apply(fiveTimes, 5)
	fmt.Println(result3)

	multiplyByTwo := multiplyBy(2)
	multiplyByThree := multiplyBy(3)

	result4 := multiplyByTwo(5)
	fmt.Println(result4)

	result5 := multiplyByThree(3)
	fmt.Println(result5)

	//var perform func(int, int) int
	//perform = arithmeticOperation("add")
	//result6 := perform(1, 2)
	//fmt.Println(result6)

	var perform1 operation
	perform1 = arithmeticOperation("add")
	result6 := perform1(1, 2)
	fmt.Println(result6)

	var perform2 operation
	perform2 = arithmeticOperation("subtract")
	result7 := perform2(3, 2)
	fmt.Println(result7)

}

func double(x int) int {
	return x * 2
}

func triple(x int) int {
	return x * 3
}

func apply(f func(int) int, x int) int {
	return f(x)
}

//func multiplyBy(m int) func(int) int {
//	return func(i int) int {
//		return i * m
//	}
//}

func multiplyBy(m int) multiplier {
	return func(i int) int {
		return i * m
	}
}

func arithmeticOperation(op string) operation {
	switch op {
	case "add":
		return func(i1, i2 int) int {
			return i1 + i2
		}
	case "subtract":
		return func(i1, i2 int) int {
			return i1 - i2
		}
	default:
		return func(i1, i2 int) int {
			return i1 + i2
		}
	}
}
