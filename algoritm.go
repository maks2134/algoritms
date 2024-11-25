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

func SelectionSort(list []int) []int {
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

func BubleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

func Factorial(value int) int {
	if value == 0 {
		return 1
	}
	return value * Factorial(value-1)
}

func RecursiveSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + RecursiveSum(arr[1:])
}

func RecursiveCount(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + RecursiveCount(arr[1:])
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}

func InsertionSort(arr []int) []int {
	i := 1
	for i < len(arr) {
		j := i
		for j >= 1 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
		i++
	}
	return arr
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])
	return merge(first, second)
}

func DoubleToArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
	return arr
}

func DoubleFirstElementArray(arr []int) []int {
	arr[0] *= 2
	return arr
}

func main() {
	//fmt.Println(BinarySearchInt(array, 3))
	//newarray := bubleSort(array)
	//for i := 0; i < len(newarray); i++ {
	//	fmt.Print(newarray[i], " ")
	//}
	arr := []int{1, 3, 5, 64, 64, 5, 423, -123, 5834, 934, 0, -423}
	//arr = QuickSort(arr, 0, len(arr)-1)
	arr = MergeSort(arr)
	fmt.Println(arr)
}
