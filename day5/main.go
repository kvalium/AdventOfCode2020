package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
)

var seatSplitRegex = regexp.MustCompile(`([B-F]{7})([L-R]{3})`)
var seatIDMultiplier = 8

func main() {
	fmt.Println("First exercise:", getMaxSeatID())
	fmt.Println("Second exercise:", getFirstAvailableSeat())
}

func getMaxSeatID() (max int) {
	seatIDS := getSeatIDs()
	return seatIDS[len(seatIDS)-1]
}

func getFirstAvailableSeat() (seatID int) {
	seatIDS := getSeatIDs()
	startSeat := seatIDS[0]
	for i, seatID := range seatIDS {
		if startSeat+i != seatID {
			return startSeat + i
		}
	}
	return -1
}

func getSeatIDs() (seatIDs []int) {
	for _, seat := range seats {
		seatIDs = append(seatIDs, getSeatID(seat))
	}
	sort.Ints(seatIDs)
	return
}

func getSeatID(seat string) int {
	s := seatSplitRegex.FindStringSubmatch(seat)
	row := getRow(s[1])
	column := getColumn(s[2])
	return row*seatIDMultiplier + column
}

func getRow(r string) int {
	return codeToInt(r, 'B')
}

func getColumn(c string) int {
	return codeToInt(c, 'R')
}

func codeToInt(cc string, upper rune) (r int) {
	l := len(cc) - 1
	for i, c := range []rune(cc) {
		if c == upper {
			pow := float64(l - i)
			r += int(math.Pow(2, pow))
		}
	}
	return
}
