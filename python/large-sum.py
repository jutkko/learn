import unittest
from functools import reduce

def sum_array(list):
    return sum(list)

class TestLargeSum(unittest.TestCase):
    def testExample(self):
        self.assertEqual(sum_array([1000000001, 1000000002, 1000000003, 1000000004, 1000000005]), 5000000015)
