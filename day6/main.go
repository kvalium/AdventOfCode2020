package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type groupAnswers []answers
type answers []rune

func main() {
	totalPassengerAnswers := getPassengersAnswers("./answers")
	fmt.Println("First exercise:", sumGroupAnswered(totalPassengerAnswers))
	fmt.Println("Second exercise:", sumSameAnswered(totalPassengerAnswers))
}

func getPassengersAnswers(inputFilePath string) (totalAnswers []groupAnswers) {
	f, err := os.Open(inputFilePath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(f)

	var groupAnswers groupAnswers
	for {
		line, err := rd.ReadString('\n')

		if err == io.EOF {
			totalAnswers = append(totalAnswers, groupAnswers)
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if len(line) == 1 {
			totalAnswers = append(totalAnswers, groupAnswers)
			groupAnswers = nil
			continue
		}

		groupAnswers = append(groupAnswers, []rune(line))
	}
	return
}

func sumGroupAnswered(totalAnswers []groupAnswers) (sum int) {
	for _, groupAnswers := range totalAnswers {
		var uniqueAnswers answers
		for _, passengerAnswers := range groupAnswers {
			for _, letter := range passengerAnswers {
				if letter == '\n' || runeInSlice(letter, uniqueAnswers) {
					continue
				}
				uniqueAnswers = append(uniqueAnswers, letter)
				sum++
			}
		}
	}
	return
}

func sumSameAnswered(totalAnswers []groupAnswers) (sum int) {

	for _, gAnswers := range totalAnswers {
		var commonGAnswers answers
		firstPassengersAnswers := gAnswers[0]
		for _, letterToTest := range firstPassengersAnswers {
			if letterToTest == '\n' {
				continue
			}
			occ := 0
			for _, passengerAnswer := range gAnswers {
				if runeInSlice(letterToTest, passengerAnswer) {
					occ++
				}
			}
			if occ == len(gAnswers) {
				commonGAnswers = append(commonGAnswers, letterToTest)
			}
		}
		sum += len(commonGAnswers)
	}
	return
}

func runeInSlice(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
