package main

import "fmt"

// 左闭右闭
func search(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {

	var num [6]int = [...]int{-1, 0, 3, 5, 9, 12}
	res := search(num[:], 9)
	fmt.Println(res)
}

/*
坚持根据查找区间的定义来做边界处理
要先规定好区间规则，决定了mid的定位
*/
