package controllers

import (
	"fmt"
	service "parallel-sorting-go-/app/services"
	"parallel-sorting-go-/app/views"
)

func MainController() {
	n, files := views.ArgReader()

	if n == 0 || len(files) == 0 {
		return
	}

	fmt.Printf("N threads: %d\n", n)
	fmt.Print("Files: \n")
	for idx, val := range files {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	allNumbers := ParallelReadingController(files)
	chunks := service.NumbersSeparator(allNumbers, n)
	fmt.Println(chunks)
	ParallelSortingController(chunks, n)
}
