package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read in the test input
	// Parse it correctly
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Need to split the two columns and store them in a ds; slice
	var lList []int
	var rList []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Split the lines by a space
		sp := strings.Fields(scanner.Text())
		fmt.Printf("%+v\n", sp)
		// Separate each slice based on length, and store the left and right
		// values into their respective slices
		if len(sp) >= 2 {
			firstPart, errFp := strconv.Atoi(sp[0])
			secondPart, errSp := strconv.Atoi(sp[1])
			if errFp != nil {
				log.Fatal(errFp)
			}
			if errSp != nil {
				log.Fatal(errSp)
			}

			lList = append(lList, firstPart)
			rList = append(rList, secondPart)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%#v\n", partOne(lList, rList))
	fmt.Println(partTwo(lList, rList))
}

// partOne solves part one of day 1.
// This function needs to sort both lists, calculate their 1-D euclidean distance,
// and then sum up those distances and return a total distance.
func partOne(leftList, rightList []int) int {
	sort.Slice(leftList, func(i, j int) bool {
		return leftList[i] < leftList[j]
	})
	sort.Slice(rightList, func(i, j int) bool {
		return rightList[i] < rightList[j]
	})

	var distance int

	for i := range leftList {
		diff := math.Abs(float64(leftList[i]) - float64(rightList[i]))
		distance += int(diff)
	}
	return distance
}

// Compute the frequencies of the values from the left list occurring in the right list,
// and then compute the similarity score by computing the value of the left list * its frequency,
// and sum up each of these individual scores.
func partTwo(leftList, rightList []int) int {
	// Using a map as a set to store right list values
	seenMap := make(map[int]bool)
	for _, v := range rightList {
		seenMap[v] = true
	}
	// k -> left side value
	// v -> freq of k
	// Iterate over the left list, and compute frequencies for each value that
	// is seen in the right list
	freqMap := make(map[int]int)
	for _, v := range leftList {
		if seenMap[v] {
			freqMap[v]++
		}
	}

	// Compute the similarity score
	var simScore int
	for _, v := range rightList {
		individualScore := freqMap[v] * v
		simScore += individualScore
	}
	return simScore
}
