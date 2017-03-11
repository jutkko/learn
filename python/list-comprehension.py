import unittest

def cuboid(x, y, z, n):
    x += 1
    y += 1
    z += 1
    return [[x1, y1, z1] for x1 in range(x) for y1 in range(y) for z1 in range(z) if x1+y1+z1 != n]

class TestTuplify(unittest.TestCase):
    def testCuboid(self):
        self.assertEqual(cuboid(1, 1, 1, 2), [[0, 0, 0], [0, 0, 1], [0, 1, 0], [1, 0, 0], [1, 1, 1]])
