package service

func NumbersSeparator(allNumbers []int, numChunks int) [][]int {
	length := len(allNumbers)

	if length == 0 || numChunks <= 0 {
		panic("Numbers Separator error.")
	}

	if numChunks > length {
		numChunks = length
	}

	baseChunkSize := length / numChunks
	rest := length % numChunks
	chunks := make([][]int, numChunks)

	start := 0
	for i := 0; i < numChunks; i++ {
		chunkSize := baseChunkSize
		if i < rest {
			chunkSize++
		}

		end := start + chunkSize

		chunks[i] = allNumbers[start:end]
		start = end
	}
	return chunks
}
