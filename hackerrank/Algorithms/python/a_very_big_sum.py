#!/bin/python3

import os
import sys


def aVeryBigSum(n, ar):
	z = 1000000000
	for x in ar:
		z+=x
	return z - 1000000000

if __name__ == '__main__':
  f = open(os.environ['OUTPUT_PATH'], 'w')

  n = int(input())

  ar = list(map(int, input().rstrip().split()))

  result = aVeryBigSum(n, ar)

  f.write(str(result) + '\n')

  f.close()