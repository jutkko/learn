import unittest

def concat_string(n):
    result = ""
    for num in range(1, n+1):
        result += str(num)

    return result

class TestPrintSequence(unittest.TestCase):
    def test10(self):
        self.assertEqual("12345678910", concat_string(10))
    def test5(self):
        self.assertEqual("12345", concat_string(5))
    def test1(self):
        self.assertEqual("1", concat_string(1))
