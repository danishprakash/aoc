package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cmds struct {
	dir   string
	value int
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	input := make([]cmds, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		tmp := scanner.Text()
		text := strings.Split(tmp, " ")

		dir := text[0]
		value, err := strconv.Atoi(text[1])
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, cmds{
			dir,
			value,
		})
	}

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []cmds) int {
	var fwd, depth int
	for _, cmd := range input {
		if cmd.dir == "forward" {
			fwd += cmd.value
		} else if cmd.dir == "up" {
			depth -= cmd.value
		} else {
			depth += cmd.value
		}
	}
	return fwd * depth
}

func part2(input []cmds) int {
	var fwd, depth, aim int
	for _, cmd := range input {
		if cmd.dir == "forward" {
			fwd += cmd.value
			depth += aim * cmd.value
		} else if cmd.dir == "up" {
			aim -= cmd.value
		} else {
			aim += cmd.value
		}
	}
	return fwd * depth
}
