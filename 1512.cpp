/*Given an array of integers nums.

A pair (i,j) is called good if nums[i] == nums[j] and i < j.

Return the number of good pairs.

 

Example 1:

Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
*/

class Solution {
public:
    int numIdenticalPairs(vector<int>& nums) {
        int result = 0;
        for(size_t i = 0; i < nums.size(); ++i) {
            for (size_t j = i + 1; j < nums.size(); ++j) {
                if (nums[i] == nums[j]) {
                    ++result;
                }
            }
        }
       return result; 
    }
};

/*
class Solution {
public:
    int numIdenticalPairs(vector<int> nums) {
        unordered_map<int,int> frequency;
        int good_pair=0;
        for(auto num : nums){
            if(frequency.find(num) != frequency.end()){ // if we have already seen this element such that i<j
                good_pair+=frequency[num]; // add the number of time u have seen it
            }
            frequency[num]++;
        }
        return good_pair;
    }
};
*/