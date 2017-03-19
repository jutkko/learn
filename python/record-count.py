import unittest

def break_max_min(list):
    if len(list) < 1:
        return [0, 0]
    else:
        result = [0] * 2
        min = list[0]
        max = list[0]
        for n in list:
            if n < min:
                result[1] += 1
                min = n
            if n > max:
                result[0] += 1
                max = n
        return result

class TestTuplify(unittest.TestCase):
    def testExample(self):
        self.assertEqual(break_max_min([3, 4, 21, 36, 10, 28, 35, 5, 24, 42]), [4, 0])
