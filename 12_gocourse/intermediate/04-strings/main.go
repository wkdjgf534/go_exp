package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message1 := "Hello, \nGo!"
	message2 := "Hello, \tGo!" // Hello,  Go!
	message3 := "Hello, \rGo!" // Go!lo,
	rawMessage := `Hello\nGo`  // Hello\nGo - escape sequence doesn't work for string with backticks

	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)
	fmt.Println(rawMessage)

	fmt.Println("Length of message1 variable is", len(message1))
	fmt.Println("Length of rawMessage1 variable is", len(rawMessage))

	fmt.Println("Print the first character in message variable:", message1[0]) // ASCII
	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"          // A has an ASCII value of 65
	str2 := "banana"         // b has an ASCII value of 98
	str3 := "app"            // a has an ASCII value 97
	str4 := "apple"          // a has an ASCII value of 97
	fmt.Println(str1 < str2) // true according to lexical graphical comparison
	fmt.Println(str3 < str1) // false
	fmt.Println(str4 > str1) // true
	fmt.Println(str4 > str3) // true apple > apple, here the lentgh of strings

	for _, char := range message1 {
		// %d, %c - placeholders in C lang, format verbs in Go lang
		//fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))
	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	jch := '日'
	fmt.Println(ch)  // 97
	fmt.Println(jch) // 26085

	fmt.Printf("%c\n", ch)  // a
	fmt.Printf("%c\n", jch) // 日

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日本語" // Japanese text
	fmt.Println(NIHONGO)

	jhello := "こんにちは" // Japanese "Hello"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}
}
