package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalFuel := 0

	for scanner.Scan() {
		dist, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Print(err)
		}

		fuelForVessel := dist/3 - 2
		fuelForFuel := calculateFuelNeededForFuel(fuelForVessel)

		// Total fuel for vessel = Total fuel plus third of the distance rounded down minus two. And consider fuel needed for fuel weight
		totalFuel = totalFuel + fuelForVessel + fuelForFuel
	}

	fmt.Println(totalFuel)
}

func calculateFuelNeededForFuel(fuelForVessel int) int {
	fuelForFuel := 0

	fuelForVessel = fuelForVessel/3 - 2

	for fuelForVessel > 0 {
		fuelForFuel = fuelForFuel + fuelForVessel

		fuelForVessel = fuelForVessel/3 - 2
	}

	return fuelForFuel
}
