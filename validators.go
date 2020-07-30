package main

import "os"

func pathIsValid(path *string) bool {
	if file, err := os.Stat(*path); err == nil {
		// The file exists, but only return true if it's not a directory.
		return !file.IsDir()
	}

	return false
}
