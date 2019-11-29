package shortt

import (
	"math/rand"
)

func QuickShortTypeInt(val []int) []int {
	// Check if total array < 2
	if len(val) < 2 {
		return val
	}
	// Getting value left, right and pivot
	left, right, pivot := 0, len(val)-1, rand.Int()%len(val)
	val[pivot], val[right] = val[right], val[pivot]

	for i, _ := range val {
		if val[i] < val[right] {
			val[left], val[i] = val[i], val[left]
			left++
		}
	}

	val[left], val[right] = val[right], val[left]
	QuickShortTypeInt(val[:left])
	QuickShortTypeInt(val[left+1:])
	return val
}
