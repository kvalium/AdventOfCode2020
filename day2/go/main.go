package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type policy struct {
	a      int
	b      int
	letter string
}

func main() {
	inRange := inRangePolicy()
	byPosition := byPositionPolicy()
	fmt.Println("First exercise:", inRange)
	fmt.Println("Second exercise:", byPosition)
}

func byPositionPolicy() (count int) {
	for _, p := range Passwords {
		policy, password, e := getPasswordAndPolicy(p)
		if e != nil {
			continue
		}
		passwordLetters := []rune(password)
		letter := []rune(policy.letter)[0]
		inFirst := passwordLetters[policy.a-1] == letter
		inSecond := passwordLetters[policy.b-1] == letter
		if !inFirst == inSecond {
			count++
		}
	}
	return
}

func inRangePolicy() (count int) {
	for _, p := range Passwords {
		policy, password, e := getPasswordAndPolicy(p)
		if e != nil {
			continue
		}
		letterOcc := strings.Count(password, policy.letter)
		isValid := letterOcc >= policy.a && letterOcc <= policy.b
		if isValid {
			count++
		}
	}
	return
}

func getPasswordAndPolicy(s string) (*policy, string, error) {
	var policyExtractRegex = regexp.MustCompile(`(\d+)-(\d+)\s(\w):\s(\w+)`)
	m := policyExtractRegex.FindSubmatch([]byte(s))
	if len(m) < 5 {
		return nil, "", errors.New("invalid policy format")
	}
	a, err1 := strconv.Atoi(string(m[1]))
	if err1 != nil {
		return nil, "", err1
	}
	b, err2 := strconv.Atoi(string(m[2]))
	if err2 != nil {
		return nil, "", err2
	}
	letter := string(m[3])
	policy := policy{
		a,
		b,
		letter,
	}
	password := string(m[4])
	return &policy, password, nil
}
