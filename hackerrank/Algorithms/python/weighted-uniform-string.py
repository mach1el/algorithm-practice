#!/usr/bin/env python3

s = input().strip()
n = int(input().strip())
seq=[]

for num in range(1,27): seq.append(0)
for x in sorted(set(s)):
  i = 1;
  while x * i in s:
    i += 1
  seq[ord(x)-97]=i-1

finale = set()

for index,every in enumerate(seq):
  for sval in range(every): finale.add((index+1)*(sval+1))

for a0 in range(n):
  x = int(input().strip())
  print("Yes" if x in finale else "No")