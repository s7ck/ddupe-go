package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

// TODO -> Handle invalid CLI flags.
// TODO -> Support stdin
func main() {
	// Setup the args we want to support.
	inFilePtr := flag.String("source", "", "Path to the source file from which duplicates will be removed.")
	outFilePtr := flag.String("output", "", "Path to an output file of your choosing. The file does not need to exist.")
	sortPtr := flag.Bool("sort", false, "Indicates if you want to sort the output (ascending)")
	ignoreBlanksPtr := flag.Bool("ignore-blank", false, "Ignores blank lines in the source file")
	// replaceFilePtr := flag.Bool("replace", false, "Replace source file with output")

	// Now we can parse those args.
	flag.Parse()

	// Make sure the path provided is valid and the file exists...
	if !pathIsValid(inFilePtr) {
		//... if not, show an error and exit.
		log.Fatalf("The source file '%v' does not exist.", *inFilePtr)
	} else {
		sourceFile, err := os.Open(*inFilePtr)
		defer sourceFile.Close()

		handleError(err)

		fmt.Println("Reading the source file...")

		uniqueLines := handleSourceFile(sourceFile)

		if *sortPtr {
			sort.Strings(uniqueLines)
		}

		if *outFilePtr != "" {
			handleOutFile(outFilePtr, uniqueLines, *ignoreBlanksPtr)
		} else {
			for _, val := range uniqueLines {
				fmt.Println(val)
			}
		}
		// uniqueSet := handleSourceFile(sourceFile)
	}
}
