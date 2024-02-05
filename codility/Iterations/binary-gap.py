#!/usr/bin/env python

def solution(N):
  bin_var = format(N,'b')
  return len(max([i for i in bin_var.strip("0").split("1")]))