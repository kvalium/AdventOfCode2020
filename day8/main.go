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

type instruction struct {
	cmd  string
	val  int
	line int
}

func main() {
	instructions := getInstructions("./boot")
	a, infL := run(instructions)
	fmt.Println("First exercise:", a, "(inf. loop starts at line #", infL, ")")
	b := changeIns(instructions)
	fmt.Println("Second exercise:", b)
}

func changeIns(instructions []instruction) int {
	for i, ins := range instructions {
		if ins.cmd == "acc" {
			continue
		}
		newIns := "jmp"
		if ins.cmd == "jmp" {
			newIns = "nop"
		}
		instCopy := append([]instruction(nil), instructions...)
		b := instCopy[i]
		b.cmd = newIns
		instCopy[i] = b
		acc, inf := run(instCopy)
		if inf == -1 {
			return acc
		}
	}
	return -1
}

func run(instructions []instruction) (acc int, infiniteLoopLine int) {
	var passedInstructionLines []int
	currLine := 0
	for currLine <= len(instructions)-1 {
		ins := instructions[currLine]
		currLine = ins.line
		if inSlice(passedInstructionLines, currLine) {
			return acc, currLine
		}
		passedInstructionLines = append(passedInstructionLines, currLine)
		switch ins.cmd {
		case "jmp":
			currLine += ins.val
			continue
		case "acc":
			acc += ins.val
		}
		currLine++
	}
	return acc, -1
}

func getInstructions(inputFilePath string) (instructions []instruction) {
	splitInstRegex := regexp.MustCompile(`(nop|jmp|acc)\s(.*)`)
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

		ins := splitInstRegex.FindStringSubmatch(instructionLine)
		v, _ := strconv.Atoi(ins[2])
		instructions = append(instructions, instruction{cmd: ins[1], val: v, line: line})
		line++
	}
	return
}

func inSlice(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
