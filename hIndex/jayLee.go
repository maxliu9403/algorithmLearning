package hIndex

import (
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	var h int
	for i := range citations {
		if citations[i] > h {
			h++
		}
	}
	return h
}
