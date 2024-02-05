#!/usr/bin/env python

def solution(ele,count):
  if ele != []:
    for i in range(1,count+1):
      last_ele = ele[-1]
      ele.pop(-1)
      ele = [last_ele]+ele
  return(ele)