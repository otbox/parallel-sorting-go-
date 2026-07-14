package views

import (
	"fmt"
	"os"
	"strconv"
)

func ArgReader() (nThreads int, filesPath []string) {
	if len(os.Args) < 3 {
		fmt.Println("Is required least 3 argumnets")
		return 0, nil
	}

	value, err := strconv.Atoi(os.Args[1])

	if err != nil {
		return 0, nil
	}

	nThreads = value
	filesPath = append(filesPath, os.Args[2:]...)
	return
}
