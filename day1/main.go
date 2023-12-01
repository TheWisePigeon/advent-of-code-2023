package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getCalibrationValue() (int, error) {
	calibrationValue := 0
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	arr := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if intValue, err := strconv.Atoi(string(line[i])); err != nil {
				continue
			} else {
				arr = append(arr, intValue)
			}
		}
		lineValue, err := strconv.Atoi(fmt.Sprintf("%d%d", arr[0], arr[len(arr)-1]))
		if err != nil {
			return 0, err
		}
		calibrationValue += lineValue
    arr = []int{}

	}
  println(calibrationValue)
	return calibrationValue, nil
}

func main() {
	getCalibrationValue()
}
