/* You are given a 0-indexed string s consisting of only lowercase English letters, where each letter in s appears exactly twice. You are also given a 0-indexed integer array distance of length 26.

Each letter in the alphabet is numbered from 0 to 25 (i.e. 'a' -> 0, 'b' -> 1, 'c' -> 2, ... , 'z' -> 25).

In a well-spaced string, the number of letters between the two occurrences of the ith letter is distance[i]. If the ith letter does not appear in s, then distance[i] can be ignored.

Return true if s is a well-spaced string, otherwise return false. 

Constraints:
    2 <= s.length <= 52
    s consists only of lowercase English letters.
    Each letter appears in s exactly twice.
    distance.length == 26
    0 <= distance[i] <= 50

*/
#include <string>
#include <vector>
#include <unordered_map>
using namespace std;

class Solution {
public:
    bool checkDistances(string s, vector<int>& distance) {
        for (int i = 0; i < s.size() - 1; i++) { // we technically don't need to check the very last letter so we can be blazing fast ;)
            int index = s[i] - 'a';
            if (distance[index] < 0) { // we'll make distance[index] negative once we're done checking each letter
                continue;
            }
            int next_index = i + distance[index] + 1;
            if (next_index < s.size() && s[next_index] == s[i]) {
                distance[index] = -1; // we're done checking this letter
            } else {
                return false;
            }
        }
        return true;
    }
};