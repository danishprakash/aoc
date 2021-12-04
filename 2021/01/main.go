package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	values := make([]int, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
	}

	fmt.Println(day01(values))
	fmt.Println(day02(values))
}

func day01(nums []int) int {
	var count int
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		}
	}
	return (count + 1)
}

func day02(nums []int) int {
	var count, prevSum int
	for i := 0; i < len(nums); i++ {
		if i > 1 {
			currSum := nums[i] + nums[i-1] + nums[i-2]
			if prevSum > 0 && currSum > prevSum {
				count++
			}
			prevSum = currSum
		}
	}
	return count
}
