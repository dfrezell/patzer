#!/usr/bin/env python

sq = 16

sqs = range(sq*8)

grid = ""
row  = ""

for x in sqs:
    row += "%4d" % (x)
    if x % sq == sq - 1:
        grid = row + "\n" + grid
        row = ""

print grid

grid = ""
row  = ""

for x in sqs:
    row += "(%2x,%2x) " % (x/16, x%16)
    if x % sq == sq - 1:
        grid = row + "\n" + grid
        row = ""

print grid


grid = ""
row  = ""

for x in sqs:
    row += "(%8s) " % ("{0:b}".format(x))
    if x % sq == sq - 1:
        grid = row + "\n" + grid
        row = ""

print grid

