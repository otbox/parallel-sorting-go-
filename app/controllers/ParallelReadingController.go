package controllers

import (
	"fmt"
	service "parallel-sorting-go-/app/services"
)

func processReading(id int, allNumbers *[]int, file string) {
	numbers := service.Execute(id, func() []int { return service.Reader(file) })

	*allNumbers = append(*allNumbers, numbers...)
}

func ParallelReadingController(files []string) []int {
	service.InitGroup(len(files))

	var allNumbers []int

	fmt.Print("1. Paralallel Reading Files: \n")
	for idx, path := range files {
		processReading(idx, &allNumbers, path)
	}

	return allNumbers
}
