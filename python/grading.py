import unittest

def grade_student(grade):
    if grade < 38:
        return grade
    if (grade + 2) // 5 != grade // 5:
        return ((grade + 2) // 5) * 5
    return grade

class TestGradeStudent(unittest.TestCase):
    def testExample(self):
        self.assertEqual(grade_student(73), 75)
    def testExample2(self):
        self.assertEqual(grade_student(67), 67)
    def testExample3(self):
        self.assertEqual(grade_student(38), 40)
    def testExample4(self):
        self.assertEqual(grade_student(33), 33)
