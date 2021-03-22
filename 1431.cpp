/*
Given the array candies and the integer extraCandies, 
where candies[i] represents the number of candies that the ith kid has.
For each kid check if there is a way to distribute extraCandies among 
the kids such that he or she can have the greatest number of candies among them. 
Notice that multiple kids can have the greatest number of candies.
Input: candies = [4,2,1,1,2], extraCandies = 1
Output: [true,false,false,false,false] 
Explanation: There is only 1 extra candy, 
therefore only kid 1 will have the greatest number of candies among the kids regardless of who takes the extra candy.
*/

#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {
              vector<bool> result;
              int max_candies = *max_element(candies.begin(), candies.end());
              for(const auto& kid : candies) {
                  if(kid + extraCandies >= max_candies) {
                      result.push_back(true);
                  } else {
                      result.push_back(false);
                  }
              }
              return result;
    }
};