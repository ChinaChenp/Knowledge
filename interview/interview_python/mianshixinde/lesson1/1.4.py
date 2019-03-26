#回文字符串
#回文，英文palindrome，指一个顺着读和反过来读都一样的字符串，比如madam、我爱我，这样的短句在智力性、趣味性和艺术性上都颇有特色，中国历史上还有很多有趣的回文诗。
#那么，我们的第一个问题就是：判断一个字串是否是回文？

def IsPalindrome(str):
    beg, end = 0, len(str) - 1

    while beg < end:
        if str[beg] != str[end]:
            return False
        beg += 1
        end -= 1
    return True

print(IsPalindrome("fqfsajf"))
print(IsPalindrome("asdfdsa"))
print(IsPalindrome("1234565432"))