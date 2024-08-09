#!/bin/python3

import math
import os
import random
import re
import sys

def migratoryBirds(arr):
  count_ids = { value:arr.count(value) for value in list(set(arr)) }
  return max(count_ids, key=lambda k: count_ids.get(k))

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  arr_count = int(input().strip())

  arr = list(map(int, input().rstrip().split()))

  result = migratoryBirds(arr)

  fptr.write(str(result) + '\n')

  fptr.close()