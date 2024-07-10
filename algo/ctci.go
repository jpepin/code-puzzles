package algo

import (
	"fmt"
)

func Run() {
	k := 6
	nums := []int{1, 7, 5, 9, 2, 12, 3}
	k_diff_brute(k, nums)
	k_diff_hash(k, nums)

}

// CTCI p.66-68, Optimize and Solve
// For an array of distinct integer values,
// find the count of pairs with a difference of k
func k_diff_brute(k int, nums []int) {
	// brute force, go through array
	// and check each value against the others
	k_pairs := 0
	for i, inum := range nums {
		// we're at the end of the array
		if i+1 == len(nums) {
			break
		}
		// the other values in the array
		sub := nums[i+1:]
		for _, jnum := range sub {
			if inum+k == jnum {
				k_pairs++
			}
			if inum-k == jnum {
				k_pairs++
			}
		}
	}
	fmt.Printf("pairs: %d\n", k_pairs)
}

func k_diff_hash(k int, nums []int) {
	// avoid a nested loop by storing the integers in
	// a hash map so that we only loop through once
	k_pairs := 0
	val_map := make(map[int]bool)
	for _, num := range nums {
		if _, ok := val_map[num+k]; ok {
			k_pairs++
		}
		if _, ok := val_map[num-k]; ok {
			k_pairs++
		}
		// add num after checks to avoid checking
		// against self
		val_map[num] = true
	}
	fmt.Printf("pairs: %d\n", k_pairs)
}
