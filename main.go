package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func getUniqueLinesFromFile(sourceFile *os.File) map[string]struct{} {
	// We need somewhere to store the unique lines from the source file so we're
	// using a map for it. This is a map of strings with structs as values. The
	// structs will always be empty for two reasons:
	//   1. No value is actually needed. The key is sufficient.
	//   2. An empty struct takes up no memory.
	uniqueLines := make(map[string]struct{})

	numSourceLines := 0
	scanner := bufio.NewScanner(sourceFile)

	for scanner.Scan() {
		numSourceLines++

		thisLine := scanner.Text()

		if _, ok := uniqueLines[thisLine]; !ok {
			uniqueLines[thisLine] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d of %d lines were unique", len(uniqueLines), numSourceLines)

	return uniqueLines
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func validateArgs(argMap map[string]struct{}) bool {
	return true
}

// TODO -> Handle invalid CLI flags.
func main() {
	// TODO -> Support stdin

	// Make and parse the args we want to caputre.
	// filePathPtr := flag.String("path", "", "Path to the file from which duplicates will be removed")
	filePathPtr := flag.String("source", "", "Path to the source file from which duplicates will be removed")

	// sortPtr := flag.Bool("sort", false, "Indicates if you want to sort the output (ascending)")
	// replaceFilePtr := flag.Bool("replace", false, "Replace source file with output")
	// outputFilePtr := flag.String("output", "", "Output file")

	// Do the actual arg parsing.
	flag.Parse()

	if pathIsValid(*filePathPtr) {
		fmt.Println("Opening " + *filePathPtr)
		sourceFile, err := os.Open(*filePathPtr)

		if err != nil {
			panic(err)
		}

		// if *sortPtr {
		// 	fmt.Println("Sorting")
		// }

		// if *replaceFilePtr {
		// 	fmt.Println("Replacing source file")
		// } else {
		// 	if *outputFilePtr == "" {
		// 		fmt.Println("You must specify -ouput if you're not using -replace.")
		// 	}
		// }

		if err != nil {
			log.Fatal(err)
		}

		defer sourceFile.Close()

		// uniqueSet := getUniqueLinesFromFile(sourceFile)
		getUniqueLinesFromFile(sourceFile)
	} else {
		fmt.Printf("The source file '%v' does not exist.", *filePathPtr)
	}
}
