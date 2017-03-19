import unittest
import math

def sort_probability(list):
    total = math.factorial(len(list))
    counts = find_counts(list)
    
    good = 1
    for c in counts:
        good *= math.factorial(c)
    return float(good)/float(total)

# returns the count for the elems in the list
# sorted order
def find_counts(list):
    uniqueList = set(list)
    result = []
    for u in uniqueList:
        count = list.count(u)
        if count > 1:
            result.append(count)
    return result

def is_sorted(list):
    if len(list) > 0:
        curr = list[0]
        for e in list:
            if curr <= e:
                curr = e
                continue
            else:
                return False
    return True

class TestTypeCount(unittest.TestCase):
    def testExample(self):
        self.assertEqual(sort_probability([1, 2]), 0.500000)
    def testExample1(self):
        self.assertEqual(sort_probability([5, 2, 19, 15, 19, 23, 6, 11, 2, 18, 6, 30, 28, 5, 12]), 1.2235461970911706e-11)
    def testMadeUp(self):
        self.assertEqual(find_counts([4, 1, 1, 1, 4, 4, 4, 5, 3]), [3, 4])
    def testAll(self):
        self.assertEqual("%.6f" % (1/sort_probability([5, 2, 19, 15, 19, 23, 6, 11, 2, 18, 6, 30, 28, 5, 12])), "81729648000.000000")
    def testIsSorted(self):
        self.assertEqual(is_sorted([1, 2, 1, 3]), False)
    def testIsSorted1(self):
        self.assertEqual(is_sorted([1, 2, 3]), True)
