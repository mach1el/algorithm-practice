#!/bin/python3

import os

def pageCount(n, p):
  if n == 2: return 0
  if n-p == 1 and n%2 != 0: return 0
  elif n-p == 1 and n%2 == 0: return 1
  else: return(min(p//2,(n-p)//2))
    
if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  n = int(input().strip())

  p = int(input().strip())

  result = pageCount(n, p)
  print(result)

  fptr.write(str(result) + '\n')

  fptr.close()