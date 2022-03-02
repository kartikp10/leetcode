from typing import List


class Solution:
    def __init__(self):
        pass

    def plusOne(self, digits: List[int]) -> List[int]:
        digits[-1] += 1
        for i in range(len(digits) - 1, -1, -1):
            if digits[i] == 10 and i > 0:
                digits[i] = 0
                digits = self.plusOne(digits[0:i]) + digits[i:]
        if digits[0] == 10:
            return [1, 0] + digits[1:]
        return digits


if __name__ == '__main__':
    solution1 = Solution().plusOne([9])
    print(solution1)
    solution2 = Solution().plusOne([8, 9, 9])
    print(solution2)
    solution3 = Solution().plusOne([9])
    print(solution3)
