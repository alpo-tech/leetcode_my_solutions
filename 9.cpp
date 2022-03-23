// Given an integer x, return true if x is palindrome integer.

// An integer is a palindrome when it reads the same backward as forward.

// For example, 121 is a palindrome while 123 is not.

bool isPalindrome(int x) {
    if (x < 0) { 
        return false;
    }
    int64_t answer = 0;
    int origin = x;
    while (origin) {
        answer = answer * 10 + origin % 10;
        origin /= 10;
    }
    return answer == x;
}