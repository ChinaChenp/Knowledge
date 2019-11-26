#最长不重复子串


def LongestSubstring(s):
    if len(s) <= 1:
        return len(s)

    sets = set()
    i = j = 0
    max_len = 0
    while i < len(s) and j < len(s):
        if s[j] not in sets:
            sets.add(s[j])
            j += 1
            max_len = max(max_len, j - i)
        else:
            sets.remove(s[i])
            i += 1
    return max_len


re = LongestSubstring("abcadaeffghi")
print(re)