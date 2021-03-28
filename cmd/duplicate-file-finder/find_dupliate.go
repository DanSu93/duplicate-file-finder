// Package provides duplicate file search and duplicate file removal operation
// Support operations: CheckDuplicate, deleteDuplicate

package duplicateFileFinder

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// CheckDuplicate searches for duplicate files in the specified directory
func CheckDuplicate(dir string, delete bool) error{
	var (
		compareFiles   = make(map[[sha512.Size]byte]string)
		duplicateFiles []string
	)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if f.IsDir() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		hash := sha512.Sum512(data)

		if _, ok := compareFiles[hash]; ok {
			duplicateFiles = append(duplicateFiles, path)
		} else {
			compareFiles[hash] = path
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	if len(duplicateFiles) == 0 {
		fmt.Println("There are no duplicate files")
		return nil
	}

	if len(duplicateFiles) > 0 && delete {
		deleteDuplicate(duplicateFiles)
	} else {
		fmt.Println("List of duplicate files:")
		for _, val := range duplicateFiles {
			fmt.Println(val)
		}
	}
	return nil
}

// deleteDuplicate removes duplicate files found in the specified directory
func deleteDuplicate(files []string) {
	var (
		err error
		wg  sync.WaitGroup
	)

	fmt.Println("List of deleted duplicate files:")
	for _, val := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err = os.Remove(val)
			if err != nil {
				panic(err)
			}
			fmt.Println(val)
		}()
		wg.Wait()
	}
}
