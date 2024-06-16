package main

import (
	"fmt"
	"os"
)

// Function to append data to a file
func appendToFile(filename string, data string) {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Append the data
	if _, err := file.WriteString(data + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Data successfully appended to file")
	}
}

func main() {
	// Data to append
	data := "This is the status data to be appended."
	println("heelwo")
	// Append the data to status.txt
	appendToFile("status.txt", data)
}
