package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func readInput() [][]int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var records [][]int

	for scanner.Scan() {
		line := scanner.Text()
		levels := strings.Fields(line)

		var record []int
		for _, level := range levels {
			value, _ := strconv.Atoi(level)
			record = append(record, value)
		}

		records = append(records, record)
	}

	return records
}

func isSafe(record []int) bool {

	direction := 1
	if record[1] < record[0] {
		direction = -1
	}

	for i := 0; i < len(record)-1; i++ {
		diff := direction * (record[i+1] - record[i])
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true

}

func part1() int {
	records := readInput()
	safeReports := 0

	for _, record := range records {
		if isSafe(record) {
			safeReports++
		}
	}

	return safeReports
}

func part2() int {
	reports := readInput()
	safeReports := 0

	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		} else {
			for j := range report {
				newReport := append([]int{}, report[:j]...)
				newReport = append(newReport, report[j+1:]...)
				if isSafe(newReport) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}
