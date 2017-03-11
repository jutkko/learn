import unittest

# the preconditio here is position must be in range
def mutate_string(line, position, char):
    chars = list(line)
    chars[position] = char
    return "".join(chars)

class TestMutateString(unittest.TestCase):
    def testExample(self):
        self.assertEqual(mutate_string("abracadabra", 5, 'k'), "abrackdabra")
