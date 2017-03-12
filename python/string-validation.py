import unittest

def any(line, func):
    for c in line:
        if func(c):
            return True
    return False

class TestStringValidation(unittest.TestCase):
    def testExample(self):
        self.assertEqual(any("qA2", str.isalnum), True)
        self.assertEqual(any("qA2", str.isalpha), True)
        self.assertEqual(any("qA2", str.isdigit), True)
        self.assertEqual(any("qA2", str.islower), True)
        self.assertEqual(any("qA2", str.isupper), True)
    def testExample1(self):
        self.assertEqual(any("qa2", str.isalnum), True)
        self.assertEqual(any("qa2", str.isalpha), True)
        self.assertEqual(any("qa2", str.isdigit), True)
        self.assertEqual(any("qa2", str.islower), True)
        self.assertEqual(any("qa2", str.isupper), False)
