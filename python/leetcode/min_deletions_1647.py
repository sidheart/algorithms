class Solution:
    def minDeletions(self, s: str) -> int:
        mydict = {}
        for char in s:
            if char in mydict:
                mydict[char] += 1
            else:
                mydict[char] = 1
        same_freq_dict = {}
        for k,v in mydict.items():
            if v in same_freq_dict:
                same_freq_dict[v] += 1
            else:
                same_freq_dict[v] = 0
        reloop = True
        ans = 0
        print(f"same_freq_dict is {same_freq_dict}")
        while reloop:
            reloop = False
            for freq, num_same in same_freq_dict.items():
                if num_same > 0:
                    ans += 1
                    same_freq_dict[freq] -= 1
                    if freq-1 > 0:
                        if freq-1 in same_freq_dict:
                            same_freq_dict[freq-1] += 1
                        else:
                            same_freq_dict[freq-1] = 0
                    reloop = True
                    break
        return ans