#!/bin/python3

import os
import sys

def timeConversion(s):
	h,m,se = s.split(':')

	if se.upper().endswith('AM'):
		se = se.replace('AM','')
		if h == "12":
			h = "00"
		
	else:
		se = se.replace('PM','')
		if h != "12":
			h = int(h)+12
	return("{0}:{1}:{2}".format(str(h),m,se))
    

if __name__ == '__main__':
	f = open(os.environ['OUTPUT_PATH'], 'w')

	s = input()

	result = timeConversion(s)

	f.write(result + '\n')

	f.close()
