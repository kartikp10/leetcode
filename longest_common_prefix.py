from typing import List


def longest_common_prefix(strs: List[str]) -> str:
    pf = strs[0]
    for s in strs[1:]:
        if len(s) < len(pf):
            pf = pf[:len(s)]
        for i in range(len(pf)):
            if i == 0:
                if s[i] != pf[i]:
                    return ""
            if s[i] != pf[i]:
                pf = pf[:i]
                break
    return pf


if __name__ == '__main__':
    pf = longest_common_prefix(["a", "b"])
    print(pf)

