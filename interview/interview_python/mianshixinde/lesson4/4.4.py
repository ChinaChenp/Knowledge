# 字符串查找
# 暴力求解

def MatchStr(str1, str2):
    len1 = len(str1)
    len2 = len(str2)

    i, j = 0, 0
    while i < len1 and j < len2:
        if str1[i] == str2[j]:
            i += 1
            j += 1
        else:
            #i 需要回退回去
            i = i - j + 1
            j = 0
    if j == len2:
        return True
    return False

str1 = list("BBC ABCDAB ABCDABCDABDE")
str2 = list("ABCDABDEF")
print(MatchStr(str1, str2))