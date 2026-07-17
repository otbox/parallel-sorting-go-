package service

// func up (heap []int, i int) {
// 	n := len(heap)

//     for {
// 		dad := (i - 1) / 2
// 		if dad < n && heap[dad] < heap[i] {
// 			temp := heap[dad]
// 			heap[dad] = heap[i]
// 			heap[i] = temp
// 		} else {
// 			break
// 		}
// 	}

// }

func up(heap []int, i int) {
	for i > 0 {
		dad := (i - 1) / 2

		if heap[dad] < heap[i] {
			heap[dad], heap[i] = heap[i], heap[dad]
			i = dad
		} else {
			break
		}
	}
}

func down(heap []int, i int) {
	n := len(heap)

	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < n && heap[left] > heap[largest] {
			largest = left
		}
		if right < n && heap[right] > heap[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		heap[i], heap[largest] = heap[largest], heap[i]
		i = largest
	}
}
func buildHeap(heap []int) {
	IndexOfLastDad := len(heap)/2 - 1
	for x := IndexOfLastDad; x >= 0; x-- {
		down(heap, x)
	}
}

func HeapSort(heap []int) {
	buildHeap(heap)

	size := len(heap) - 1
	for ; size > 0; size-- {
		heap[0], heap[size] = heap[size], heap[0]
		down(heap[:size], 0)
	}
}
