// package main

// import "fmt"

// func main() {
// 	nums := [4]int{3, 3, 1, 2}
// 	target := 6
// 	l := 9
// 	fmt.Println(twoSum(nums[0:], target))
// 	fmt.Println(massiv(l))
// }

// func massiv(l int) []int {
// 	golem := make([]int, l)
// 	for f := 0; f < len(golem); f++ {
// 		golem[f] = f
// 	}
// 	return golem
// }

// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := 1 + i; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return nums
// }

package main

import "fmt"

func main() {
	array := [7]int{2, 6, 8, 12, 86, 123, 3657}
	var lookingFor int = 3657
	fmt.Println(binarySearch(array[0:], lookingFor))
}

func binarySearch(array []int, lookingFor int) int {
	var mid int
	min := 0
	high := len(array) - 1

	for min <= high {
		mid = (min + high) / 2
		if array[mid] == lookingFor {
			return mid
		}
		if array[mid] > lookingFor {
			high = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}

// package main

// import "fmt"

// func main() {
// 	var i int = 0
// 	for i < 10 {
// 		i += 1
// 		fmt.Println(i)
// 	}
