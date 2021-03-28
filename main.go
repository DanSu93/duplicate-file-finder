package main

import (
	duplicatefilefinder "DanSu93/duplicate-file-finder/cmd/duplicate-file-finder"
	"flag"
	"os"
)

var (
	dir           *string
	delete        *bool
	currentDir, _ = os.Getwd()
)

func init() {
	dir = flag.String("dir", currentDir, "directory in which duplicates are searched")
	delete = flag.Bool("delete", false, "remove duplicates")
	flag.Parse()
}

func main() {
	if *delete {
		duplicatefilefinder.DeleteDuplicateFiles(*dir)
	} else {
		duplicatefilefinder.FindAllFiles(*dir)
	}
}
