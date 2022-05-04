package main

import (
	"fmt"
)

// Khởi tạo mảng rank bằng 0
func khoitaorank(arr []int, ranklist chan []int) {
	rankkhoitao := make([]int, len(arr))
	// var rankkhoitao []int
	// for i := 0; i < len(arr)-1; i++ {
	// 	rankkhoitao = append(rankkhoitao, 0)
	// }
	ranklist <- rankkhoitao
}

// rank sort
func rankSort(arr []int, ranklistkhoitao []int, ranksort_trave chan []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] >= arr[j] {
				ranklistkhoitao[i]++
			} else {
				ranklistkhoitao[j]++
			}
		}
	}
	ranksort_trave <- ranklistkhoitao
}

// Sắp xếp mảng
func sortlist(arr []int, ranklist []int, arrsort chan []int) {
	kq := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		kq[ranklist[i]] = arr[i]
	}
	arrsort <- kq
}

// in
func in(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func main() {
	var slice1 = []int{1, 2, 5, 2, 1, 60, 3}
	rank_list := make(chan []int)
	go khoitaorank(slice1, rank_list)
	//
	rank := <-rank_list
	ranksort_trave := make(chan []int)
	go rankSort(slice1, rank, ranksort_trave)
	//
	ranksort := <-ranksort_trave
	arr_sort := make(chan []int)
	go sortlist(slice1, ranksort, arr_sort)
	//
	kq := <-arr_sort
	fmt.Print("Mang ban dau: ")
	in(slice1)
	fmt.Print("Rank sort: ")
	in(ranksort)
	fmt.Print("Arr Sort: ")
	in(kq)
}
