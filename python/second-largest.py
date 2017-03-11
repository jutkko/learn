import unittest

def find_second_largest(list):
    max = list[0]
    for x in list:
        if x > max:
            max = x

    not_max = None
    for x in list:
        if x != max:
            not_max = x
            break

    if not_max == None:
        return None

    result = None
    for x in list:
        if x >= not_max and x != max:
            result = x
    return result

class TestSecondLargest(unittest.TestCase):
    def test1(self):
        self.assertEqual(find_second_largest([1, 2, 3, 5, 4]), 4)
    def test2(self):
        self.assertEqual(find_second_largest([5, 5, 5, 5, 5]), None)
    def test3(self):
        self.assertEqual(find_second_largest([5, 5, 5, -5, 5]), -5)
