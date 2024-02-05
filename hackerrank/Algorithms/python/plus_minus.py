#!/bin/python3

import os
import sys

def plusMinus(n,arr):
  pos = 0
  neg = 0
  zero = 0
  
  for x in arr:
    if x > 0: pos += 1
    elif x < 0: neg += 1
    else: zero += 1

  print("%.6f" % (pos/len(arr)))
  print("%.6f" % (neg/len(arr)))
  print("%.6f" % (zero/len(arr)))

if __name__ == '__main__':
  n = int(input())

  arr = list(map(int, input().rstrip().split()))

  plusMinus(n,arr)