package main

import (
	"bufio"
	"fmt"
	// "fmt"
	"os"
	"strconv"
	"strings"
)

var colorToMaxValues = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func setIsValid(set string) bool {
	setData := strings.Split(set, ",")
	for _, dataStr := range setData {
		data := strings.Split(strings.TrimSpace(dataStr), " ")
		amountStr, color := data[0], data[1]
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			panic(err)
		}
		if colorToMaxValues[color] < amount {
			return false
		}
	}
	return true
}

func partOne() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	gameId := 1
	validGames := []int{}
	for s.Scan() {
		line := s.Text()
		sets := strings.Split(strings.Split(line, ":")[1], ";")
		setNum := len(sets)
		validSets := 0
		for _, set := range sets {
			if !setIsValid(set) {
				break
			}
			validSets += 1
		}
		if validSets == setNum {
			validGames = append(validGames, gameId)
		}
		gameId += 1
	}
	sum := 0
	for _, gameId := range validGames {
		sum += gameId
	}
	fmt.Printf("Solution is %d", sum)
}

func partTwo() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	powerSum := 0
	for s.Scan() {
		line := s.Text()
		minimumRequiredCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		_ = minimumRequiredCubes
		sets := strings.Split(strings.Split(line, ":")[1], ";")
		for _, set := range sets {
			setData := strings.Split(set, ",")
			for _, dataStr := range setData {
				data := strings.Split(strings.TrimSpace(dataStr), " ")
				amountStr, color := data[0], data[1]
				amount, err := strconv.Atoi(amountStr)
				if err != nil {
					panic(err)
				}
				for range minimumRequiredCubes {
					if minimumRequiredCubes[color] < amount {
						minimumRequiredCubes[color] = amount
					}
				}
			}
		}
		setsPower := 1
		for _, v := range minimumRequiredCubes {
			setsPower *= v
		}
		powerSum += setsPower

	}
	fmt.Println("Part 2 solution is ", powerSum)
}

func main() {
	partTwo()
}
