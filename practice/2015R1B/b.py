from codejam import CodeJam, parsers
import json
import itertools

class record:
  def __init__(self):
    self.possible = {}
    self.maxSerie = 0

def addAll(r, currentSum, c, denominations):
  if len(denominations) <= 0:
    r.possible[currentSum] = 1
    if currentSum == r.maxSerie + 1:
      r.maxSerie = currentSum
    return

  head, *tail = denominations

  for k in range(0, c+1):
    addAll(r, currentSum + k*head, c, tail)

def solve(c, d, v, denominations):
  newDen = 0

  r = record()
  # Calculate all possible
  addAll(r, 0, c, denominations)
  print(r.possible)
  print(r.maxSerie)
  return newDen

@parsers.iter_parser
def parse(next):
  c, d, v = next().strip().split(' ')
  denominations = [int(x) for x in next().strip().split(' ')]
  return int(c), int(d), int(v), denominations

if __name__ == "__main__":
  CodeJam(parse, solve).main()
