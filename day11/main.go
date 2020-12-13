package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var seatMap [][]rune
var nbRows, nbCols int
var occSeat = '#'
var freeSeat = 'L'
var floor = '.'

type seat struct {
	x, y int
}

func main() {
	start := time.Now()
	setSeatMap("./example")
	origSeatMap := make([][]rune, len(seatMap))
	for i := range seatMap {
		origSeatMap[i] = make([]rune, len(seatMap[i]))
		copy(origSeatMap[i], seatMap[i])
	}
	placePassengers()
	fmt.Println("First exercice:", countOccupiedSeats())
	r := checkForOccupiedSeat(&seat{x: 0, y: 0}, seat{x: 0, y: 0})
	fmt.Println(r)
	elapsed := time.Since(start)
	fmt.Println("exec. time:", elapsed)
}

func placePassengers() {
	changes := 1
	rounds := 0
	// fmt.Println("Placing passengers...")
	for changes > 0 {
		changes = addNewPassengersWave()
		rounds++
		// fmt.Printf("* round %v (%v changes)\n", rounds, changes)
	}
	// fmt.Println("done")
}

func countOccupiedSeats() (occupiedSeats int) {
	for r := 0; r < nbRows; r++ {
		for _, state := range seatMap[r] {
			if state == occSeat {
				occupiedSeats++
			}
		}
	}
	return
}

func addNewPassengersWave() (stateChanges int) {
	newSeatMap := make([][]rune, len(seatMap))
	for i := range seatMap {
		newSeatMap[i] = make([]rune, len(seatMap[i]))
		copy(newSeatMap[i], seatMap[i])
	}

	for r := 0; r < nbRows; r++ {
		for c := 0; c < nbCols; c++ {
			seat := seat{x: c, y: r}
			if seatMap[r][c] == floor {
				continue
			}
			newState, change := newSeatState(seat)
			if change {
				stateChanges++
				newSeatMap[r][c] = newState
			}
		}
	}
	seatMap = newSeatMap
	return
}

func setSeatMap(inputFilePath string) {
	f, err := os.Open(inputFilePath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(f)
	for {
		seatRow, err := rd.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		var cleanRow []rune
		for _, c := range []rune(seatRow) {
			if c == '\n' {
				continue
			}
			cleanRow = append(cleanRow, c)
		}
		seatMap = append(seatMap, cleanRow)
	}
	nbRows = len(seatMap)
	nbCols = len(seatMap[0])
}

func countAdjacentOccupiedSeat(s seat) (ct int) {
	positions := []seat{
		{x: s.x, y: s.y - 1},     // TOP
		{x: s.x + 1, y: s.y - 1}, // TOP-RIGHT
		{x: s.x + 1, y: s.y},     // RIGHT
		{x: s.x + 1, y: s.y + 1}, // BOTTOM-RIGHT
		{x: s.x, y: s.y + 1},     // BOTTOM
		{x: s.x - 1, y: s.y + 1}, // BOTTOM-LEFT
		{x: s.x - 1, y: s.y},     // LEFT
		{x: s.x - 1, y: s.y - 1}, // TOP-LEFT
	}
	for _, p := range positions {
		if isOccupied(p) {
			ct++
		}
	}
	return
}

func checkForOccupiedSeat(s *seat, dir seat) bool {
	checkSeat := seat{x: s.x + dir.x, y: s.y + dir.y}
	fmt.Printf("check seat %+v\n", checkSeat)
	fmt.Println("VALID", isValid(checkSeat), "ISFLOOR", isFloor(checkSeat), "ISOCCUPIED", isOccupied(checkSeat))

	if !isValid(checkSeat) {
		return false
	}
	if isFloor(checkSeat) {
		return checkForOccupiedSeat(&checkSeat, dir)
	}
	return isOccupied(checkSeat)
}

func newSeatState(s seat) (newState rune, hasSwitched bool) {
	nbAdjacentOccupiedSeats := countAdjacentOccupiedSeat(s)
	currentSeatState := seatMap[s.y][s.x]
	if currentSeatState == freeSeat && nbAdjacentOccupiedSeats == 0 {
		return occSeat, true
	}
	if currentSeatState == occSeat && nbAdjacentOccupiedSeats >= 4 {
		return freeSeat, true
	}
	return currentSeatState, false
}

func isOccupied(s seat) bool {
	if !isValid(s) {
		return false
	}
	return seatMap[s.y][s.x] == occSeat
}

func isFloor(s seat) bool {
	return seatMap[s.y][s.x] == floor
}

func isValid(s seat) bool {
	return s.x >= 0 && s.y >= 0 && s.y < nbRows-1 && s.x < nbCols-1
}
