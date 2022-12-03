package main

import (
	"fmt"
	"github.com/vplemyannik/wordcount/app"
	"os"
)

func main() {
	src, err := readInput()
	if err != nil {
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%d", app.Calculate(src))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	src = os.Args[1]
	return src, nil
}
