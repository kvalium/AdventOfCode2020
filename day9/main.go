package main

import "fmt"

func main() {
	unmat := getFirstUnmatching()
	fmt.Println("First exercise:", unmat)
	t := getXMasCodeBreak(unmat)
	fmt.Println(t)
}

func getXMasCodeBreak(t int) int {
	for i, n := range xMasOutput {
		ct := n
		r := []int{xMasOutput[i]}
		for j := i + 1; j < len(xMasOutput); j++ {
			ct += xMasOutput[j]
			if ct > t {
				break
			}
			r = append(r, xMasOutput[j])
			if ct == t {
				min, max := minMax(r)
				return min + max
			}

		}
	}
	return -1
}

func getFirstUnmatching() int {
	for i, n := range xMasOutput {
		if i < xMasPreambleSize {
			continue
		}
		stIdx := i - xMasPreambleSize
		found := false
		for j := 0; j < xMasPreambleSize; j++ {
			if found {
				break
			}
			for k := j + 1; k < xMasPreambleSize; k++ {
				a := xMasOutput[stIdx+j]
				b := xMasOutput[stIdx+k]
				if a+b == n {
					found = true
				}
			}
		}
		if !found {
			return n
		}
	}
	return -1
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
