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

type passport = []passportField

var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	rawPassports := extractRawPassports("input")
	completePassports := getPassportsWithAllFields(rawPassports)
	completePassportsWithValidation := getPassportsWithValidation(completePassports)
	fmt.Println("First exercise:", len(completePassports))
	fmt.Println("Second exercise:", len(completePassportsWithValidation))
}

func extractRawPassports(inputFilePath string) (rawPassports []passport) {
	f, err := os.Open(inputFilePath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(f)
	passportSplitRegex := regexp.MustCompile(`(?m)([a-z]{3}):([a-zA-Z0-9#]+)`)

	var rawPassport passport
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
		pFields := passportSplitRegex.FindAllStringSubmatch(line, -1)
		for _, pField := range pFields {
			pKey := pField[1]
			pValue := pField[2]
			rawPassport = append(rawPassport, passportField{key: pKey, value: pValue})
		}
	}
	return
}

func getPassportsWithAllFields(rawPassports []passport) (validPassports []passport) {
	for _, rawPassport := range rawPassports {
		if ContainsMandatoryFields(rawPassport) {
			validPassports = append(validPassports, rawPassport)
		}
	}
	return
}

func getPassportsWithValidation(rawPassports []passport) (validPassports []passport) {
	for _, rawPassport := range rawPassports {
		if ValidatePassport(rawPassport) {
			validPassports = append(validPassports, rawPassport)
		}
	}
	return
}
