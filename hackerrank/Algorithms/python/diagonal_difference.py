#!/bin/python3

import os
import sys

def diagonalDifference(n,a):
	x=sum([a[x][x] for x in range(n)])
	y=sum([a[x][n-1-x] for x in range(n)])	
	return(abs(x-y))

if __name__ == '__main__':
  f = open(os.environ['OUTPUT_PATH'], 'w')

  n = int(input())

  a = []

  for _ in range(n): a.append(list(map(int, input().rstrip().split())))

  result = diagonalDifference(n,a)

  f.write(str(result) + '\n')

  f.close()