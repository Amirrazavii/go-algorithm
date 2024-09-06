package main

import (
	"fmt"

	Array "algo/array"
)

func main() {
	// var a = Array.MergeAlternately("fghtd", "ds")
	// fmt.Println(a)
	// 	var a = Array.GcdOfStrings("ABCABC", "ABC")
	// fmt.Println(a)
	var arr = []int{6, 5, 3, 5, 3}
	var a = Array.KidsWithCandies(arr, 9)
	fmt.Println(a)

}
