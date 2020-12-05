package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

type passportField struct {
	key, value string
}

var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	rawPassports := extractRawPassports("input")
	completePassports := countValidPassports(rawPassports)
	fmt.Println("First exercise:", completePassports)
}

func extractRawPassports(inputFilePath string) (rawPassports [][]passportField) {
	f, err := os.Open(inputFilePath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(f)
	re := regexp.MustCompile(`(?m)([a-z]{3}):([a-zA-Z0-9#]+)`)

	var rawPassport []passportField
	for {
		line, err := rd.ReadString('\n')

		if err == io.EOF {
			rawPassports = append(rawPassports, rawPassport)
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if len(line) == 1 {
			rawPassports = append(rawPassports, rawPassport)
			rawPassport = nil
			continue
		}
		pFields := re.FindAllStringSubmatch(line, -1)
		for _, pField := range pFields {
			pKey := pField[1]
			pValue := pField[2]
			rawPassport = append(rawPassport, passportField{key: pKey, value: pValue})
		}
	}
	return
}

func countValidPassports(rawPassports [][]passportField) (count int) {
	for _, rawPassport := range rawPassports {
		if len(rawPassport) < len(mandatoryFields) {
			continue
		}
		if containsMandatoryFields(rawPassport) {
			count++
		}
	}
	return
}

func containsMandatoryFields(x []passportField) bool {
	var keys []string
	for _, pField := range x {
		if pField.key == "cid" {
			continue
		}
		keys = append(keys, pField.key)
	}
	return sameStringSlice(keys, mandatoryFields)
}

func sameStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		diff[_x]++
	}
	for _, _y := range y {
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y]--
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}
