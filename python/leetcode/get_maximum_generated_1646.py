class Solution:
    def getMaximumGenerated(self, n: int) -> int:
        if n == 0 or n == 1:
            return n
        maximum = 1
        gen_list = [0, 1]
        for i in range(2, n+1):
            if i % 2 == 0:
                item = gen_list[i//2]
                if item > maximum:
                    maximum = item
                gen_list.append(item)
            else:
                item = gen_list[i//2] + gen_list[(i//2) + 1]
                if item > maximum:
                    maximum = item
                gen_list.append(item)
        print(gen_list)
        return maximum