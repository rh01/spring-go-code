#coding: utf-8
import sys

def solution(arr, target):
    dp = [0] * (target+1)
    arr.sort()
    print(arr)

    res = 0
    for i, j in enumerate(arr):
        target -= j
        if target < 0:
            return res
        res+=1
    return res


if __name__ == "__main__":
    params = sys.stdin.readline().strip().split()
    n, k = int(params[0]), int(params[1])
    arr =  sys.stdin.readline().strip().split()
    intarr = map(int, arr)
    print(solution(intarr[:n], k))
