package duplicate_file_finder

import (
	"fmt"
	"io/ioutil"
)

func DeleteDuplicateFiles(dir string) []string {
	return []string{}
}

func FindAllFiles(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Sys())
	}
}
