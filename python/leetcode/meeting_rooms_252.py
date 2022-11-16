from typing import List


# See https://leetcode.com/problems/meeting-rooms/
class Solution:
    def canAttendMeetings(self, intervals: List[List[int]]) -> bool:
        sorted_intervals = sorted(intervals, key=lambda interval: interval[1])
        for i in range(0, len(sorted_intervals)-1):
            if sorted_intervals[i][1] > sorted_intervals[i+1][0]:
                return False
        return True