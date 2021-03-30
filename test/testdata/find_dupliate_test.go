package duplicateFileFinder

import (
	duplicateFileFinder "DanSu93/duplicate-file-finder/cmd/duplicate-file-finder"
	"bufio"
	"log"
	"os"
	"testing"
	"time"
)

func createTestData(data string) error {
	err := os.MkdirAll("./test/data/", 0777)
	if err != nil {
		return err
	}

	firstFile, err := os.Create("./test/data/test.txt")
	if err != nil {
		return err
	}
	defer firstFile.Close()

	time.Sleep(time.Second * 1)
	secondFile, err := os.Create("./test/data/test.txt")
	if err != nil {
		return err
	}
	defer secondFile.Close()

	if data != "" {
		w := bufio.NewWriter(secondFile)
		_, err = w.WriteString(data)
		if err != nil {
			return err
		}

		w.Flush()
	}
	return nil
}

func TestStart(t *testing.T) {
	tests := []struct {
		name   string
		delete bool
		data   string
		expect uint8
	}{
		{name: "delete no duplicates", delete: true, data: "test", expect: 0},
		{name: "delete with duplicates", delete: true, data: "", expect: 1},
		{name: "no delete no duplicates", delete: false, data: "test", expect: 0},
		{name: "no delete with duplicates", delete: false, data: "", expect: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				err := os.RemoveAll("./test/data/")
				if err != nil {
					log.Fatalf("cannot remove test data: %v\n", err)
				}
			}()

			err := createTestData(tc.data)
			if err != nil {
				log.Println(err)
				return
			}

			err = duplicateFileFinder.CheckDuplicate("./test/data/", tc.delete)
			if err != nil {
				log.Println(err)
				return
			}

			if tc.expect == 1 {
				if _, err := os.Stat("./test/data/test.txt"); err == nil {
					log.Println("file was not deleted")
					return
				}
			}
			if tc.expect == 0 {
				if _, err := os.Stat("./test/data/test.txt"); err != nil {
					log.Println("cannot find file")
					return
				}
				if _, err := os.Stat("./test/data/test.txt"); err != nil {
					log.Println("cannot find file")
					return
				}
			}
		})
	}
}
