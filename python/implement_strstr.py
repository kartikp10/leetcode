# https://leetcode.com/problems/implement-strstr/submissions/

class Solution:
    def __init__(self):
        pass

    def strStr(self, haystack: str, needle: str) -> int:
        if needle.strip() == "":
            return 0
        needle_len = len(needle)
        haystack_len = len(haystack)
        if needle_len > haystack_len:
            return -1
        index = 0
        while index + needle_len <= haystack_len:
            if haystack[index: needle_len + index] == needle:
                return index
            index += 1
        return -1


if __name__ == '__main__':
    solution = Solution().strStr("a", "a")
    print(solution)