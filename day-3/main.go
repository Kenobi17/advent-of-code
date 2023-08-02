package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// over-engineered the shit out of this day so i can play a little bit with structs and methods

const input string = "./big-input.txt"

type Compartment string

type Inventory string

type Badge string

type Rucksack struct {
	left  Compartment
	right Compartment
	items Inventory
}

type Group struct {
	badge Badge
	a     Rucksack
	b     Rucksack
	c     Rucksack
}

func (r Rucksack) GetRepeatedItem() string {
	for _, c := range r.left {
		item := string(c)
		if r.right.ContainsItem(item) {
			return item
		}
	}
	return ""
}

func getCommonItem(a, b Rucksack) (string, bool) {
	for _, i := range a.items {
		item := string(i)
		if a.items.ContainsItem(item) {
			return item, true
		}
	}
	return "", false
}

func (c Compartment) ContainsItem(i string) bool {
	return strings.Index(string(c), i) >= 0
}

func (I Inventory) ContainsItem(i string) bool {
	return strings.Index(string(I), i) >= 0
}

func generateBadge(g Group) Badge {
	for _, r := range g.a.items {
		item := r
		if g.b.items.ContainsItem(string(item)) && g.c.items.ContainsItem(string((item))) {
			return Badge(item)
		}
	}

	return Badge("")
}

var prioValues = map[string]int{}

func initPrioValues() {
	for i := 'a'; i <= 'z'; i++ {
		prioValues[string(i)] = int(i - 'a' + 1)
	}
	for i := 'A'; i <= 'Z'; i++ {
		prioValues[string(i)] = int(i - 'A' + 27)
	}
}

func generateGroups(rs []Rucksack) []Group {
	gs := []Group{}

	for i := range rs {
		if i%3 == 0 {
			g := Group{a: rs[i], b: rs[i+1], c: rs[i+2]}
			g.badge = generateBadge(g)
			gs = append(gs, g)
		}
	}

	return gs
}

func getRucksacksData() []Rucksack {
	rs := []Rucksack{}

	buffer, _ := os.Open(input)

	s := bufio.NewScanner(buffer)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		items := Inventory(s.Text())

		mid := len(items) / 2

		left, right := items[:mid], items[mid:]

		r := Rucksack{Compartment(left), Compartment(right), items}

		rs = append(rs, r)
	}

	return rs
}

func calculatePrioritySum(items []string) (prioSum int) {
	for _, i := range items {
		prioSum = prioSum + prioValues[i]
	}
	return
}

func getRucksacksRepeated(rs []Rucksack) (repeated []string) {
	for _, r := range rs {
		repeated = append(repeated, r.GetRepeatedItem())
	}
	return
}

func getGroupsBadges(gs []Group) (bs []string) {
	for _, g := range gs {
		bs = append(bs, string(g.badge))
	}
	return
}

func main() {
	initPrioValues()

	rs := getRucksacksData()
	gs := generateGroups(rs)

	repeated := getRucksacksRepeated(rs)
	badges := getGroupsBadges(gs)

	prioSum := calculatePrioritySum(repeated)
	prioSumBadges := calculatePrioritySum(badges)

	fmt.Println(prioSum)
	fmt.Println(prioSumBadges)
}
