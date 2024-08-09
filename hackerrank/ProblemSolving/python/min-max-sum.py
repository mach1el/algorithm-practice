#!/bin/python3

import os
import sys

def miniMaxSum(arr):
  total = sum([x for x in arr])
  print(total - max(arr), total - min(arr))
        
if __name__ == '__main__':
  arr = list(map(int, input().rstrip().split()))

  miniMaxSum(arr)