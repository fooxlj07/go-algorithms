package main

import (
	"fmt"
	"time"
)

func bubbleSort(arr []int) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println("used", time.Since(start))
	}(start)
	var tmp int
	for i, element := range arr {
		for j := i; j < len(arr); j++ {
			if element > arr[j] {
				tmp = arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
	fmt.Println(arr)
}

func selectionSort(arr []int) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println("used", time.Since(start))
	}(start)
	var minIndex int
	for i := 0; i < len(arr); i++ {
		minIndex = i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		tmp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}
	fmt.Println(arr)
}

func insertSort(arr []int) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println("used", time.Since(start))
	}(start)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] >= arr[i] {
			continue
		}
		for j := i + 1; j >= 0; j-- {
			if arr[i] < arr[j] {
				arr[j] = arr[j-1]
			}
		}
	}
	fmt.Println(arr)
}

func shellSort(arr []int) {
	var tmp int
	for gap := len(arr) / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len(arr); i += gap {
			if arr[i-gap] > arr[i] {
				tmp = arr[i]
				arr[i] = arr[i-gap]
				arr[i-gap] = tmp
			}
		}
	}
	fmt.Println(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, left[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func quickSort(arr []int) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println("used", time.Since(start))
	}(start)
	partition(arr, 0, len(arr)-1)
}

func partition(arr []int, left, right int) {
	if right < left {
		return
	}
	i, j, tmp := left+1, right, 0
	for i < j {
		if arr[j] > arr[left] {
			j--
		}
		if arr[i] < arr[left] {
			i++
		}
		if arr[j] <= arr[left] && arr[i] > arr[left] {
			tmp = arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
		}
	}
	tmp = arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
	partition(arr, j+1, right)
	partition(arr, left, j-1)
}

func heapSort(arr []int) {
	start := time.Now()
	defer func(start time.Time) {
		fmt.Println("used", time.Since(start))
	}(start)
	for i := len(arr) - 1; i > 0; i-- {
		for j := i; j/2 >= 0; j -= 2 {
			biggerIndex := j
			if j%2 == 1 && arr[j] < arr[j-1] {
				biggerIndex = j - 1
			}
			if arr[biggerIndex] > arr[j/2-1] {
				arr[j/2-1], arr[biggerIndex] = arr[biggerIndex], arr[j/2-1]
			}
			fmt.Println("@@@   ", arr)
		}
		arr[i], arr[0] = arr[0], arr[i]
		fmt.Println(arr)
	}
	//	fmt.Println(arr)
}
