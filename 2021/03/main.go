package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const BIT_SIZE = 12

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	values := make([]string, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		value := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, value)
	}

	fmt.Println(part1(values))
	fmt.Println(part2(values))
}

func part1(input []string) int {
	var oneCount, zeroCount [BIT_SIZE]int
	for _, v := range input {
		value := strings.Split(v, "")
		for i, digit := range value {
			if digit == "1" {
				oneCount[i]++
			} else {
				zeroCount[i]++
			}
		}
	}

	var gammaBin, epsBin strings.Builder
	for i := 0; i < BIT_SIZE; i++ {
		if zeroCount[i] > oneCount[i] {
			gammaBin.WriteString("0")
			epsBin.WriteString("1")
		} else {
			gammaBin.WriteString("1")
			epsBin.WriteString("0")
		}
	}

	gamma := binToDec(gammaBin.String())
	epsilon := binToDec(epsBin.String())

	return gamma * epsilon
}

func part2(values []string) int {
	return recurse(values, 0, "O2") * recurse(values, 0, "CO2")
}

func recurse(input []string, level int, rating string) int {
	if len(input) == 1 {
		return binToDec(input[0])
	}

	var oneCount, zeroCount int
	var oneArr, zeroArr []string
	var resArr []string

	for i, v := range input {
		digit := strings.Split(v, "")[level]
		if digit == "1" {
			oneCount++
			oneArr = append(oneArr, input[i])
		} else {
			zeroCount++
			zeroArr = append(zeroArr, input[i])
		}
	}

	if ratingComparison(zeroCount, oneCount, rating) {
		resArr = oneArr
	} else {
		resArr = zeroArr
	}

	return recurse(resArr, level+1, rating)
}

func ratingComparison(zeroBits, oneBits int, rating string) bool {
	switch rating {
	case "O2":
		return oneBits >= zeroBits
	case "CO2":
		return oneBits < zeroBits
	}

	return false
}

func binToDec(bin string) int {
	dec, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic("failed to convert bin to dec")
	}

	return int(dec)
}
