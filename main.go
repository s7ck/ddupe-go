package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Make sure the path provided is valid and the file exists...
func removeDuplicates(args map[string]interface{}) {
	// Convert the paths, which came in as interfaces, to strings.
	inFilePath := fmt.Sprintf("%s", args["inFile"])
	outFilePath := fmt.Sprintf("%s", args["outFile"])

	// Convert the bools that came in as interfaces to strings...
	sortOutputStr := fmt.Sprintf("%t", args["sort"])
	ignoreBlanksStr := fmt.Sprintf("%t", args["ignoreBlanks"])
	// replaceInFileStr := fmt.Sprintf("%t", args["replaceFile"])

	//... then convert them to bools.
	sortOutput, _ := strconv.ParseBool(sortOutputStr)
	ignoreBlanks, _ := strconv.ParseBool(ignoreBlanksStr)
	// replaceInFile, _ := strconv.ParseBool(replaceInFileStr)

	if !pathIsValid(inFilePath) {
		//... if not, show an error and exit.
		log.Fatalf("The source file '%v' does not exist.", inFilePath)
	} else {
		log.Println("Opening source file...")

		sourceFile, err := os.Open(inFilePath)
		defer sourceFile.Close()

		handleError(err)

		uniqueLines := handleSourceFile(sourceFile)

		if sortOutput {
			log.Println("Sorting output...")
			sort.Strings(uniqueLines)
		}

		if args["outputPath"] != "" {
			handleOutFile(outFilePath, uniqueLines, ignoreBlanks)
		} else {
			for _, val := range uniqueLines {
				fmt.Println(val)
			}
		}
		// uniqueSet := handleSourceFile(sourceFile)
	}
}

// TODO -> Handle invalid CLI flags.
// TODO -> Support stdin
func main() {
	args := make(map[string]interface{})

	// Setup the args we want to support.
	inFilePtr := flag.String("source", "", "Path to the source file from which duplicates will be removed.")
	outFilePtr := flag.String("output", "", "Path to an output file of your choosing. The file does not need to exist.")
	sortPtr := flag.Bool("sort", false, "Indicates if you want to sort the output (ascending)")
	ignoreBlanksPtr := flag.Bool("ignore-blank", false, "Ignores blank lines in the source file")
	replaceFilePtr := flag.Bool("replace", false, "Replace source file with output")

	// Now we can parse those args.
	flag.Parse()

	args["inFile"] = *inFilePtr
	args["outFile"] = *outFilePtr
	args["sort"] = *sortPtr
	args["ignoreBlanks"] = *ignoreBlanksPtr
	args["replaceFile"] = *replaceFilePtr

	removeDuplicates(args)
}
