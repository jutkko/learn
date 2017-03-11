import unittest

def tuplify(line):
    ints = tuple(map(int, line.split(' ')))
    return hash(ints)

class TestTuplify(unittest.TestCase):
    def test12(self):
        self.assertEqual(tuplify("1 2"), 3713081631934410656)
