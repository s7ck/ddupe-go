package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleOutFile(outputPath *string, uniqueLines []string) {
	outputFile, err := os.OpenFile(*outputPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer outputFile.Close()
	handleError(err)

	writer := bufio.NewWriter(outputFile)

	for _, value := range uniqueLines {
		writer.WriteString(value + "\n")
	}

	writer.Flush()
}

// func handleSourceFile(sourceFile *os.File) map[string]struct{} {
func handleSourceFile(sourceFile *os.File) []string {
	// We need somewhere to store the unique lines from the source file so we're
	// using a map for it. This is a map of strings with structs as values. The
	// structs will always be empty for two reasons:
	//   1. No value is actually needed. The key is sufficient.
	//   2. An empty struct takes up no memory.
	uniqueLines := make(map[string]struct{})
	orderedSlice := []string{}

	numSourceLines := 0
	scanner := bufio.NewScanner(sourceFile)

	for scanner.Scan() {
		numSourceLines++

		thisLine := scanner.Text()

		// Check if the current line is already in the unique list...
		if _, ok := uniqueLines[thisLine]; !ok {
			//... if not, add it with the empty struct.
			uniqueLines[thisLine] = struct{}{}
			// And then add the known unique line to the slice.
			orderedSlice = append(orderedSlice, thisLine)
		}
	}

	fmt.Printf("%d of %d lines were unique.\n", len(uniqueLines), numSourceLines)

	// return uniqueLines
	return orderedSlice
}
