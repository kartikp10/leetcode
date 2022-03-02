class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 0:
            return 0
        if x == 1:
            return 1
        left = 0
        right = x
        while right >= left:
            mid = (left + right) // 2
            if mid * mid > x:
                right = mid -1
            elif mid * mid <= x < (mid + 1)*(mid + 1):
                return mid
            else:
                left = mid + 1
        return x


if __name__ == '__main__':
    solution = Solution().mySqrt(4)
    print(solution)