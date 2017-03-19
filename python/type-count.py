import unittest

def count(line):
    maxCount = 0
    result = None
    # there are only 5 types
    for i in range(1, 5 + 1):
        count = line.count(i)
        if count > maxCount:
            maxCount = count
            result = i
    return result

class TestTypeCount(unittest.TestCase):
    def testExample(self):
        self.assertEqual(count([1, 4, 4, 4, 5, 3]), 4)
    def testMadeUp(self):
        self.assertEqual(count([1, 1, 1, 4, 4, 4, 5, 3]), 1)
