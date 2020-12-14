package main

import "fmt"

var coordinates = []rune{
	'N',
	'E',
	'S',
	'W',
}

var actionEffect map[rune]func(position, direction) position

func initActions() {
	actionEffect = map[rune]func(position, direction) position{
		'N': moveNorth,
		'E': moveEast,
		'S': moveSouth,
		'W': moveWest,
		'F': moveForward,
		'L': turnLeft,
		'R': turnRight,
	}
}

func moveNorth(p position, d direction) position {
	return position{x: p.x, y: p.y + d.value, dir: p.dir}
}

func moveSouth(p position, d direction) position {
	return position{x: p.x, y: p.y - d.value, dir: p.dir}
}

func moveEast(p position, d direction) position {
	return position{x: p.x + d.value, y: p.y, dir: p.dir}
}

func moveWest(p position, d direction) position {
	return position{x: p.x - d.value, y: p.y, dir: p.dir}
}

func moveForward(p position, d direction) position {
	return actionEffect[p.dir](p, d)
}

func turnRight(p position, d direction) position {
	i := d.value / 90
	nDir := newDirection(p.dir, i)
	return position{x: p.x, y: p.y, dir: nDir}
}

func turnLeft(p position, d direction) position {
	i := d.value / 90
	nDir := newDirection(p.dir, -i)
	return position{x: p.x, y: p.y, dir: nDir}
}

func newDirection(d rune, n int) rune {
	var currIdx int
	coordsLen := len(coordinates)
	for curr := 0; curr < coordsLen; curr++ {
		if coordinates[curr] == d {
			currIdx = curr
		}
	}
	newIdx := currIdx + n
	fmt.Println(string(d), n, newIdx)

	if newIdx >= coordsLen {
		newIdx = coordsLen - newIdx
	}
	if newIdx < 0 {
		newIdx = coordsLen + newIdx
	}
	return coordinates[newIdx]
}
