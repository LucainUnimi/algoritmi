package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isLowest(heightmap [][]int, posx, posy int) bool {
	return (posx-1 < 0 || heightmap[posx-1][posy] > heightmap[posx][posy]) &&
		(posx+1 == len(heightmap) || heightmap[posx+1][posy] > heightmap[posx][posy]) &&
		(posy-1 < 0 || heightmap[posx][posy-1] > heightmap[posx][posy]) &&
		(posy+1 == len(heightmap[0]) || heightmap[posx][posy+1] > heightmap[posx][posy])
}

func check(heightmap [][]int, posx, posy, risklevel int) int {
	actualx := posx / len(heightmap[0])
	actualy := posy % len(heightmap[0])
	if actualx == len(heightmap)-1 && actualy == len(heightmap[0])-1 {
		return risklevel
	}
	if isLowest(heightmap, actualx, actualy) {
		return check(heightmap, (posx + 1), (posy + 1), risklevel+1+heightmap[actualx][actualy])
	}
	return check(heightmap, (posx + 1), (posy + 1), risklevel)
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
	/*
		for i := 0; i < len(heightmap); i++ {
			for j := 0; j < len(heightmap[0]); j++ {
				fmt.Printf("%d ", heightmap[i][j])
			}
			fmt.Println()
		}
	*/
	fmt.Printf("check(heightmap, len(heightmap)-1, len(heightmap[0])-1, 0): %v\n", check(heightmap, 0, 0, 0))
}
