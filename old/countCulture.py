from codejam import CodeJam, parsers

def solve(line):
  n = int(line[0])
  counter = 1
  current = "1"

  while int(current) < n:
    inverted = str(current[::-1])
    if int(inverted) <= n and int(inverted) > int(current):
      current = inverted
      #print("Choosing inverted: " + current)
    else:
      current = str(int(current) + 1)
      #print("Choosing current + 1: " + current)
    counter += 1

  return str(counter)

#@parsers.iter_parser
#def parse(next):
#  n = next()
#  return int(n)

if __name__ == "__main__":
  CodeJam(parsers.ints, solve).main()
