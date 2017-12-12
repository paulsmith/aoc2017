"""
Advent of Code 2017

http://adventofcode.com/2017/day/1

You're standing in a room with "digitization quarantine" written in LEDs along
one wall. The only door is locked, but it includes a small interface.
"Restricted Area - Strictly No Digitized Users Allowed."

It goes on to explain that you may only leave by solving a captcha to prove
you're not a human. Apparently, you only get one millisecond to solve the
captcha: too fast for a normal human, but it feels like hours to you.

The captcha requires you to review a sequence of digits (your puzzle input) and 
find the sum of all digits that match the next digit in the list. 
The list is circular, so the digit after the last digit is the first digit in
the list.

For example:

* 1122 produces a sum of 3 (1 + 2) because the first digit (1) matches the second
  digit and the third digit (2) matches the fourth digit.
* 1111 produces 4 because each digit (all 1) matches the next.
* 1234 produces 0 because no digit matches the next.
* 91212129 produces 9 because the only digit that matches the next one is the
  last digit, 9.
"""

import sys

def part1():
    input = sys.stdin.readline().strip()
    sum = 0
    l = len(input)
    for i, c in enumerate(input):
        if c == input[(i+1) % l]:
            sum += int(c)
    print(sum)

def part2():
    input = sys.stdin.readline().strip()
    sum = 0
    l = len(input)
    half = l//2
    for i, c in enumerate(input[:half]):
        if c == input[i+half]:
            sum += int(c)*2
    print(sum)

def part2_alt():
    input = sys.stdin.readline().strip()
    sum = 0
    a, b = input[:len(input)//2], input[len(input)//2:]
    for i, c in enumerate(a):
        if c == b[i]:
            sum += int(c)
    for i, c in enumerate(b):
        if c == a[i]:
            sum += int(c)
    print(sum)

{'part1': part1,
 'part2': part2,
 'part2_alt': part2_alt,
}[sys.argv[1]]()