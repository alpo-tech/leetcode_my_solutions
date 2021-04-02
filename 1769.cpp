/*
The idea here is to use use Prefix sum array.
We iterate the given array twice, but in opposite direction. In 1st pass we move from left to right, storing the total balls to be moved from 0th index to current ith index (ie. all elements to left), in a variable (here sum) and res stores the operations needed to move all balls to ith index. Thus, res[i]= res[i-1]+ sum, gives minimum number of operations needed to move all balls from the left indices.

Similarily, we need to iterate from right to left and storing total balls in a variable (sum) and using it to find res[i], which is res[i]=res[i-1]+sum.

Time Complexity : O(N) due to the linear traversal from 1st element to last element & vice versa
Space Complexity: O(N) due to the space used to store the prefix sum and ultimately the result.


vector<int> minOperations(string boxes) {
        int n=boxes.length();
        vector<int>res(n);
        int last_cost=0, sum=boxes[0]-'0';
        for(int i=1;i<n;i++){
            res[i]+=sum+last_cost;
            last_cost=res[i];
            sum+=(boxes[i]-'0');
        }

        last_cost=0,sum=boxes[n-1]-'0';

        for(int i=n-2;i>=0;i--){
            res[i]+=sum+last_cost;
            last_cost=(sum+last_cost);
            sum+=(boxes[i]-'0');
        }
        return res;
    }
*/

class Solution {
public:
    vector<int> minOperations(string boxes) {
         vector<int> result;
         vector<int> tmp;
         for(char& letter : boxes) {
               result.push_back(letter - '0');
         }
         tmp = result;
         for(int i = 0; i < static_cast<int>(result.size()); ++i) {
             unsigned long summa = 0;
             for(int j = 0; j < static_cast<int>(result.size()); ++j) {
                 summa += tmp[j] * abs(j - i);
             }
             result[i] = summa;
         }
         return result;
    }
};