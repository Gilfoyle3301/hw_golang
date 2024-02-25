package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type resultSlice struct {
	words string
	count uint
}

func Top10(s string) []string {
	var (
		rSlice  = make([]resultSlice, 0, len(strings.Fields(s)))
		result  = make([]string, 0, 10)
		forSort = make(map[string]uint, 0)
		counts  uint
	)
	for _, v := range strings.Fields(s) {
		if _, ok := forSort[v]; ok {
			forSort[v]++
			continue
		}
		forSort[v] = counts + 1
	}

	for k, v := range forSort {
		rSlice = append(rSlice, resultSlice{k, v})
	}
	sort.Slice(rSlice, func(i, j int) bool {
		if rSlice[i].count == rSlice[j].count {
			return rSlice[i].words < rSlice[j].words
		}
		return rSlice[i].count > rSlice[j].count
	})
	for i, v := range rSlice {
		if i < 10 {
			result = append(result, v.words)
		}
	}
	return result
}
