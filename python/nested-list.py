import unittest

def me(students):
    students.sort(key = lambda x: x[1])

    secondindex = 0
    for s in students:
        if s[1] != students[0][1]:
            break
        secondindex += 1

    results = [s[0] for s in students if s[1] == students[secondindex][1]]

    results.sort()
    return results

class TestMe(unittest.TestCase):
    def testExample(self):
        self.assertEqual(me([['Harry', 37.21], ['Berry', 37.21], ['Tina', 37.2], ['Akriti', 41], ['Harsh', 39]]), ["Berry", "Harry"])
    def testMadeUp(self):
        self.assertEqual(me([['Harry', 37.21], ['Berry', 37.2], ['Tina', 37.2], ['Akriti', 41], ['Harsh', 39]]), ["Harry"])
