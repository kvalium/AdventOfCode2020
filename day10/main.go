package main

import (
	"fmt"
	"sort"
)

func main() {
	d := getDiffBetweenAdaptersT()
	fmt.Println("First exercice:", d[1]*d[3])
	arr := getAllPossibleArrangements()
	fmt.Println("Second exercice:", arr)
}

func getDiffBetweenAdaptersT() map[int]int {
	var differences = map[int]int{1: 1, 2: 1, 3: 1}
	sort.Ints(adaptersT)
	for i, a := range adaptersT {
		if i+1 == len(adaptersT) {
			continue
		}
		diff := adaptersT[i+1] - a
		differences[diff] = differences[diff] + 1
	}
	return differences
}

func getAllPossibleArrangements() (nbArr int) {
	sort.Ints(adaptersT)
	return walkAdaptersT(adaptersT, 0, 1)
}

func walkAdaptersT(adaptersT []int, i, nb int) int {
	curr := adaptersT[i]
	if inSlice(adaptersT, curr+1) {
		nb++
		nb += walkAdaptersT(adaptersT, i+1, nb)
	}
	if inSlice(adaptersT, curr+2) {
		nb++
		nb += walkAdaptersT(adaptersT, i+2, nb)
	}
	if inSlice(adaptersT, curr+3) {
		nb++
		nb += walkAdaptersT(adaptersT, i+3, nb)
	}
	return nb
}

func inSlice(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
