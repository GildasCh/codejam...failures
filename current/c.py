from codejam import CodeJam, parsers
import sys, traceback

def getSs(jpList, dj, dp, overall, s, k):
  if len(jpList) == 0:
    return True, []

  head, *tail = jpList
  cj, cp = head

  for a in range(0, s):
    # if not dj.get((cj, a), 0) < k:
    #   print("%d,%d already taken" % (cj, a))
    # elif not dp.get((cp, a), 0) < k:
    #   print("%d,%d already taken" % (cp, a))
    # elif not overall.get((cj, cp, a), 0) < 1:
    #   print("%d,%d,%d already taken" % (cj, cp, a))
    #   print(overall)
    # else:
    if dj.get((cj, a), 0) < k and dp.get((cp, a), 0) < k and overall.get((cj, cp, a), 0) < 1:
      dj[(cj, a)] = dj.get((cj, a), 0) + 1
      dp[(cp, a)] = dp.get((cp, a), 0) + 1
      overall[(cj, cp, a)] = overall.get((cj, cp, a), 0) + 1
      # print("Chose %d" % a)
      ok, ret = getSs(tail, dj, dp, overall, s, k)
      if ok:
        ret.append(a)
        return True, ret
      #print("Falling back %d" % a)
      #print("New values: %d, %d, %d" %
            # (dj.get((cj, a), 1) - 1,
            #  dp.get((cp, a), 1) - 1,
            #  dp.get((cj, cp, a), 1) - 1))
      dj[(cj, a)] = dj.get((cj, a), 1) - 1
      dp[(cp, a)] = dp.get((cp, a), 1) - 1
      overall[(cj, cp, a)] = overall.get((cj, cp, a), 1) - 1

  #print("FAILED!")
  return False, [77]

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

  if k < s:
    overall = {}
    dj = {}
    dp = {}
    jpList = []
    for i in range(0, realMax):
      jpList.append((cj, cp))
      cp += 1
      if cp >= p:
        cp = 0
        cj = (cj + 1) % j

    # print(jpList)
    ok, sList = getSs(jpList, dj, dp, overall, s, k)

    if not ok:
      print("FAILDE")
      print(jpList)
      print(sList)

    for i in range(0, realMax):
      cj, cp = jpList[i]
      cs = sList[i]
      output += "%d %d %d\n" % (cj + 1, cp + 1, cs + 1)

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
