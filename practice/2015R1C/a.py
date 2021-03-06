from codejam import CodeJam, parsers

def solve(*lines):
  return sum((sum(line) for line in lines)) # This is where you put your solution

@parsers.iter_parser
def parse(next):
  n = int(next())
  return [int(next()) for unused in range(n)]

if __name__ == "__main__":
  CodeJam(parse, solve).main()
