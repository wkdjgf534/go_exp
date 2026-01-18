package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("letters.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := countAlphabets(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Letters from file: %d\n", n)

	str := strings.NewReader("Hello Worldddd*****")

	n, err = countAlphabets(str)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Letters from string: %d\n", n)
}

func countAlphabets(r io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)

	for {
		n, err := r.Read(buffer)
		for _, l := range buffer[:n] {
			if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
				count++
			}
		}
		if err == io.EOF {
			return count, nil
		}
		if err != nil {
			return 0, err
		}
	}
}
