package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("writing.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := writeString("Hello World!!!", file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Written bytes: %d\n", n)
}

func writeString(s string, w io.Writer) (int, error) {
	n, err := w.Write([]byte(s))
	if err != nil {
		return 0, fmt.Errorf("error occurred while writing: %w", err)
	}

	return n, nil
}
