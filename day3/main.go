package main

import (
	"fmt"
)

type position struct {
	x int
	y int
}

var tree = rune('#')
var slopePatterns = []position{
	{x: 1, y: 1},
	{x: 3, y: 1},
	{x: 5, y: 1},
	{x: 7, y: 1},
	{x: 1, y: 2},
}

func main() {
	fmt.Println("First exercise: ", countEncounteredTrees(slopePatterns[1]), "trees - slope", slopePatterns[1])

	fmt.Println("\nSecond exercise:")
	result := 1
	for _, p := range slopePatterns {
		slopeTrees := countEncounteredTrees(p)
		fmt.Println(slopeTrees, "trees - slope", p)
		result *= slopeTrees
	}
	fmt.Println("Result:", result)
}

func countEncounteredTrees(slopePattern position) (nbTrees int) {
	currentPosition := position{x: 1, y: 1}
	nbRows := len(Treemap)
	lineSize := len(Treemap[0])
	for currentPosition.y <= nbRows {
		if isTreeAt(currentPosition) {
			nbTrees++
		}
		currentPosition = nextMove(currentPosition, lineSize, slopePattern)
	}
	return
}

func nextMove(p position, lineSize int, slopePattern position) position {
	nextX := p.x+slopePattern.x
	if(nextX > lineSize) {
		nextX = nextX - lineSize
	}
	return position{
		x:nextX,
		y: p.y+slopePattern.y,
	}
}

func isTreeAt(pos position) bool {
	row := []rune(Treemap[pos.y - 1]);
	return row[pos.x - 1] == tree
}