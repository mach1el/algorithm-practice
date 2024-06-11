#!/bin/python3

import math
import os
import random
import re
import sys
from datetime import datetime


def dayOfProgrammer(year):
  if year in range(1700,1917):
    if year%4==0: return f'12.09.{year}'
    else: return f'13.09.{year}'
  
  elif year==1918: return '26.09.1918'

  else:
    if (year%400==0) or (year%4==0 and not year%100==0): return f'12.09.{year}'
    else: return f'13.09.{year}'

if __name__ == '__main__':
  fptr = open(os.environ['OUTPUT_PATH'], 'w')

  year = int(input().strip())

  result = dayOfProgrammer(year)

  fptr.write(result + '\n')

  fptr.close()