#!/bin/python3

import math
import os
import random
import re
import sys

def divisibleSumPairs(n, k, ar):
  count = 0
  for index, value in enumerate(ar):
    list_value = list(map(lambda x:value+x, ar[index+1:]))
    count += len([num for num in list_value if num%k==0])
  return count

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  first_multiple_input = input().rstrip().split()

  n = int(first_multiple_input[0])

  k = int(first_multiple_input[1])

  ar = list(map(int, input().rstrip().split()))

  result = divisibleSumPairs(n, k, ar)

  fptr.write(str(result) + '\n')

  fptr.close()