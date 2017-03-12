import unittest

def lonely_integer(list):
    myHash = {}
    for i in list:
        if i in myHash:
            myHash[i] = -1
        else:
            myHash[i] = 1

    for key, value in myHash.items():
        if value == 1:
            return key
    return None

class TestLonelyInteger(unittest.TestCase):
    def testExample(self):
        self.assertEqual(lonely_integer([1]), 1)
    def testExample1(self):
        self.assertEqual(lonely_integer([1, 1, 2]), 2)
    def testExample2(self):
        self.assertEqual(lonely_integer([0, 0, 1, 1, 2]), 2)
