package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func readInput() ([1000]int, [1000]int) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1 [1000]int
	var list2 [1000]int

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		b := strings.Fields(line)
		list1[i], _ = strconv.Atoi(b[0])
		list2[i], _ = strconv.Atoi(b[1])
		i++
	}

	return list1, list2
}

func part1() int {
	list1, list2 := readInput()

	sort.Ints(list1[:])
	sort.Ints(list2[:])
	diff := 0

	for i := 0; i < len(list1); i++ {
		a := list1[i]
		b := list2[i]
		c := a - b
		if c < 0 {
			c *= -1
		}
		diff += c
	}

	return diff
}

func part2() int {
	list1, list2 := readInput()
	similarity := 0

	for i := 0; i < len(list1); i++ {
		count := 0
		a := list1[i]
		for i := 0; i < len(list2); i++ {
			if a == list2[i] {
				count++
			}
		}
		similarity += a * count
	}

	return similarity
}
