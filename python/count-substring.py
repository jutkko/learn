import unittest

def count_substring(string, sub_string):
    result = 0
    index = 0
    while index != -1:
        index = string.find(sub_string, index)
        if index == -1:
            break
        index += 1
        result += 1

    return result

class TestCountSubstring(unittest.TestCase):
    def testExample(self):
        self.assertEqual(count_substring("ABCDCDC", "CDC"), 2)
    def testMadeUp(self):
        self.assertEqual(count_substring("ABCDCDCAB", "AB"), 2)
    def testMadeUp1(self):
        self.assertEqual(count_substring("ABCDCDCAB", "U"), 0)
    def testMadeUp2(self):
        self.assertEqual(count_substring("ABCDCDCAB", "ABCDCDCAB"), 1)
