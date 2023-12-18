package main

import "fmt"

// for integer array
func int_array_once(array []int) {
	fmt.Println("This is a array:")
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func int_array_with_index(array []int) {
	fmt.Println("This is a array with index:")
	for index, value := range array {
		fmt.Println("[", index, "] ->", value)
	}
}

// for string array
func str_array_once(array []string) {
	fmt.Println("This is a array:")
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func str_array_with_index(array []string) {
	fmt.Println("This is a array with index:")
	for index, value := range array {
		fmt.Println("[", index, "] ->", value)
	}
}

func test_array() {
	int_arr := []int{1, 2, 3, 4, 5}
	str_arr := []string{"Google", "Apple", "Microsoft", "Amazon", "Facebook"}

	fmt.Println("----------")
	int_array_once(int_arr)
	fmt.Println("----------")
	int_array_with_index(int_arr)
	fmt.Println("----------")
	str_array_once(str_arr)
	fmt.Println("----------")
	str_array_with_index(str_arr)
}
