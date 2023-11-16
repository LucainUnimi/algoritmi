package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(heightmap [][]int, posx, posy int, all_size []int, size int) []int {
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
