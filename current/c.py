from codejam import CodeJam, parsers

def solve(line):
  j = int(line[0])
  p = int(line[1])
  s = int(line[2])
  k = int(line[3])

  max1 = j*p*s
  #print("max1:" + str(max1))
  max2 = j*p*k
  #print("max2:" + str(max2))
  realMax = min(max1, max2)

  output = str(realMax) + "\n"

  cj = 0
  cp = 0
  cs = 0
  ck = 0
  if k < s:
    biais = 0
    for i in range(0, realMax):
      output += "%d %d %d\n" % (cj + 1, ((cp + biais) % p) + 1, cs + 1)
      cs = (cs + 1) % s
      ck += 1
      if ck % k == 0:
        cs = int(ck / k) % s
        cp += 1
        if cp >= p:
          cp = 0
          cj = (cj + 1) % j
        if cs == 0:
          biais += 1
    return output

  for i in range(0, realMax):
    output += "%d %d %d\n" % (cj + 1, cp + 1, cs + 1)
    cs += 1
    if cs >= s:
      cs = 0
      cp += 1
      if cp >= p:
        cp = 0
        cj = (cj + 1) % j

  return output

if __name__ == "__main__":
  CodeJam(parsers.ints, solve).main()
