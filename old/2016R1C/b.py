from codejam import CodeJam, parsers

def solve(c, d, v, denominations):
  return "solution"

@parsers.iter_parser
def parse(next):
  c, d, v = next().strip().split(' ')
  denominations = [int(x) for x in next().strip().split(' ')]
  return int(c), int(d), int(v), denominations

if __name__ == "__main__":
  CodeJam(parse, solve).main()
