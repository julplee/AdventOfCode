package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	str := string(dat)

	s := strings.Split(str, ",")

	var slice []int
	for _, digit := range s {
		stoi, _ := strconv.Atoi(digit)
		slice = append(slice, stoi)
	}

	// Part 1 result
	slice1 := make([]int, len(slice))
	copy(slice1, slice)

	// Program instructions to overwrite
	slice1[1] = 12
	slice1[2] = 2

	processOpcode(slice1)
	fmt.Println(arrayToString(slice1, ","))

	// Part 2 result
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			slice2 := make([]int, len(slice))
			copy(slice2, slice)

			slice2[1] = noun
			slice2[2] = verb

			processOpcode(slice2)

			if slice2[0] == 19690720 {
				fmt.Println(slice2)
				fmt.Println("noun=", noun)
				fmt.Println("verb=", verb)
				fmt.Println("100*noun+verb=", 100*noun+verb)

				break
			}

		}
	}
}

func processOpcode(slice []int) {
	for i := 0; i < len(slice)-3; i = i + 4 {
		index := slice[i]

		if index == 99 {
			break
		}

		firstPointer := slice[i+1]
		secondPointer := slice[i+2]
		writePointer := slice[i+3]

		var third int
		if index == 1 {
			third = slice[firstPointer] + slice[secondPointer]
		} else {
			third = slice[firstPointer] * slice[secondPointer]
		}

		slice[writePointer] = third
	}
}

func arrayToString(slice []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(slice), " ", delim, -1), "[]")
}
