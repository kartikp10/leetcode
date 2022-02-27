# https://leetcode.com/problems/remove-duplicates-from-sorted-array/

from typing import List


class Solution:
    def __init__(self):
        pass

    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return 1
        unique_len = 1
        for i in range(len(nums) -1):
            if nums[i] != nums[i+1]:
                nums[unique_len] = nums[i+1]
                unique_len += 1
        return unique_len


if __name__ == '__main__':
    solution = Solution().removeDuplicates([1, 1, 2, 4, 7, 8, 8, 9])
    print(solution)

