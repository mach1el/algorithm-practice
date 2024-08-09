#!/bin/python3

import math
import os
import random
import re
import sys

all_squares_flat = [
  [2, 9, 4, 7, 5, 3, 6, 1, 8],
  [6, 1, 8, 7, 5, 3, 2, 9, 4],
  [6, 7, 2, 1, 5, 9, 8, 3, 4],
  [8, 3, 4, 1, 5, 9, 6, 7, 2],
  [8, 1, 6, 3, 5, 7, 4, 9, 2],
  [4, 9, 2, 3, 5, 7, 8, 1, 6],
  [4, 3, 8, 9, 5, 1, 2, 7, 6],
  [2, 7, 6, 9, 5, 1, 4, 3, 8],
]

def formingMagicSquare(s):
  
  def square_cost(s1, s2):
    return sum([abs(x[0] - x[1]) for x in zip(s1, s2)])

  s = s[0] + s[1] + s[2]
  return min([square_cost(s, x) for x in all_squares_flat])

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  s = []

  for _ in range(3): s.append(list(map(int, input().rstrip().split())))

  result = formingMagicSquare(s)

  fptr.write(str(result) + '\n')

  fptr.close()