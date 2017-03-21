import unittest
import math

def encrypt(line):
    dimensions = dimension(len(line))
    result = ""
    for i in range(dimensions[1]):
        for j in range(dimensions[0]):
            index = j * dimensions[1] + i
            if index < len(line):
                result += line[index]

        if i != dimensions[1] - 1:
            result += " "

    return result

def dimension(length):
    column = int(math.ceil(math.sqrt(length)))
    row = int(math.floor(math.sqrt(length)))
    if column * row < length:
        return [row + 1, column]
    else:
        return [row, column]

class TestEncryption(unittest.TestCase):
    def testExample(self):
        self.assertEqual(encrypt("haveaniceday"), "hae and via ecy")
    def testDimension(self):
        self.assertEqual(dimension(80), [9, 9])
    def testDimension1(self):
        self.assertEqual(dimension(71), [8, 9])
    def testDimension2(self):
        self.assertEqual(dimension(54), [7, 8])
    def testDimension3(self):
        self.assertEqual(dimension(2), [1, 2])
    def testDimension4(self):
        self.assertEqual(dimension(5), [2, 3])
