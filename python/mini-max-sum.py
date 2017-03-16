import unittest

def mini_sum(list):
    return sum(sorted(list)[:-1])

def max_sum(list):
    return sum(sorted(list)[1:])

class TestMiniMaxSum(unittest.TestCase):
    def testExample(self):
        self.assertEqual(mini_sum([1, 2, 3, 4, 5]), 10)
    def testExample1(self):
        self.assertEqual(max_sum([1, 2, 3, 4, 5]), 14)
