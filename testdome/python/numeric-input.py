#!/usr/bin/env python

class TextInput:
  def __init__(self):
    self.vals = ""
  def add(self,character):
    self.vals+=character
  def get_value(self):
    return self.vals

class NumericInput(TextInput):
  def __init__(self):
    super().__init__()
  def add(self,character):
    if character.isdigit(): self.vals+=character

if __name__ == '__main__':
  input = NumericInput()
  input.add("1")
  input.add("a")
  input.add("0")
  print(input.get_value())