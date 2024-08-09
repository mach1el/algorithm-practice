#!/bin/python3

import os
import sys

def solve(score1, score2):
  Alice = 0
  Bob = 0
  
  for i in range(len(score1)):
    if int(score1[i]) > int(score2[i]): Alice+=1
    elif int(score1[i]) < int(score2[i]): Bob+=1

  print( Alice,Bob ) 

if __name__ == '__main__':
  f = open(os.environ['OUTPUT_PATH'], 'w')

  score1 = input().split()

  score2 = input().split()

  result = solve(score1, score2)

  f.write(' '.join(map(str, result)))
  f.write('\n')

  f.close()