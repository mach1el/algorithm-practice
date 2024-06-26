#!/usr/bin/evn python3

from operator import xor
from functools import reduce

t = int(input())

for _ in range(t):
  _, bins=input(),list(map(int, input().split()))
  if len(bins)==1: print('Second')
  else:
    win=reduce(xor,(i if bins[i]%2 else 0 for i in range(1,len(bins))))
    print('First' if win else 'Second')