package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var allSquaresFlat = [][]int{
	{2, 9, 4, 7, 5, 3, 6, 1, 8},
	{6, 1, 8, 7, 5, 3, 2, 9, 4},
	{6, 7, 2, 1, 5, 9, 8, 3, 4},
	{8, 3, 4, 1, 5, 9, 6, 7, 2},
	{8, 1, 6, 3, 5, 7, 4, 9, 2},
	{4, 9, 2, 3, 5, 7, 8, 1, 6},
	{4, 3, 8, 9, 5, 1, 2, 7, 6},
	{2, 7, 6, 9, 5, 1, 4, 3, 8},
}

func squareCost(s1, s2 []int) int {
	total := 0
	for i := 0; i < len(s1); i++ {
		total += int(math.Abs(float64(s1[i] - s2[i])))
	}
	return total
}

func formingMagicSquare(s [][]int) int {
	flatS := []int{}
	for i := 0; i < 3; i++ {
		flatS = append(flatS, s[i]...)
	}

	minCost := math.MaxInt32
	for _, square := range allSquaresFlat {
		cost := squareCost(flatS, square)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	s := [][]int{}

	for i := 0; i < 3; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		numStrs := strings.Split(line, " ")
		row := []int{}
		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			row = append(row, num)
		}
		s = append(s, row)
	}

	result := formingMagicSquare(s)
	fmt.Fprintf(writer, "%d\n", result)
}
