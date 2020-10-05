package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func newSlice(slice []int, a int, b int) []int {
	newS := make([]int, len(slice))

	//for a; a <= b; a++ {
	//		newS[a] = slice[a]
	//	}
	return newS
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	var i = 0
	for _, s := range slice {
		slice[i] = f(s)
		i++
	}
}

func mapArray(f func(a int) int, array [5]int) {
	for s := range array {
		array[s] = f(s)
	}
}

func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, intSlice)
	fmt.Println(intSlice)
	intArr := [5]int{1, 2, 3, 4, 5}
	mapArray(addOne, intArr)
	fmt.Println(intArr)

	sl := newSlice(intSlice, 2, 4)
	fmt.Println(sl)
}
