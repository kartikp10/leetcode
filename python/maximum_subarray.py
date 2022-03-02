# https://leetcode.com/problems/maximum-subarray/
# Kadane’s Algorithm

from typing import List


class Solution:
    def __init__(self):
        pass

    def maxSubArray(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        local_max = global_max = nums[0]
        for i, n in enumerate(nums[1:], start=1):
            local_max = max(n, local_max + n)
            global_max = max(local_max, global_max)
        return global_max


if __name__ == '__main__':
    solution = Solution().maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4])
    print(solution)
