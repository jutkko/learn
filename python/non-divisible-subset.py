import unittest

def get_max_non_divisible_subset(myList, k):
    remainderList = list(map(lambda x: x % k, myList))

    remainderCount = {}
    for r in remainderList:
        if r in remainderCount:
            remainderCount[r] += 1
        else:
            remainderCount[r] = 1
    
    # consider even k, when the number of k / 2 elems is 1
    count = 0
    
    if k % 2 == 0 and k / 2 in remainderCount:
        count += 1
        del remainderCount[k / 2]
        
    if 0 in remainderCount:
        count += 1
        del remainderCount[0]
        
    count += get_sum_except_for_value(remainderCount, k)

    return count
    
def get_sum_except_for_value(myMap, k):
    result = 0
    toExclude = []
    for key, val in myMap.items():
        if k - key != key and key not in toExclude:
            toExclude.append(key)
            toExclude.append(k - key)
            otherValue = 0
            if k - key in myMap:
                otherValue = myMap[k - key]
            result += max(val, otherValue)

    return result

class TestTypeCount(unittest.TestCase):
    def testExample(self):
        self.assertEqual(get_max_non_divisible_subset([1, 7, 2, 4], 3), 3)
    def testExample1(self):
        self.assertEqual(get_max_non_divisible_subset([1], 3), 1)
    def testExample2(self):
        self.assertEqual(get_max_non_divisible_subset([1, 1, 1, 2, 2, 3, 4, 4, 4, 5, 5, 5], 6), 7)
    def testExample3(self):
        self.assertEqual(get_max_non_divisible_subset([422346306, 940894801, 696810740, 862741861, 85835055, 313720373], 9), 5)
    def testExample4(self):
        self.assertEqual(get_max_non_divisible_subset([1, 2, 3, 4, 5], 1), 1)
