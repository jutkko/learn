import unittest
import math

def gcd(list):
    result = list[0]
    for n in list:
        result = math.gcd(n, result)
    return result

def lcm(list):
    result = list[0]
    for n in list:
        result = lcm_helper(n, result)
    return result

def lcm_helper(a, b):
    return a * b // math.gcd(a, b)

def between(x, y):
    count = 0
    m = x
    while m <= y:
        if y % m == 0:
            count += 1
        m += x

    return count

class TestBetweenTwoSets(unittest.TestCase):
    def testExample(self):
        self.assertEqual(between(4, 16), 3)
    def testExample1(self):
        self.assertEqual(between(2, 16), 4)
    def testExample2(self):
        self.assertEqual(lcm([2, 4]), 4)
    def testExample3(self):
        self.assertEqual(gcd([16, 32, 64]), 16)
