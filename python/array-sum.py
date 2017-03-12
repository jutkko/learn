import unittest

def array_sum(list):
    result = 0
    for i in list:
        result += i
    return result

class TestArraySum(unittest.TestCase):
    def testExample(self):
        self.assertEqual(array_sum([1, 2, 3, 4, 10, 11]), 31)
