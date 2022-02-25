# https://leetcode.com/problems/valid-parentheses/

from typing import List


class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        openings = ['(', '[', '{']
        closings = [')', ']', '}']
        valid = False
        for i, c in enumerate(s):
            if i == 0 and c not in openings:
                return False
            if c in closings and len(stack) == 0:
                return False
            if c in openings:
                stack.append(c)
            elif c == closings[openings.index(stack[-1])]:
                valid = True
                stack.pop()
            # if c is not another opening or closing i=0, string is invalid
            else:
                return False
        if len(stack) > 0:
            return False
        return valid


if __name__ == '__main__':
    a = Solution().isValid("(){}}{")
    print(a)
