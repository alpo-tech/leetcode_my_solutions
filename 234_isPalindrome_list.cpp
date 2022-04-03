//Given the head of a singly linked list, return true if it is a palindrome.

#include <iostream>
#include <vector>

//Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};


class Solution {
public:
    bool isPalindrome(ListNode* head) {
        ListNode *slow = head;
        ListNode *fast = head;
        ListNode *prev;
        ListNode *temp;

        while(fast && fast->next) {
            slow = slow->next;
            fast = fast->next->next;
        }
        prev = slow;
        slow = slow->next;
        prev->next = nullptr;
        while(slow) {
            temp = slow->next;
            slow->next = prev;
            prev = slow;
            slow = temp;
        }
        fast = head;
        slow = prev;
        while(slow) {
            if (slow->val != fast->val) {
                return false;
            } else {
                slow = slow->next;
                fast = fast->next;
            }
        }
        return true;
    }

    bool isPalindrome_bruteforce(ListNode* head) {
        std::vector<int> vector;
        while(head) {
            vector.push_back(head->val);
            head = head->next;
        }
        for(size_t i = 0; i < vector.size() / 2; ++i) {
            if (vector[i] != vector[vector.size() - i - 1]) {
                return false;
            }
        }
        return true;
    }

};



int main() {
    ListNode l1(1);
    ListNode l2(1);
    ListNode l3(2);
    ListNode l4(1);
    ListNode l5(1);


    l1.next = &l2;
    l2.next = &l3;
    l3.next = &l4;
    l4.next = &l5;

    Solution solution;
    std::cout << solution.isPalindrome(&l1) << std::endl;
}