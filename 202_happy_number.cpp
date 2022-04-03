
#include <iostream>
#include <string>
#include <set>

// Write an algorithm to determine if a number n is happy.

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

bool isHappy_my(int n) {
    std::set <int64_t> set_number;
    int64_t number = n;
    while(true) {
        if (number != 1) {
            if (set_number.count(number)) {
                return false;
            } else {
                set_number.insert(number);
                std::string number_str = std::to_string(number);
                number = 0;
                for(char cifra : number_str) {
                    number += ((cifra - '0') * (cifra - '0'));
                }
            }
        } else {
            return true;
        }
    }
}

int main() {
    std::cout << isHappy_my(2) << std::endl;
}