package service

// type chunkTable struct {
// 	id int
// 	value int
// 	idx int
// }

// func KWayMerge (chunks [][]int) {
// 	var chunkIndexTables []chunkTable
// 	for x := 0; x < len(chunks); x++ {
// 		chunkIndexTables = append(chunkIndexTables, chunkTable{id: x, value: chunks[x][0], idx: 0})
// 	}

// 	var mergedFinalArray []int

// 	minorPos := 0
// 	for x := 0; x < len(chunkIndexTables); x++ {
// 		if (chunkIndexTables[minorPos].value > chunkIndexTables[x].value) {
// 			minorPos = x
// 		}
// 	}
// 	minEntry := chunkIndexTables[minorPos]
// 	mergedFinalArray = append(mergedFinalArray, minEntry.value)
// 	chunkIndexTables[minorPos].idx++
// }

type Node struct {
	value   int
	chunkID int
	index   int
}

type MinHeap struct {
	data []Node
}

func (h *MinHeap) Len() int {
	return len(h.data)
}

// When i add somthing will run up
func (h *MinHeap) Insert(node Node) {
	h.data = append(h.data, node)
	h.up(h.Len() - 1)
}

func (h *MinHeap) up(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[parent].value <= h.data[i].value {
			break
		}
		h.data[parent], h.data[i] = h.data[i], h.data[parent]
		i = parent
	}
}

func (h *MinHeap) down(i int) {
	n := h.Len()

	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && h.data[left].value < h.data[smallest].value {
			smallest = left
		}

		if right < n && h.data[right].value < h.data[smallest].value {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

func (h *MinHeap) extractMin() (Node, bool) {
	if h.Len() == 0 {
		return Node{}, false
	}

	min := h.data[0]
	last := h.data[h.Len()-1]
	h.data = h.data[:h.Len()-1]

	if h.Len() > 0 {
		h.data[0] = last
		h.down(0)
	}

	return min, true
}

func KWayMerge(chunks [][]int) []int {
	var heap MinHeap
	var merged []int

	// Initialization with the firsts elements
	for chunkID := 0; chunkID < len(chunks); chunkID++ {
		if len(chunks[chunkID]) > 0 {
			heap.Insert(Node{
				value:   chunks[chunkID][0],
				chunkID: chunkID,
				index:   0,
			})
		}
	}

	for heap.Len() > 0 {
		minNode, _ := heap.extractMin()
		merged = append(merged, minNode.value)

		nextIndex := minNode.index + 1
		if nextIndex < len(chunks[minNode.chunkID]) {
			heap.Insert(Node{
				value:   chunks[minNode.chunkID][nextIndex],
				chunkID: minNode.chunkID,
				index:   nextIndex,
			})
		}
	}

	return merged
}
