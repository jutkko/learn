import unittest

# x1 != x2
def kangaroo(x1, v1, x2, v2):
    if v1 == v2:
        print("NO")
    elif (x2 - x1) % (v1 - v2) != 0:
        return "NO"
    elif (x2 - x1) * (v1 - v2) < 0:
        return "NO"
    else:
        return "YES"

class TestKangaroo(unittest.TestCase):
    def testExample(self):
        self.assertEqual(kangaroo(0, 3, 4, 2), "YES")
    def testExample1(self):
        self.assertEqual(kangaroo(0, 2, 5, 3), "NO")
