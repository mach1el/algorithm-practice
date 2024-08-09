#!/bin/python3

import os

def appendAndDelete(s, t, k):
  i = 0
  for x in range(min(len(s),len(t))):
    i+=1
    if s[x] != t[x]:
      break

  num=k-len(s)-len(t)+2*i

  if num<0:
    print("No")
  elif s == t:
    print("Yes")
  elif num<=2*i:
    if num%2 != 0:
      print("No")
    elif num%k == 0:
      print("No")
    else:
      print("Yes")
  else:
    print("Yes")

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  s = input()

  t = input()

  k = int(input().strip())

  result = appendAndDelete(s, t, k)

  fptr.write(result + '\n')

  fptr.close()