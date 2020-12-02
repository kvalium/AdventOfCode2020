package main

import (
	"errors"
	"fmt"
)

func main() {
	rFirst, eFirst := oneTermTo2020()
	if eFirst != nil {
		fmt.Println("first exercice failed:", eFirst)
		return
	}
	fmt.Println("first exercise: ", rFirst)

	rSecond, eSecond := twoTermsTo2020()
	if eSecond != nil {
		fmt.Println("second exercice failed:", eSecond)
		return
	}
	fmt.Println("second exercise: ", rSecond)
}

func oneTermTo2020() (int, error) {
	for _, n := range Numbers {
		rest := 2020 - n
		if contains(Numbers, rest) {
			return n * rest, nil
		}
	}
	return -1, errors.New("no match")
}

func contains(nn []int, test int) bool {
	for _, n := range nn {
		if n == test {
			return true
		}
	}
	return false
}

func twoTermsTo2020() (int, error) {
	for _, n := range Numbers {
		rest := 2020 - n
		inferiorToRest := func(i int) bool { return i < rest }
		restCandidates := filter(Numbers, inferiorToRest)
		for _, restA := range restCandidates {
			for _, restB := range restCandidates {
				if restA+restB == rest {
					return n * restA * restB, nil
				}
			}
		}
	}
	return -1, errors.New("no match")
}

func filter(nn []int, test func(int) bool) (ret []int) {
	for _, n := range nn {
		if test(n) {
			ret = append(ret, n)
		}
	}
	return
}
