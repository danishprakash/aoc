package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var spaceRegexp = regexp.MustCompile(`\s+`)
var globalWinCount = 0
var board = make([][][]int, 100)
var seen [100]int

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	var draws []int
	values := make([]string, 0)
	scanner := bufio.NewScanner(fp)
	var pos int
	boardCount := 0
	for k := 0; k < 100; k++ {
		board[k] = make([][]int, 5)
		for i := 0; i < 5; i++ {
			board[k][i] = make([]int, 5)
		}
	}

	for scanner.Scan() {
		text := scanner.Text()

		// parse the draws
		if pos == 0 {
			for _, v := range strings.Split(text, ",") {
				intV, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				draws = append(draws, intV)
			}
		} else if len(text) != 0 {
			for i := 0; i < 5; i++ {
				text := strings.TrimSpace(spaceRegexp.ReplaceAllString(scanner.Text(), " "))
				v := strings.Split(text, " ")
				for j := 0; j < 5; j++ {
					num, err := strconv.Atoi(v[j])
					if err != nil {
						log.Fatal(err)
					}
					board[boardCount][i][j] = num
					// scanner.Scan()
				}
				scanner.Scan()
			}
			boardCount++
		}

		values = append(values, text)
		pos++
	}

	fmt.Println(part1(draws))
	fmt.Println(part2(draws))
}

func part1(draws []int) int {
	return iterate(draws, "first")
}

func part2(draws []int) int {
	return iterate(draws, "second")
}

func iterate(draws []int, kind string) int {
	for _, d := range draws {
		for i := 0; i < 100; i++ {
			if seen[i] == 1 {
				continue
			}

			for j := 0; j < 5; j++ {
				for k := 0; k < 5; k++ {
					if board[i][j][k] == d {
						board[i][j][k] = -board[i][j][k]
						out := checkWin(draws, i, j, k, d)
						if out > 0 {
							if kind == "first" {
								return out
							}

							globalWinCount++
							if globalWinCount == 100 {
								return out
							}
							seen[i] = 1
						}
					}
				}
			}
		}
	}

	return 0
}

func visited(draws []int, item, d int) bool {
	if item < 0 {
		return true
	} else if item == 0 {
		for _, v := range draws {
			if v == item {
				return true
			}

			if v == d {
				return false
			}
		}
	}

	return false
}

func unvisited(draws []int, item, d int) bool {
	if item > 0 {
		return true
	} else if item == 0 {
		for _, v := range draws {
			if v == item {
				return true
			}

			if v == d {
				return false
			}
		}
	}

	return false
}

func getScore(draws []int, idx, d int) int {
	unmarkedSum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if unvisited(draws, board[idx][i][j], d) {
				unmarkedSum += board[idx][i][j]
			}
		}
	}
	return unmarkedSum * d
}

func checkWin(draws []int, i, j, k, d int) int {
	var sum int
	count := 0
	for z := 0; z < 5; z++ {
		if visited(draws, board[i][j][z], d) {
			sum += int(math.Abs(float64(board[i][j][z])))
			count++
		}
	}

	if count == 5 {
		return getScore(draws, i, d)
	}

	sum = 0
	count = 0
	for z := 0; z < 5; z++ {
		if visited(draws, board[i][z][k], d) {
			sum += int(math.Abs(float64(board[i][z][k])))
			count++
		}
	}

	if count == 5 {
		return getScore(draws, i, d)
	}

	return -1
}
