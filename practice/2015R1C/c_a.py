from codejam import CodeJam, parsers
import json

def solve(c, d, v, denominations):
  newDen = 0
  for i in range(1, v):
    ok, closest = isPossible(i, c, denominations)
    if not ok:
      print("Adding " + str(closest))
      denominations.append(closest)
      newDen += 1
  return newDen

def isPossible(i, c, denominations):
  return isPossibleRec(i,
                       0,
                       sum(denominations) * c,
                       list(zip(denominations, [c]*len(denominations))))

def isPossibleRec(i, currentSum, denSum, denTuples):
  #print(json.dumps(denTuples))
  if currentSum == i:
    return True, 0

  if currentSum > i:
    return False, -1

  if len(denTuples) <= 0:
    return False, i - currentSum

  newSum = currentSum + denTuples[-1][0]
  denSum -= denTuples[-1][0]
  denTuples[-1] = (denTuples[-1][0], denTuples[-1][1] - 1)
  if denTuples[-1][1] <= 0:
    newDenTuples = denTuples[:-1]
  else:
    newDenTuples = denTuples

  ok, closest = isPossibleRec(i, newSum, denSum, newDenTuples)
  if ok:
    return True, 0

  ok, closest2 = isPossibleRec(i, currentSum, denSum, newDenTuples)
  if ok:
    return True, 0

  if closest == -1:
    proposal = closest2
  else:
    proposal = min(closest, closest2)
  if proposal in denTuples:
    proposal = -1
  return False, proposal

@parsers.iter_parser
def parse(next):
  c, d, v = next().strip().split(' ')
  denominations = [int(x) for x in next().strip().split(' ')]
  return int(c), int(d), int(v), denominations

if __name__ == "__main__":
  CodeJam(parse, solve).main()
