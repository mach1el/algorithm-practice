#!/bin/python3

import math
import os
import random
import re
import sys

def countingValleys(steps, path):
  count = 0
  valley = 0
  for x in path:
    if x == "U": 
      count+=1
      if count==0: valley+=1
    else:
      count-=1
  return valley

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  steps = int(input().strip())

  path = input()

  result = countingValleys(steps, path)

  print(result)

  fptr.write(str(result) + '\n')

  fptr.close()