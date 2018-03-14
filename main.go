package main

import (
	"time"
	"fmt"
	"sort"
	"math/rand"
	// https://github.com/0xAX/go-algorithms
)

func createArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i <= n - 1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

func sortInt(arr []int) {
	start := time.Now()

	sort.Ints(arr)

	end := time.Now()
	fmt.Println(len(arr), end.Sub(start))
}
func mSort(arr []int) {
	start := time.Now()
	MergeSort(arr)
	end := time.Now()
	fmt.Println(len(arr), end.Sub(start))
}
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := (len(arr)) / 2

	return Merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))

}
func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func gnomeSort(arr []int) {
	start := time.Now()
	for d := int(len(arr)/2); d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			for j := i; j >= d && arr[j-d] > arr[j]; j -= d {
				arr[j], arr[j-d] = arr[j-d], arr[j]
			}
		}
	}
	end := time.Now()
	fmt.Println(len(arr), end.Sub(start))
}

func mSortGo() {
	arr1000 := createArray(1000)
	mSort(arr1000)

	arr10000 := createArray(10000)
	mSort(arr10000)

	arr100000 := createArray(100000)
	mSort(arr100000)

	arr1000000 := createArray(1000000)
	mSort(arr1000000)

	arr10000000 := createArray(10000000)
	mSort(arr10000000)

	arr100000000 := createArray(100000000)
	mSort(arr100000000)
}

func nativeQSort() {
	arr1000 := createArray(1000)
	sortInt(arr1000)

	arr10000 := createArray(10000)
	sortInt(arr10000)

	arr100000 := createArray(100000)
	sortInt(arr100000)

	arr1000000 := createArray(1000000)
	sortInt(arr1000000)

	arr10000000 := createArray(10000000)
	sortInt(arr10000000)

	arr100000000 := createArray(100000000)
	sortInt(arr100000000)
}

func gnomeSortGo() {
	arr1000 := createArray(1000)
	gnomeSort(arr1000)

	arr10000 := createArray(10000)
	gnomeSort(arr10000)

	arr100000 := createArray(100000)
	gnomeSort(arr100000)

	arr1000000 := createArray(1000000)
	gnomeSort(arr1000000)

	arr10000000 := createArray(10000000)
	gnomeSort(arr10000000)

	arr100000000 := createArray(100000000)
	gnomeSort(arr100000000)
}

func main() {
	gnomeSortGo()
}
