package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var lettersToInt = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getNextWord(line, word string, i int) string {
	return line[i : i+len(word)]
}

func notEnoughLengthLeftForWord(line, word string, i int) bool {
	return i+(len(word)-1) > len(line)-1
}

func pseudoTokenizer(char, line string, i int) (int, bool) {
	switch char {
	case "o":
		if notEnoughLengthLeftForWord(line, "one", i) {
			return 0, false
		}
		if getNextWord(line, "one", i) == "one" {
			return lettersToInt["one"], true
		}
		return 0, false
	case "t":
		if notEnoughLengthLeftForWord(line, "two", i) {
			return 0, false
		}
		if getNextWord(line, "two", i) == "two" {
			return lettersToInt["two"], true
		}
		if notEnoughLengthLeftForWord(line, "three", i) {
			return 0, false
		}
		if getNextWord(line, "three", i) == "three" {
			return lettersToInt["three"], true
		}
		return 0, false
	case "f":
		if notEnoughLengthLeftForWord(line, "four", i) {
			return 0, false
		}
		if getNextWord(line, "four", i) == "four" {
			return lettersToInt["four"], true
		}
		if getNextWord(line, "five", i) == "five" {
			return lettersToInt["five"], true
		}
		return 0, false
	case "s":
		if notEnoughLengthLeftForWord(line, "six", i) {
			return 0, false
		}
		if getNextWord(line, "six", i) == "six" {
			return lettersToInt["six"], true
		}
		if notEnoughLengthLeftForWord(line, "seven", i) {
			return 0, false
		}
		if getNextWord(line, "seven", i) == "seven" {
			return lettersToInt["seven"], true
		}
		return 0, false
	case "e":
		if notEnoughLengthLeftForWord(line, "eight", i) {
			return 0, false
		}
		if getNextWord(line, "eight", i) == "eight" {
			return lettersToInt["eight"], true
		}
		return 0, false
	case "n":
		if notEnoughLengthLeftForWord(line, "nine", i) {
			return 0, false
		}
		if getNextWord(line, "nine", i) == "nine" {
			return lettersToInt["nine"], true
		}
		return 0, false
	default:
		return 0, false
	}
}

func getCalibrationValue() (int, error) {
	calibrationValue := 0
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		arr := []int{}
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if intValue, err := strconv.Atoi(char); err != nil {
				value, ok := pseudoTokenizer(char, line, i)
				if !ok {
					continue
				}
				arr = append(arr, value)
			} else {
				arr = append(arr, intValue)
			}
		}
		lineValue, err := strconv.Atoi(fmt.Sprintf("%d%d", arr[0], arr[len(arr)-1]))
		if err != nil {
			return 0, err
		}
		calibrationValue += lineValue
	}
	return calibrationValue, nil
}

func main() {
}
