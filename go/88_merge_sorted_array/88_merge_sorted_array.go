package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	p1 := m - 1
	p2 := n - 1
	p3 := m + n - 1

	for ; p2 >=0; p3-- {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p3] = nums1[p1]
			p1--
		} else {
			nums1[p3] = nums2[p2]
			p2--
		}
	}
    return nums1
}

