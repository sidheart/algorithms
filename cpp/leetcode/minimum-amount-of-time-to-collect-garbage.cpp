/* You are given a 0-indexed array of strings garbage where garbage[i] represents the assortment of garbage at the ith house. garbage[i] consists only of the characters 'M', 'P' and 'G' representing one unit of metal, paper and glass garbage respectively. Picking up one unit of any type of garbage takes 1 minute.

You are also given a 0-indexed integer array travel where travel[i] is the number of minutes needed to go from house i to house i + 1.

There are three garbage trucks in the city, each responsible for picking up one type of garbage. Each garbage truck starts at house 0 and must visit each house in order; however, they do not need to visit every house.

Only one garbage truck may be used at any given moment. While one truck is driving or picking up garbage, the other two trucks cannot do anything.

Return the minimum number of minutes needed to pick up all the garbage. */
class Solution {
public:
    int garbageCollection(vector<string>& garbage, vector<int>& travel) {
        int time = 0;
        unordered_map<char, int> garbage_to_last_house; // furthest house containing each garbage type
        garbage_to_last_house['M'] = -1;
        garbage_to_last_house['P'] = -1;
        garbage_to_last_house['G'] = -1;
        for (int i = 0; i < garbage.size(); i++) {
            time += garbage[i].size();
            for (char g : garbage[i]) {
                garbage_to_last_house[g] = i;
            }
        }
        // convert travel vector into total travel time
        for (int i = 1; i < travel.size(); i++) {
            travel[i] += travel[i-1];
        }
        // add each truck's travel time
        if (garbage_to_last_house['M'] > 0) {
            time += travel[garbage_to_last_house['M']-1];
        }
        if (garbage_to_last_house['P'] > 0) {
            time += travel[garbage_to_last_house['P']-1];
        }
        if (garbage_to_last_house['G'] > 0) {
            time += travel[garbage_to_last_house['G']-1];
        }
        return time;
    }
};

