package controllers

import (
	"fmt"
	service "parallel-sorting-go-/app/services"
)

func processReading(id int, allNumbers *[]int, file string) {
	numbers := service.Execute(id, func() []int { return service.Reader(file) })

	*allNumbers = append(*allNumbers, numbers...)
}

func ParallelReadingController(files []string) {
	service.InitGroup(len(files))

	var allNumbers []int

	for idx, path := range files {
		processReading(idx, &allNumbers, path)
	}

	fmt.Println(allNumbers)

	// service.
}
