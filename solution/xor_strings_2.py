# https://www.hackerrank.com/challenges/one-month-preparation-kit-strings-xor/problem
def strings_xor(s, t):
    res = ""
    for i in range(len(s)):
        if s[i] == t[i]:
            res += '0'
        else:
            res += '1'
    return res
