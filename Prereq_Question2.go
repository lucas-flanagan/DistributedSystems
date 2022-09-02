package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Get_Integer() int {
	var x int
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&x)
	return x
}

func Generate_Randoms(x int) []int {

	// Initialize slice and populate with random integers from 0-100
	// Source: gobyexample.com random-numbers
	slce := make([]int, x)
	for i := 0; i < x; i++ {
		slce[i] = rand.Intn(100)
	}
	return slce
}

type byValue []int

func (a byValue) Len() int {
	return len(a)
}

func (a byValue) Less(i, j int) bool {
	return i < j
}

func (a byValue) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Sort_Slice(slce []int) time.Duration {
	start := time.Now()
	sort.Sort(byValue(slce))
	sort_durr := time.Since(start)

	return sort_durr
}

func Stable_Slice(slce []int) time.Duration {
	start := time.Now()
	sort.Stable(byValue(slce))
	stable_durr := time.Since(start)

	return stable_durr
}

func main() {
	x := Get_Integer()
	slce := Generate_Randoms(x)
	Sort_Slice := Sort_Slice(slce)
	Stable_Slice := Stable_Slice(slce)
	fmt.Println("Sort duration for Slice: ", Sort_Slice)
	fmt.Println("Stable duration for Slice: ", Stable_Slice)
}
