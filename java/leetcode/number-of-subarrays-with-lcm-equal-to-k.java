/*
Given an integer array nums and an integer k, return the number of subarrays of nums where the least common multiple of the subarray's elements is k.

A subarray is a contiguous non-empty sequence of elements within an array.

The least common multiple of an array is the smallest positive integer that is divisible by all the array elements.
*/
class Solution {
    public int gcd(int a, int b) {
        int remainder;
        do {
            remainder = a%b;
            a = b;
            b = remainder;
        } while (b != 0);
        return a;
    }
    
    public int lcm(int a, int b) {
        if (a == 0 || b == 0) {
            return 0;
        }
        return (a * b) / gcd(a, b);
    }
    
    public int subarrayLCM(int[] nums, int k) {
        int n = 0;
        for (int start = 0; start < nums.length; start++) {
            int currentLCM = nums[start];
            for (int end = start; end < nums.length; end++) {
                currentLCM = lcm(currentLCM, nums[end]);
                if (currentLCM == k) {
                    n++;
                }
            }
        }
        return n;
    }
}