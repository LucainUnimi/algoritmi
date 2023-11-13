package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
