package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	coords := make(map[string]string)

	scanner := bufio.NewScanner(file)

	var wire string
	for scanner.Scan() {
		str := scanner.Text()
		pathArray := strings.Split(str, ",")

		cursorX := 0
		cursorY := 0

		if wire == "" {
			wire = "A"
		} else {
			wire = "B"
		}

		fmt.Println(pathArray)
		for i := 0; i < len(pathArray); i++ {
			path := pathArray[i]

			pathDirection := path[0:1]
			pathDistance, _ := strconv.Atoi(path[1:len(path)])

			if pathDirection == "R" {
				for i := 1; i <= pathDistance; i++ {
					s := strconv.Itoa(cursorX+i) + ";" + strconv.Itoa(cursorY)
					if !strings.Contains(coords[s], wire) {
						coords[s] = coords[s] + wire
					}
				}

				cursorX = cursorX + pathDistance
			} else if pathDirection == "L" {
				for i := 1; i <= pathDistance; i++ {
					s := strconv.Itoa(cursorX-i) + ";" + strconv.Itoa(cursorY)
					if !strings.Contains(coords[s], wire) {
						coords[s] = coords[s] + wire
					}
				}

				cursorX = cursorX - pathDistance
			} else if pathDirection == "U" {
				for i := 1; i <= pathDistance; i++ {
					s := strconv.Itoa(cursorX) + ";" + strconv.Itoa(cursorY+i)
					if !strings.Contains(coords[s], wire) {
						coords[s] = coords[s] + wire
					}
				}

				cursorY = cursorY + pathDistance
			} else if pathDirection == "D" {
				for i := 1; i <= pathDistance; i++ {
					s := strconv.Itoa(cursorX) + ";" + strconv.Itoa(cursorY-i)
					if !strings.Contains(coords[s], wire) {
						coords[s] = coords[s] + wire
					}
				}

				cursorY = cursorY - pathDistance
			}
		}

		fmt.Println(coords)
	}

	var manhattanDists []int

	for k, coord := range coords {
		if coord == "AB" {
			s := strings.Split(k, ";")
			fmt.Println(s)

			x, _ := strconv.Atoi(s[0])
			y, _ := strconv.Atoi(s[1])
			xdist := abs(x)
			ydist := abs(y)
			manhattanDists = append(manhattanDists, int(xdist+ydist))
		}
	}

	fmt.Println(manhattanDists)
	fmt.Println(min(manhattanDists))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	x := arr[0]

	for i := 1; i < len(arr); i++ {
		y := arr[i]
		if y < x {
			x = y
		}
	}

	return x
}
