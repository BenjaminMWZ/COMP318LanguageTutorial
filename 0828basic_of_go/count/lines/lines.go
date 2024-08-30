package lines

import (
	"bufio"
	"os"
)

func CountLines(file *os.File) (int, error) {
	var count int

	// Initialize a new scanner
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	for scanner.Scan() {
		count++
	}

	// Check for errors
	err := scanner.Err()
	if err != nil {
		return 0, err
	}

	return count, nil

}
