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
Before solving the question we will keep two thing in mind.
1.The use of Vectors
2.The use of Hashing.

Here, this question can be solved by a direct formula which we will use to find the number of good pairs.
The formula is (n*(n-1))/2.

*First we find the count of number occurences of a number.
*Then store the number count in the dictionary .
[1,2,3,1,1,3]

key - value pair
1 : 3
2 : 1
3 : 2

*Then we consider n = value in the hash table
*using loop calculate the good pairs.

//Program to find the number of good pairs
//Good pairs are those in which nums[i] == nums[j] and i<j
class Solution {
public:
    int numIdenticalPairs(vector<int>& nums) 
    {
   
        unordered_map<int,int> umap; //Initializing a Hash Table
         
        for(int i=0;i<nums.size();i++) //Iterating through the vector
        {
            ++umap[nums[i]];  //Counting the occurences of a number and storing it in value.
            
        }
        int good_pairs = 0;
        for(auto i:umap) //Using the formula 
        {
            int n = i.second; //i.second implies -- value of hash table
            good_pairs += ((n)*(n-1))/2;
            
        }
        return good_pairs;
        
        
    }
};
*/