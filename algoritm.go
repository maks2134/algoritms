package main

import "fmt"

func BinarySearchInt(list []int, value int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == value {
			return mid
		}
		if guess > value {
			high = mid - 1
		}
		if guess < value {
			low = mid + 1
		}
	}
	return -1
}

func DefaultSearch(list []int, value int) int {
	low := 0
	for low <= len(list)-1 {
		if list[low] == value {
			return low
		}
		if list[low] != value {
			low += 1
		}
	}
	return -1
}

func selectionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
	return list
}

func bubleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func factorial(value int) int {
	if value == 0 {
		return 1
	}
	return value * factorial(value-1)
}

func main() {
	//array := []int{56, 3, 421, 23, -5}
	//fmt.Println(BinarySearchInt(array, 3))
	//newarray := bubleSort(array)
	//for i := 0; i < len(newarray); i++ {
	//	fmt.Print(newarray[i], " ")
	//}

	fmt.Println(factorial(5))
}
