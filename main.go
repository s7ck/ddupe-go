package main

import (
	"flag"
	"fmt"
	"os"
)

func getUniqueLinesFromFile(sourceFile) map {
	scanner := bufio.NewScanner(sourceFile)

	for scanner.Scan() {
		numSourceLines++

		thisLine := scanner.Text()

		if _, ok := uniqueLines[thisLine]; !ok {
			uniqueLines[thisLine] = struct{}{}
		}
	}
}
//testing git
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func validateArgs(argMap map) bool {
	return true
}

func main() {
	// TODO -> Support stdin

	// We need somewhere to store the unique lines from the source file so we're
	// using a map for it. This is a map of strings with structs as values. The
	// structs will always be empty for two reasons:
	//   1. No value is actually needed. The key is sufficient.
	//   2. An empty struct takes up no memory.
	uniqueLines := make(map[string]struct{})
	numSourceLines := 0

	// Make and parse the args we want to caputre.
	filePathPtr := flag.String("path", "", "Path to the file from which duplicates will be removed")
	sortPtr := flag.Bool("sort", false, "Indicates if you want to sort the output (ascending)")
	replaceFilePtr := flag.Bool("replace", false, "Replace source file with output")
	outputFilePtr := flag.String("output", "", "Output file")

	// Do the actual arg parsing.
	flag.Parse()

	fmt.Println("Opening " + *filePathPtr)

	if fileExists(*filePathPtr) {
		file, err := os.Open(*filePathPtr)
	}

	check(err)

	if *sortPtr {
		fmt.Println("Sorting")
	}

	if *replaceFilePtr {
		fmt.Println("Replacing source file")
	} else {
		if *outputFilePtr == "" {
			fmt.Println("You must specify -ouput if you're not using -replace.")
		}
	}

	sourceFile, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer sourceFile.Close()

	uniqueSet := getUniqueLinesFromFile(sourceFile)

	fmt.Printf("%d of %d lines were unique", len(uniqueLines), numSourceLines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
