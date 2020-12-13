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

var positions = []seat{
	{x: 0, y: -1},  // TOP
	{x: 1, y: -1},  // TOP-RIGHT
	{x: 1, y: 0},   // RIGHT
	{x: 1, y: 1},   // BOTTOM-RIGHT
	{x: 0, y: 1},   // BOTTOM
	{x: -1, y: 1},  // BOTTOM-LEFT
	{x: -1, y: 0},  // LEFT
	{x: -1, y: -1}, // TOP-LEFT
}

type seat struct {
	x, y int
}

func main() {
	start := time.Now()
	setSeatMap("./seats")
	placePassengers(adjacentOnly)
	fmt.Println("First exercice:", countOccupiedSeats())

	setSeatMap("./seats")
	placePassengers(untilMatch)
	fmt.Println("Second exercice:", countOccupiedSeats())

	elapsed := time.Since(start)
	fmt.Println("exec. time:", elapsed)
}

func setSeatMap(inputFilePath string) {
	seatMap = nil
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

func placePassengers(countFn func(seat) (int, int)) {
	changes := 1
	rounds := 0
	for changes > 0 {
		changes = addNewPassengersWave(countFn)
		rounds++
	}
}

func addNewPassengersWave(countFn func(seat) (int, int)) (stateChanges int) {
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
			newState, change := newSeatState(seat, countFn)
			if change {
				stateChanges++
				newSeatMap[r][c] = newState
			}
		}
	}
	seatMap = newSeatMap
	return
}

func newSeatState(s seat, countFn func(seat) (int, int)) (newState rune, hasSwitched bool) {
	nbAdjacentOccupiedSeats, switchSize := countFn(s)
	currentSeatState := seatMap[s.y][s.x]
	if currentSeatState == freeSeat && nbAdjacentOccupiedSeats == 0 {
		return occSeat, true
	}
	if currentSeatState == occSeat && nbAdjacentOccupiedSeats >= switchSize {
		return freeSeat, true
	}
	return currentSeatState, false
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
