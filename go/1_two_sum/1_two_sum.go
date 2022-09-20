package leetcode

func twoSum(nums []int, target int) []int {
	// hash map
	comp := make(map[int]int)

	for i, num := range nums {
		_, ok := comp[num]
		if ok {
			return []int{comp[num], i}
		}
		comp[target-num] = i
	}
	return []int{}
}
