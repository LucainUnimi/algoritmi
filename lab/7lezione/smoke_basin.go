package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_lower(m [][]int, x, y int) bool {
	if x != 0 && y != 0 && x != len(m)-1 && y != len(m[0])-1 {
		if m[x][y] < m[x+1][y] && m[x][y] < m[x-1][y] && m[x][y] < m[x][y+1] && m[x][y] < m[x][y-1] {
			return true
		}
	}
	return false
}

func risklevelaux(matrix [][]int, x, y, maxX, maxY int) int {
	if (x+1)%maxX == maxX-1 && (y+1)/maxY == maxY-1 {
		return 0
	} else {
		if is_lower(matrix, x%maxX, y/maxY) {
			return 1 + risklevelaux(matrix, x+1, y+1, maxX, maxY)
		} else {
			return risklevelaux(matrix, x+1, y+1, maxX, maxY)
		}
	}
}

func risklevel(matrix [][]int) int {
	return risklevelaux(matrix, 0, 0, len(matrix), len(matrix[0]))
}

func main() {
	var heightmap [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var lineint []int
		line := strings.Split(scanner.Text(), "")
		for _, v := range line {
			n, _ := strconv.Atoi(v)
			lineint = append(lineint, n)
		}
		heightmap = append(heightmap, lineint)
	}
	fmt.Printf("heightmap: %v\n", heightmap)
}
