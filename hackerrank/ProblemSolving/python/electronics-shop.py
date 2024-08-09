#!/bin/python3

import os
import sys

def getMoneySpent(keyboards, drives, b):

  if len(keyboards) > len(drives): 
    for i in keyboards: prices = [x+i for x in drives]
  else:
    for i in drives: prices = [x+i for x in keyboards]
  
  spent = [ price for price in prices if price <= b]
  
  if len(spent) == 0: return -1
  else: return max(spent)

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  bnm = input().split()

  b = int(bnm[0])

  n = int(bnm[1])

  m = int(bnm[2])

  keyboards = list(map(int, input().rstrip().split()))

  drives = list(map(int, input().rstrip().split()))

  moneySpent = getMoneySpent(keyboards, drives, b)

  fptr.write(str(moneySpent) + '\n')

  fptr.close()