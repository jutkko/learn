import unittest

def find_percentage(scoreMap, name):
    sum = 0
    scores = scoreMap[name]
    for score in scores:
        sum += score

    average = sum / len(scores)

    # for formatting to print exactly two digits after decimal point
    return average

class TestFindThePercentage(unittest.TestCase):
    def testExample(self):
        self.assertEqual(find_percentage({'Malika': [52.0, 56.0, 60.0], 'Krishna': [67.0, 68.0, 69.0], 'Arjun': [70.0, 98.0, 63.0]}, "Malika"), 56.00)
    def testExample(self):
        self.assertEqual(find_percentage({'Malika': [52.0, 56.0, 60.0], 'Krishna': [67.0, 68.0, 69.0], 'Arjun': [70.0, 98.0, 63.5]}, "Arjun"), 77.16666666666667)
