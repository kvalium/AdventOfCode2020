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

func main() {
	for i, d := range getDirections("./directions") {
		if d.action == 'R' || d.action == 'L' {
			fmt.Printf("%v => %+v\n", i, d)
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
