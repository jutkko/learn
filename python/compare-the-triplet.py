import unittest

def compare(listA, listB):
    zipped = zip(listA, listB)

    result = [0, 0]
    for pair in zipped:
        if pair[0] > pair[1]:
            result[0] += 1
        elif pair[0] < pair[1]:
            result[1] += 1

    return result

class TestTuplify(unittest.TestCase):
    def test12(self):
        self.assertEqual(compare([5, 6, 7], [3, 6, 10]), [1, 1])
