package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	diffs, sequences := getDiffBetweenAdapters()
	fmt.Println("First exercice:", diffs[1]*diffs[3])
	a := math.Pow(7, sequences[4])
	b := math.Pow(4, sequences[3])
	c := math.Pow(2, sequences[2])
	fmt.Println("Second exercice:", int(a*b*c))
}

func getDiffBetweenAdapters() (map[int]int, map[int]float64) {
	var differences = map[int]int{1: 1, 2: 1, 3: 1}
	var sequencesOfOne = map[int]float64{
		1: 0, 2: 1, 3: 0, 4: 0,
	}
	sort.Ints(adapters)
	var diffSeq = 0
	for i, a := range adapters {
		if i+1 == len(adapters) {
			continue
		}
		diff := adapters[i+1] - a
		differences[diff] = differences[diff] + 1
		if diff == 3 {
			sequencesOfOne[diffSeq] = sequencesOfOne[diffSeq] + 1
			diffSeq = 0
			continue
		}
		diffSeq++
	}
	return differences, sequencesOfOne
}
