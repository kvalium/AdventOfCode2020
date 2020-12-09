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

var bags = make(map[string]*bag)
var parentBagColors []string

// var bagCapacity = 1

var bagSplitRuleRegex = regexp.MustCompile(`([0-9])\s([a-z ]+)\sbag`)

func main() {
	getBags("./rules")
	getParentBagColors(bags["shiny gold"])
	fmt.Println("First exercise:", len(parentBagColors))
	size := getBagSize(bags["shiny gold"])
	fmt.Println("Second exercise:", size-1)
}

func getParentBagColors(bag *bag) {
	if len(bag.parents) == 0 {
		return
	}
	for _, parent := range bag.parents {
		if !inSlice(parentBagColors, parent.color) {
			parentBagColors = append(parentBagColors, parent.color)
		}
		getParentBagColors(bags[parent.color])
	}
	return
}

func getBagSize(bag *bag) int {
	cs := 1
	if len(bag.children) > 0 {
		for _, child := range bag.children {
			cs += child.count * getBagSize(bags[child.color])
		}
	}
	return cs
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

		if !isBagExists(color) {
			addNewBag(color)
		}
		for _, child := range children {
			addChildToBag(color, child)
		}
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
		if !isBagExists(color) {
			addNewBag(color)
		}
		addParentToBag(color, bagContent{color: bagColor, count: ct})
	}

	return
}

func isBagExists(color string) bool {
	_, exists := bags[color]
	return exists
}

func addNewBag(color string) {
	bags[color] = &bag{parents: nil, children: nil}
}

func addParentToBag(color string, parent bagContent) {
	b := bags[color]
	b.parents = append(bags[color].parents, parent)
	bags[color] = b
}

func addChildToBag(color string, child bagContent) {
	b := bags[color]
	b.children = append(bags[color].children, child)
	bags[color] = b
}

func inSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
