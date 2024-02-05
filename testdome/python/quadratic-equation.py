#!/usr/bin/env python

from math import sqrt

def find_roots(a, b, c):
  x = sqrt(b**2 - 4*a*c)
  if x > 0:
    x1 = (-b + x) / (2*a)
    x2 = (-b - x) / (2*a)
    return (x1,x2)
  elif x == 0:
    x1 = x2 = (-b + x) / (2*a)
    return(x1,x2)

print(find_roots(2, 10, 8));