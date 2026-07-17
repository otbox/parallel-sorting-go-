package controllers

import service "parallel-sorting-go-/app/services"

func process(id int, chunk []int) {
	service.Execute(id, func() []int {
		service.HeapSort(chunk)
		return chunk
	})
}

func ParallelSortingController(chunks [][]int, nCouroutines int) {
	service.GetGroup().Add(nCouroutines)

	for idx, chunk := range chunks {
		process(idx, chunk)
	}

}
