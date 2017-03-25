import unittest

def find(slice, val):
    return slice.index(val)

class TestFindElem(unittest.TestCase):
    def testExample(self):
        self.assertEqual(find([1, 4, 5, 7, 9, 12], 4), 1)
