import unittest

def me(students):
    sortedStudents = sorted(students, key = lambda x: x[1])

    secondindex = 0
    for s in sortedStudents:
        if s[1] != sortedStudents[0][1]:
            break
        secondindex += 1

    results = []
    for s in sortedStudents[secondindex:]:
        if s[1] == sortedStudents[secondindex][1]:
            results.append(s[0])
        else:
            break

    return sorted(results)

if __name__ == '__main__':
    for _ in range(int(input())):
        name = input()
        score = float(input())

# class TestMe(unittest.TestCase):
#     def testExample(self):
#         self.assertEqual(me([['Harry', 37.21], ['Berry', 37.21], ['Tina', 37.2], ['Akriti', 41], ['Harsh', 39]]), ["Berry", "Harry"])
#     def testMadeUp(self):
#         self.assertEqual(me([['Harry', 37.21], ['Berry', 37.2], ['Tina', 37.2], ['Akriti', 41], ['Harsh', 39]]), ["Harry"])
