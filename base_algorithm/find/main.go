package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	check()
	//t()
}

func t() {
	arr1 := []int{-781147, -752714, -732439, 44087, 466503, 840817}
	arr2 := []int{-910124, -341771, -134581, 87768, 172506, 287849}
	fmt.Println("Array1:", arr1)
	fmt.Println("Array2:", arr2)
	f := findMedianSortedArrays(arr1, arr2)
	fok := findMedianSortedArraysOK(arr1, arr2)
	fmt.Println("test result:", f, fok)
}

func checkOK() {
	arr1 := []int{1}
	arr2 := []int{}
	fok := findMedianSortedArraysOK(arr1, arr2)
	fmt.Println("Array1:", arr1)
	fmt.Println("Array2:", arr2)
	fmt.Println("test result:", fok)
}

func check() {
	rand.Seed(time.Now().UnixNano())
	rangeNum := 100000
	listMaxSize := 10
	start := time.Now()
	for i := 0; i < rangeNum; i++ {
		//n1 := rand.Intn(3)
		n1 := rand.Intn(listMaxSize)
		arr1 := randArray(uint(n1))
		n2 := rand.Intn(listMaxSize)
		arr2 := randArray(uint(n2))
		sort.Ints(arr1)
		sort.Ints(arr2)
		f := findMedianSortedArrays(arr1, arr2)
		//f := getArrayMidle(arr1, arr2)
		fok := findMedianSortedArraysOK(arr1, arr2)
		if f != fok {
			fmt.Println("Array1:", arr1)
			fmt.Println("Array2:", arr2)
			fmt.Println(f, fok)
			panic("err")
		}
		if i%(rangeNum/100) == 0 {
			fmt.Println(i/(rangeNum/100), "%")
		}
	}
	fmt.Println("speed:", time.Since(start).String())
}
