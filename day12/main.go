package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type direction struct {
	action rune
	value  int
}

type position struct {
	x, y int
	dir  rune
}

func main() {
	initActions()
	directions := getDirections("./directions")
	manhattan := manhattanTravel(directions)
	fmt.Println(manhattan)
}

func manhattanTravel(directions []direction) (manhattan int) {
	currentPosition := position{x: 0, y: 0, dir: 'E'}
	maxX, maxY := 0, 0
	for _, d := range directions {
		newPosition := actionEffect[d.action](currentPosition, d)
		fmt.Printf("%v%v => %+v (pos %v)\n", string(d.action), d.value, newPosition, string(newPosition.dir))

		if absInt(newPosition.x) > maxX {
			maxX = absInt(newPosition.x)
		}
		if absInt(newPosition.y) > maxY {
			maxY = absInt(newPosition.y)
		}
		currentPosition = newPosition
	}
	fmt.Println(maxX, maxY)
	return maxX + maxY
}

func getDirections(inputFilePath string) (directions []direction) {
	splitDirRegex := regexp.MustCompile(`(F|R|E|W|L|N|S)([0-9]+)`)
	f, err := os.Open(inputFilePath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(f)
	line := 0
	for {
		instructionLine, err := rd.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		dir := splitDirRegex.FindStringSubmatch(instructionLine)
		v, _ := strconv.Atoi(dir[2])
		a := []rune(dir[1])[0]
		directions = append(directions, direction{action: a, value: v})
		line++
	}
	return
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
