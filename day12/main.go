package main

import (
	"bufio"
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

type position struct{ x, y int }

var actionEffect = map[rune]position{
	'N': {x: 0, y: 1},
	'E': {x: 1, y: 0},
	'S': {x: 0, y: -1},
	'W': {x: 1, y: 0},
}

func main() {
	directions := getDirections("./directions")
	getFinalDestination(directions)
}

func getFinalDestination(directions []direction) {
	currentDir := 'E'
	currentPosition := position{x: 0, y: 0}
	for _, d := range directions {
		if d.action == 'F' {
			currentPosition = actionEffect[currentPosition]
		}
	}
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
