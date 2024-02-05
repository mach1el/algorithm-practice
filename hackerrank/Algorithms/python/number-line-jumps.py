#!/usr/bin/env python

x1,v1,x2,v2 = [int(x) for x in input().split()]
print("NO" if v2 >= v1 or (x2-x1) % (v1-v2) != 0 else "YES")