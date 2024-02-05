#!/bin/python3

import math
import os
import random
import re
import sys

def getTotalX(a, b):
  count = 0
  min_value, max_value = min(a), max(b)
  for n in range(min_value, max_value+1):
    factors_1 = set(map(lambda x:n%x, a))
    factors_2 = set(map(lambda x:x%n, b))
    if max(factors_1) == 0 and max(factors_2) == 0:
      count +=1

  return count

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  first_multiple_input = input().rstrip().split()

  n = int(first_multiple_input[0])

  m = int(first_multiple_input[1])

  arr = list(map(int, input().rstrip().split()))

  brr = list(map(int, input().rstrip().split()))

  total = getTotalX(arr, brr)

  fptr.write(str(total) + '\n')

  fptr.close()