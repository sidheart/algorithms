import math


# https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/
from typing import List


class Solution:
    def divsum(self, nums: List[int], divisor: int) -> int:
        sum = 0
        for num in nums:
            sum += math.ceil(num / divisor)
        return sum

    def smallestDivisor(self, nums: List[int], threshold: int) -> int:
        maximum = None
        for num in nums:
            if maximum is None or num > maximum:
                maximum = num
        ans = maximum
        lo, hi = 1, maximum
        while lo <= hi:
            mid = (lo + hi) // 2
            if self.divsum(nums, mid) <= threshold:
                ans = mid
                hi = mid - 1
            else:
                lo = mid + 1
        return ans
