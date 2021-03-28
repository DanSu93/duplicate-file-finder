//This program searches for and removes duplicates in the specified directory

package main

import (
	duplicateFileFinder "DanSu93/duplicate-file-finder/cmd/duplicate-file-finder"
	"flag"
	"os"
)

var (
	dir           *string
	del           *bool
	currentDir, _ = os.Getwd()
)

func init() {
	dir = flag.String("dir", currentDir, "directory in which duplicates are searched")
	del = flag.Bool("del", false, "remove duplicates")
	flag.Parse()
}

func main() {
	duplicateFileFinder.CheckDuplicate(*dir, *del)
}
