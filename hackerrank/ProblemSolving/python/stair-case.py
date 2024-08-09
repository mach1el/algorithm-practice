#!/bin/python3

import os
import sys

def staircase(n):
  x=1
  while x<n+1:
    print('{0}'.format(('#'*x).rjust(n)))
    x+=1

if __name__ == '__main__':
  n = int(input())
  staircase(n)