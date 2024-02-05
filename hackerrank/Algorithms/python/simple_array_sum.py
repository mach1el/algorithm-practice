#!/bin/python3

import os
import sys

def simpleArraySum(ar):
	x=0
	for n in ar:
		x += n
	return(x)

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  ar_count = int(input())

  ar = list(map(int, input().rstrip().split()))

  result = simpleArraySum(ar)

  fptr.write(str(result) + '\n')

  fptr.close()