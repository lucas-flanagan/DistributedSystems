package main

import (
	"fmt"
	"time"
)

func Get_Integer() int {
	var x int
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&x)
	return x
}

func Build_Data_Types(x int) (time.Duration, time.Duration) {

	// Initialize slice
	var slce = make([]int, x+1)
	for i, _ := range slce {
		slce[i] = i
	}

	// Increment slice values
	start := time.Now()
	for i, _ := range slce {
		slce[i] += 1
	}
	slce_duration := time.Since(start)

	// Initialize mp
	mp := make(map[int]int)
	for i := 0; i < x+1; i++ {
		mp[i] = i
	}

	// Increment map values
	start = time.Now()
	for key, val := range mp {
		mp[key] = val + 1
	}
	mp_duration := time.Since(start)

	return slce_duration, mp_duration

}

func main() {
	x := Get_Integer()
	slce, mp := Build_Data_Types(x)
	fmt.Println("Incrementation duration for slice:", slce)
	fmt.Println("Incrementation duration for map:", mp)
}
