package main

import (
	"flag"
	"fmt"
	"os"
)

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

func main() {
	// TODO -> Support stdin

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

	// reader := bufio.NewReader()
}
