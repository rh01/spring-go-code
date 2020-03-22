#coding: utf-8

import sys

def split(A, B):
    # B 可以看成空箱子，允许为空 B - 1
    # B- 2刀
    return (A + 1) ** (B-2)

if __name__ == "__main__":
    n = int(input())
    # print(n)
    for i in range(n):
        s = sys.stdin.readline().strip()
        arr = s.split()
        A, B = int(arr[0]), int(arr[1])
        print(split(A, B))
