package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// NOTES: r -> report, c -> level

func main() {
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// read each line by line, split by whitespace and append to our data slice
		line := strings.Fields(scanner.Text())
		formatted := strings.Join(line, " ")
		data = append(data, formatted)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Holds parsed output in format specified
	var safeReports int
	for i := range data {
		safeReports = DayTwoPartOne(data[i], safeReports)
	}

	fmt.Println("Safe Reports:", safeReports)
}

/*
	*
  * So, a report only counts as safe if both of the following are true:
  *
	* The levels are either all increasing or all decreasing.
	* Any two adjacent levels differ by at least one and at most three.
*/
func DayTwoPartOne(input string, safeReports int) int {
	nums := parseString(input)
	// Closure that computes the difference from two values and returns true if in the range [1, 3]
	diff := func(prev, curr int) bool {
		d := math.Abs(float64(prev) - float64(curr))
		return d >= 1 && d <= 3
	}

	isIncr := true
	isDecr := true

	for i := 1; i < len(nums); i++ {
		prev := nums[i-1]
		curr := nums[i]
		if !diff(prev, curr) {
			return safeReports
		}

		if curr <= prev {
			isIncr = false
		}
		if curr > prev {
			isDecr = false
		}
	}

	// Must be strictly incr or decr
	if !isIncr && !isDecr {
		return safeReports
	}
	return 1 + safeReports
}

// Parse a single string into a int slice.
func parseString(s string) []int {
	fields := strings.Fields(s)
	var nums []int

	for _, field := range fields {
		v, err := strconv.Atoi(field)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, v)
	}
	return nums
}
