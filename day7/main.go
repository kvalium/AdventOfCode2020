package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bagContent struct {
	color string
	count int
}

type bag struct {
	parents  []bagContent
	children []bagContent
}

var bags map[string]bag

var bagSplitRuleRegex = regexp.MustCompile(`([0-9])\s([a-z ]+)\sbag`)

func main() {
	getBags("./rules")
}

func getBags(inputFilePath string) {
	f, err := os.Open(inputFilePath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(f)

	for {
		bagRuleLine, err := rd.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		color, children := getBagRule(bagRuleLine)
		fmt.Println(color, children)
	}
	return
}

func getBagRule(bagRuleLine string) (bagColor string, children []bagContent) {
	result := strings.Split(bagRuleLine, " bags contain ")
	bagColor = result[0]

	r := bagSplitRuleRegex.FindAllSubmatch([]byte(result[1]), -1)

	for _, bc := range r {
		ct, _ := strconv.Atoi(string(bc[1]))
		color := string(bc[2])
		children = append(children, bagContent{color: color, count: ct})
	}

	return
}
