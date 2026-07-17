package controllers

import (
	"fmt"
	service "parallel-sorting-go-/app/services"
)

func process(id int, chunk []int) {
	service.Execute(id, func() []int {
		service.HeapSort(chunk)
		return chunk
	})
}

func ParallelSortingController(chunks [][]int, nCouroutines int) {
	service.GetGroup().Add(nCouroutines)

	fmt.Print("2. Parallel Heap Sorting: \n")
	for idx, chunk := range chunks {
		process(idx, chunk)
	}

	merged := service.KWayMerge(chunks)
	fmt.Println(merged)
}
