package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	slice = append(slice, slice...)
	return nil
}

func mapSlice(f func(a int) int, slice []int) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}

func mapArray(f func(a int) int, array [3]int) {
	for i, v := range array {
		array[i] = f(v)
	}
}

func main() {
	instSlice := []int{1, 2, 3, 4, 5}
	mapSlice(addOne, instSlice)
	fmt.Printf("Slice: %v\n", instSlice)

	instArray := [3]int{1, 2, 3}
	mapArray(addOne, instArray)
	fmt.Printf("Array: %v\n", instSlice)
	////////
	newSlice := instSlice[1:3]
	mapSlice(square, newSlice)
	fmt.Printf("Slice sliced squared: %v\n", newSlice)
	////////
	instSlice = double(instSlice)
	fmt.Printf("Double 'instSlice': %v\n", instSlice)

}
