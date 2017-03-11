import unittest

def find_second_largest(list):
    result = -101
    max = -101
    for x in list:
        if max < x:
            result = max
            max = x
        else:
            if result < x and x != max:
                result = x
    return result

class TestSecondLargest(unittest.TestCase):
    def test1(self):
        self.assertEqual(find_second_largest([1, 2, 3, 5, 4]), 4)
    def test2(self):
        self.assertEqual(find_second_largest([5, 5, 5, 5, 5]), -101)
    def test3(self):
        self.assertEqual(find_second_largest([5, 5, 5, -5, 5]), -5)
