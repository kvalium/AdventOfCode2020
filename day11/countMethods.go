package main

func adjacentOnly(s seat) (ct, switchSize int) {
	switchSize = 4
	for _, p := range positions {
		if isOccupied(seat{x: s.x + p.x, y: s.y + p.y}) {
			ct++
		}
	}
	return
}

func untilMatch(s seat) (ct, switchSize int) {
	switchSize = 5
	for _, p := range positions {
		if checkForOccupiedSeat(&s, p) {
			ct++
		}
	}
	return
}

func checkForOccupiedSeat(s *seat, dir seat) bool {
	checkSeat := seat{x: s.x + dir.x, y: s.y + dir.y}
	if !isValid(checkSeat) {
		return false
	}
	if isFloor(checkSeat) {
		return checkForOccupiedSeat(&checkSeat, dir)
	}
	return isOccupied(checkSeat)
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
	return s.x >= 0 && s.y >= 0 && s.y <= nbRows-1 && s.x <= nbCols-1
}
