/*
	Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

	Example 1:
	Input: nums = [1,1,1,2,2,3], k = 2
	Output: [1,2]

	Example 2:
	Input: nums = [1], k = 1
	Output: [1]

	Example 3:
	Input: arr[] = {3, 1, 4, 4, 5, 2, 6, 1}, k = 2
	Output: [4,1]
*/

package main

import (
	"fmt"
	"sort"
	"strings"

	// "reflect"
	"unicode"
)

func decryptMsg(str string) {

	// Remove all numbers 0 to 9 from the string and return the reversed message.
	var builder strings.Builder

	for _, r := range strings.ReplaceAll(str, " ", "") {
		if !unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}

	decrypted := builder.String()
	var reversed string

	for i := len(decrypted) - 1; i >= 0; i-- {
		reversed += string(decrypted[i])
	}

	fmt.Println(reversed)
}

func topKFrequent(nums []int, k int) []int {

	var want map[int]int
	want = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		want[nums[i]] = want[nums[i]] + 1
	}

	keys := make([]int, len(want))

	i := 0
	for k := range want {
		keys[i] = k
		i++
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return want[keys[i]] > want[keys[j]]
	})
	fmt.Println(keys)

	result := make([]int, k)
	x := 0
	for y := range keys {
		result[x] = keys[y]
		x++
		if x == k {
			break
		}
	}
	fmt.Println(result)

	return result
}

func main() {

	topKFrequent([]int{3, 1, 4, 4, 5, 2, 6, 1}, 2)

	// decryptMsg("  036o  l  leH1297  ")

	// type test struct {
	// 	nums []int
	// 	k    int
	// 	want []int
	// }

	// tests := []test{
	// 	{
	// 		[]int{1, 1, 1, 2, 2, 3},
	// 		2,
	// 		[]int{1, 2},
	// 	},
	// 	{
	// 		[]int{1, 1, 1, 2, 2, 3, 3, 3, 3},
	// 		2,
	// 		[]int{1, 3},
	// 	},
	// 	{
	// 		[]int{1},
	// 		1,
	// 		[]int{1},
	// 	},
	// 	{
	// 		[]int{3, 1, 4, 4, 5, 2, 6, 1},
	// 		2,
	// 		[]int{4, 1},
	// 	},
	// }

	// for i, tt := range tests {
	// 	got := topKFrequent(tt.nums, tt.k)
	// 	res := "PASSED"
	// 	if !reflect.DeepEqual(got, tt.want) {
	// 		res = "FAILED"
	// 	}
	// 	fmt.Printf("\n%s case=%d got=%v want=%v", res, i, got, tt.want)
	// }
}
