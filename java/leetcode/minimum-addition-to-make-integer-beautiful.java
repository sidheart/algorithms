/* 
You are given two positive integers n and target.

An integer is considered beautiful if the sum of its digits is less than or equal to target.

Return the minimum non-negative integer x such that n + x is beautiful. The input will be generated such that it is always possible to make n beautiful. 
*/
class Solution {
    private int sumOfDigits(long n) {
        int sum = 0;
        for (long m = n; m > 0; m /= 10) {
            sum += (int) (m % 10);
        }
        return sum;
    }

    public long makeIntegerBeautiful(long n, int target) {
        long answer = 0;
        float power = 0;
        while (sumOfDigits(n) > target) {
            long diff = (long) Math.pow(10, power) * (10 - (n % 10));
            n += 10 - (n % 10);
            n /= 10;
            answer += diff;
            power++;
        }
        return answer;
    }
}