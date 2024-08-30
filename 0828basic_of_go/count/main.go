package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BenjaminMWZ/count/lines"
)

func main() {
	// Use flag package to parse command line arguments
	// The command looks like: go run main.go -file <filename> lines/lines.go

	var filename string
	flag.StringVar(&filename, "file", "", "The name of the file to count lines in")
	flag.Parse()

	if filename == "" {
		fmt.Println("You must provide a filename")
		return
	}

	// Open the file
	file, err := os.Open(filename)

	// Check for errors
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Close the file when we're done
	defer file.Close()

	// Call the CountLines function from the lines package
	count, err := lines.CountLines(file)
	if err != nil {
		fmt.Println("Error counting lines:", err)
		return
	}
	fmt.Println(filename, count)
}
