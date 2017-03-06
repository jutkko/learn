import unittest

def is_leap(year):
    if year % 400 == 0:
        return True
    elif year % 4 == 0 and year % 100 != 0:
        return True
    else:
        return False

class TestLeapYear(unittest.TestCase):
    def test2004(self):
        self.assertTrue(is_leap(2004))
    def test2000(self):
        self.assertTrue(is_leap(2000))
    def test1800(self):
        self.assertFalse(is_leap(1800))
    def test1996(self):
        self.assertTrue(is_leap(1996))
