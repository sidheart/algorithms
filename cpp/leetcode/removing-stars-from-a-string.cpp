/* You are given a string s, which contains stars *.
 abcd****d***eee*f**
 abcd****d***e

In one operation, you can:

    Choose a star in s.
    Remove the closest non-star character to its left, as well as remove the star itself.

Return the string after all stars have been removed.

Note:

    The input will be generated such that the operation is always possible.
    It can be shown that the resulting string will always be unique. */
class Solution {
public:
    string removeStars(string s) {
        int star_count = 0;
        int insert_pos = s.size()-1;
        for (int i = s.size()-1; i >= 0; i--) {
            if (s[i] == '*') {
                star_count++;
            } else if (star_count == 0) {
                s[insert_pos] = s[i];
                insert_pos--;
            } else {
                star_count--;
            }
        }
        s.erase(0, insert_pos+1);
        return s;
    }
}
