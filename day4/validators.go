package main

import (
	"regexp"
	"strconv"
)

var heightRegex = regexp.MustCompile(`([0-9]+)(cm|in)`)
var hairColorRegex = regexp.MustCompile(`\#[0-9a-f]{6}`)
var pidRegex = regexp.MustCompile(`^\d{9}$`)
var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

var validatePassportField = map[string]func(s string) bool{
	"byr": isValidBirthYear,
	"iyr": isValidIssueYear,
	"eyr": isValidExpirationYear,
	"hgt": isValidHeight,
	"hcl": isValidHairColor,
	"ecl": isValidEyeColor,
	"pid": isValidPid,
	"cid": isValidCid,
}

// ContainsMandatoryFields checks that all mandatory passport fields are present
func ContainsMandatoryFields(p passport) bool {
	if len(p) < len(mandatoryFields) {
		return false
	}

	var keys []string
	for _, f := range p {
		if f.key == "cid" {
			continue
		}
		keys = append(keys, f.key)
	}
	return sameStringSlice(keys, mandatoryFields)
}

// ValidatePassport checks all passport properties are OK
func ValidatePassport(p passport) bool {
	if !ContainsMandatoryFields(p) {
		return false
	}
	for _, f := range p {
		isValid := validatePassportField[f.key](f.value)
		if !isValid {
			return false
		}
	}
	return true
}

func isValidBirthYear(s string) bool {
	return isBeetween(s, 1920, 2002)
}

func isValidIssueYear(s string) bool {
	return isBeetween(s, 2010, 2020)
}

func isValidExpirationYear(s string) bool {
	return isBeetween(s, 2020, 2030)
}

func isValidHeight(heightUnit string) bool {
	m := heightRegex.FindStringSubmatch(heightUnit)
	if len(m) != 3 {
		return false
	}
	height := m[1]
	unit := m[2]

	if unit == "cm" {
		return isBeetween(height, 150, 193)
	}
	return isBeetween(height, 59, 76)
}

func isValidHairColor(e string) bool {
	return hairColorRegex.MatchString(e)
}

func isValidEyeColor(e string) bool {
	_, isValid := inSlice(eyeColors, e)
	return isValid
}

func isValidPid(p string) bool {
	return pidRegex.MatchString(p)
}

func isValidCid(p string) bool {
	return true
}

func isBeetween(s string, start, end int) bool {
	i, _ := strconv.Atoi(s)
	return i >= start && i <= end
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

func inSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
