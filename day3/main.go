package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// mul ( 2, 4) -> invalid
// mul(2,4) -> valid
// scanning bytes of data
//  looking for phrase "mul" IMMEDIATELY, if ANYTHING else, (even whitespace) is next, invalid
//    -> see mul, look for "("
//      -> "(" found, i am only expecting a number next.
//          Anything else and I mark the n next tokens as invalid until i see "mul" again
//      -> "mul(2" -> look for ",", found ",", look for number, found "4", look for ")", found ")"
//          mark this as valid, and store the product of these two numbers in a stack
//          note: dont need stack, just do operation in place
//  ex) xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))
//      XCCCCCCCCXXCCCXXXXXXXXXXXXXXXCCCCCCCCXCCCCCCCCCXCCCCXCCCCCCCCC
//      sum = 40 + 88 + 25 + 8

func main() {
	f, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// var charSlice []string
	var data []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// do stuff
		line := strings.Fields(scanner.Text())
		data = append(data, line...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(data)
	fmt.Println(DayThreePartOne(data))
}

func DayThreePartOne(data []string) int {
	var sum int
	for _, s := range data {
		for i := range s {
			// Implement logic to check for sequence mul(int,int)
			// i + 2 == 'mul'
			if i+2 < len(s) && string(s[i]) == "m" && string(s[i+1]) == "u" && string(s[i+2]) == "l" {
				if i+3 < len(s) && string(s[i+3]) == "(" {
					// advance our index "mul("
					i += 4
					n1, flag, advI := parseInteger(s, i)
					if !flag {
						continue
					}
					i = advI
					if i < len(s) && string(s[i]) == "," {
						// advance our index ","
						i++
						n2, flag, advI := parseInteger(s, i)
						if !flag {
							continue
						}
						i = advI

						if i < len(s) && string(s[i]) == ")" {
							sum += n1 * n2
							i++ // advance past ")"
						} else {
							i++ // advance past some invalid rune
						}
					}
				}
			}
		}
	}
	return sum
}

// parseInteger is a helper to parse a rune into a 32-bit integer, and it returns
// the integer itself (if successful), a bool indicating success, and the length we've traversed to represent the number
func parseInteger(s string, i int) (int, bool, int) {
	// Our goal is to use atoi on a slice of s s.t. s[start:i], where we increment i
	// per digit we find in our token sequence
	start := i
	if i < len(s) && unicode.IsDigit(rune(s[i])) {
		for i < len(s) && unicode.IsDigit(rune(s[i])) {
			i++
		}
		n, err := strconv.Atoi(s[start:i])
		if err != nil {
			return 0, false, i
		}
		return n, true, i
	}
	return 0, false, i
}
