from typing import Optional


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


def lst2link(lst):
    cur = dummy = ListNode(0)
    for e in lst:
        cur.next = ListNode(e)
        cur = cur.next
    return dummy.next


def printLst(lst: Optional[ListNode]):
    arr = list()
    while lst.next:
        arr.append(lst.val)
        lst = lst.next
    arr.append(lst.val)
    print(arr)


class Solution:
    @classmethod
    def mergeTwoLists(cls, l1: Optional[ListNode], l2: Optional[ListNode]):
        current = head = ListNode(0)
        while l1 and l2:
            if l1.val < l2.val:
                current.next = l1
                l1 = l1.next
            else:
                current.next = l2
                l2 = l2.next
            current = current.next
        current.next = l1 or l2
        return head.next


if __name__ == '__main__':
    l1 = lst2link([1, 2, 3, 4])
    l2 = lst2link([2, 4, 6])
    solution = Solution.mergeTwoLists(l1=l1, l2=l2)
    printLst(solution)
