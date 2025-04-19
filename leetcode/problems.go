package leetcode

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type stackInt struct {
	stack []int
	lock  sync.RWMutex
}

func (stack *stackInt) Push(value int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.stack = append(stack.stack, value)
}

func (stack *stackInt) Pop() (int, error) {
	if stack.Size() > 0 {
		stack.lock.Lock()
		defer stack.lock.Unlock()
		value := stack.stack[stack.Size()-1]
		stack.stack = stack.stack[:stack.Size()-1]
		return value, nil
	}
	return 0, fmt.Errorf("Size stack = 0")
}

func (stack *stackInt) Empty() bool {
	return stack.Size() == 0
}

func (stack *stackInt) Front() (int, error) {
	if stack.Size() > 0 {
		stack.lock.Lock()
		defer stack.lock.Unlock()
		value := stack.stack[stack.Size()-1]
		return value, nil
	}
	return 0, fmt.Errorf("Size stack = 0")
}

func (stack *stackInt) Size() int {
	return len(stack.stack)
}

func newStack() stackInt {
	return stackInt{
		stack: make([]int, 0),
	}
}

func TwoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for index, value := range nums {
		i, ok := result[target-value]
		if ok != false {
			return []int{i, index}
		}
		result[value] = index
	}
	return []int{0, 0}
}

func IsPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	for index := 0; index < len(strX)/2; index++ {
		if strX[index] != strX[len(strX)-index-1] {
			return false
		}
	}
	return true
}

func IsPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}

	reverseX := 0
	copyX := x

	// сравниваем первую цифру и последнюю
	for copyX > 0 {
		reverseX = reverseX*10 + copyX%10
		copyX = copyX / 10
	}

	return reverseX == x
}
func RomanToInt(s string) int {
	b := map[int32]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	result := 0
	previosValue := 10000000

	for _, value := range s {
		currenValue := b[value]
		if currenValue > previosValue {
			result -= previosValue
			result += (currenValue - previosValue)
		} else {
			result += currenValue
		}

		previosValue = currenValue
	}

	return result
}

func LongestCommonPrefix(strs []string) string {
	commonPrefix := strs[0]
	for _, str := range strs[1:] {
		sizePrefix := 0
		if len(str) > len(commonPrefix) {
			sizePrefix = len(commonPrefix)
		} else {
			sizePrefix = len(str)
			commonPrefix = commonPrefix[:sizePrefix]
		}

		for index := range str[:sizePrefix] {
			if commonPrefix[index] != str[index] {
				commonPrefix = str[:index]
				break
			}
		}

	}

	return commonPrefix
}

func ValidParentheses(str string) bool {

	reverse := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := make([]int32, 0, 10)

	for _, skobka := range str {
		if skobka == '(' || skobka == '[' || skobka == '{' {
			stack = append(stack, skobka)
		} else if skobka == ')' || skobka == ']' || skobka == '}' {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if reverse[last] != skobka {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func ValidParenthesesV2(str string) bool {

	fmt.Println("test")
	customStack := newStack()

	for _, value := range str {
		switch value {
		case '{':
			customStack.Push(int(value))
			break
		case '[':
			customStack.Push(int(value))
			break
		case '(':
			customStack.Push(int(value))
			break
		case '}':
			valueStack, err := customStack.Pop()
			if err != nil || valueStack != '{' {
				return false
			}
			break
		case ']':
			valueStack, err := customStack.Pop()
			if err != nil || valueStack != '[' {
				return false
			}
			break
		case ')':
			valueStack, err := customStack.Pop()
			if err != nil || valueStack != '(' {
				return false
			}
			break
		default:
			return false
		}
	}

	return customStack.Empty()
}

func RemoveDuplicatess_26(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[k] {
			k++
			nums[k] = nums[i]
		} else {
			nums[k] = nums[i]
		}
	}

	return k + 1
}

func RemoveElement(nums []int, val int) int {
	len := len(nums)
	count := 0

	for i := 0; i < len; i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}

	return len - (len - count)
}

func strStrSample(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func StrStr(haystack string, needle string) int {
	// O(N)
	prefix := make([]int, len(needle))
	j := 0
	for i := 1; i < len(needle); {
		if needle[i] == needle[j] {
			prefix[i] = j + 1
			i++
			j++
		} else {
			if j == 0 {
				prefix[i] = 0
				i++
			} else {
				j = prefix[j-1]
			}
		}
	}

	k := 0
	l := 0
	for k < len(haystack) {
		if haystack[k] == needle[l] {
			k++
			l++
			if l == len(needle) {
				return k - l
			}
		} else {
			if l == 0 {
				k++
			} else {
				l = prefix[l-1]
			}
		}
	}

	return -1
}

func SearchInsertPosition(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	// В этом месте left будет индексом, на который нужно вставить target
	return left
}

func LenghtOfLastWord(s string) int {
	result := 0

	for index := len(s) - 1; index >= 0; index-- {
		if s[index] == 32 {
			if result != 0 {
				break
			}
		} else {
			result++
		}
	}
	return result
}
func PlusOne(digits []int) []int {
	result := make([]int, len(digits)+1)
	value := 1
	for i := len(digits) - 1; i >= 0; i-- {

		value += digits[i]

		if value >= 10 {
			result[i+1] = value - 10
			value = 1
		} else {
			result[i+1] = value
			value = 0
		}
	}

	if value > 0 {
		result[0] = value
	} else {
		result = result[1:]
	}
	return result
}

// ----------minus----------//
func AddBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	indexB := len(b) - 1
	result := make([]byte, len(a))

	var shifter, sum byte
	for i := len(a) - 1; i >= 0; i-- {
		sum = shifter + a[i]
		if indexB >= 0 {
			sum += b[indexB]
			indexB--
		}

		result[i] = sum%2 + '0'
		shifter = sum >> 1 % 2
	}
	if shifter == 0 {
		return string(result)
	}

	return "1" + string(result)
}

func mySqrtBinary(x int) int {
	/*
	   Define the search space:
	       Min answer we can get is 0.
	       Max answer is x + 1 in case x = 0 or x = 1.
	*/
	left, right := 0, x+1

	for left < right {
		mid := left + (right-left)/2

		/*
		   If we overshoot, move the right pointer to the left.
		   Otherwise, move the left pointer to the right.
		*/
		if mid*mid > x {
			right = mid
		} else {
			left = mid + 1
		}
	}

	/*
	   At the end of the loop, the left pointer will be placed at ceil(n) such that n^2 = x.

	   For example, x = 8 (sqrt is 2.82842...) and the left pointer will be at 3.
	   Return left - 1 = 3 - 1 = 2

	   So we need to return (left - 1) as we are asked to round the answer down to the neares integer.
	*/
	return left - 1
}

func MySqrt(x int) int {
	if x == 0 {
		return 0
	}

	var i int
	for i = 1; i <= x; i++ {
		if i*i > x {
			return i - 1
		} else if i*i == x {
			return i
		}
	}
	return 0
}

func ClimbingsStairsReqursive(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return ClimbingsStairsReqursive(n-1) + ClimbingsStairsReqursive(n-2)
}

func ClimbingsStairsIter(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	result_one := 2
	result_two := 1

	result := result_two + result_one

	for i := 3; i < n; i++ {
		result_two = result_one
		result_one = result
		result = result_two + result_one

	}

	return result
}

func MergeSortedArray_88(nums1 []int, m int, nums2 []int, n int) {

	i := m - 1 // nums1 iter
	j := n - 1 // nums2 iter
	for k := m + n - 1; k >= 0; k-- {
		if j >= 0 && i >= 0 {
			if nums1[i] > nums2[j] {
				nums1[k] = nums1[i]
				i--
			} else {
				nums1[k] = nums2[j]
				j--
			}
		} else {
			if j < 0 {
				nums1[k] = nums1[i]
				i--
			} else {
				nums1[k] = nums2[j]
				j--
			}
		}
	}
}

// TODO: реализовать самостоятельно
func ReverseParentheses(s string) string {
	var stack [][]byte
	var top []byte
	stack = append(stack, []byte{}) // Initial segment in case we have any before the first parentheses
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, []byte{})
			continue
		case ')':
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			rev(top)
			stack[len(stack)-1] = append(stack[len(stack)-1], top...)
			continue
		default:
			stack[len(stack)-1] = append(stack[len(stack)-1], s[i])

		}
	}
	return string(stack[0])
}

func rev(seg []byte) {
	i, j := 0, len(seg)-1
	for i < j {
		seg[i], seg[j] = seg[j], seg[i]
		i++
		j--
	}
}

func GenerageTrianglePascal(numsRows int) [][]int {

	arrayLine := make([][]int, 0)

	for i := 0; i < numsRows; i++ {
		if i == 0 {
			arrayLine = append(arrayLine, []int{1})
			continue
		}

		line := make([]int, 0)

		line = append(line, 1)

		for j := 1; j < i; j++ {

			first := arrayLine[i-1][j-1]

			second := arrayLine[i-1][j]

			line = append(line, first+second)
		}

		line = append(line, 1)

		arrayLine = append(arrayLine, line)
	}

	return arrayLine

}

func GetRowTriangelPascalReqursive(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	prevRow := GetRowTriangelPascalReqursive(rowIndex - 1)

	currentRow := make([]int, len(prevRow)+1)

	currentRow[0] = 1

	for i := 1; i < len(prevRow); i++ {
		currentRow[i] = prevRow[i-1] + prevRow[i]
	}

	currentRow[len(prevRow)] = 1

	return currentRow

}

func MaxProfixBruteForce_121(prices []int) int {
	var profit = 0
	var minPrice = prices[0]

	for i := 1; i < len(prices); i++ {
		// If we find any price which is lower than the current minPrice
		// update the minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			// If diff of current stock with minPrice is greater
			// update the profit
			profit = prices[i] - minPrice
		}
	}

	return profit

}

func isLetterOrNumber(s byte) (bool, byte) {
	if s >= 65 && s <= 90 {
		return true, s + 32
	}

	if s >= 97 && s <= 122 {
		return true, s
	}

	if s >= 48 && s <= 57 {
		return true, s
	}

	return false, 0
}

func IsPalindromeString_125(s string) bool {
	begin := 0
	end := len(s) - 1
	result := true

	for begin <= end {
		begin_bool, begin_letter := isLetterOrNumber(s[begin])
		end_bool, end_letter := isLetterOrNumber(s[end])

		if begin_bool == end_bool {
			if begin_letter != end_letter {
				return false
			}
			end--
			begin++
			continue
		}

		if !begin_bool {
			begin++
			continue
		}

		if !end_bool {
			end--
			continue
		}

	}

	return result

}

func SingleNumber_136(nums []int) int {

	numsCount := make(map[int]int)

	for _, value := range nums {
		numsCount[value]++
	}

	for k, v := range numsCount {
		if v == 1 {
			return k
		}
	}

	return 0
}

func SingleNumberXOR_136(nums []int) int {
	unique := 0
	for _, num := range nums {
		unique ^= num
	}
	return unique
}

func SingleNumberSort_136(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i += 2 {
		if i+1 >= len(nums) {
			return nums[i]
		}

		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return 0
}

func ConvertToTitle_168(columnNumber int) string {
	s := make([]byte, 0)

	for columnNumber > 0 {
		columnNumber--
		ostatok := columnNumber % 26
		s = append(s, byte(ostatok+65))
		columnNumber = columnNumber / 26
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}

func MajorityElement_169(nums []int) int {
	numsCount := make(map[int]int, len(nums))

	for _, value := range nums {
		numsCount[value]++
	}

	result := 0
	maxCount := 0

	for num, count := range numsCount {
		if count > maxCount {
			maxCount = count
			result = num
		}
	}

	return result
}

func MajorityElement_VoteAlgorithm_169(nums []int) int {
	majority_element, majority_element_freq := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if majority_element_freq == 0 {
			majority_element, majority_element_freq = nums[i], 1
		} else {
			if majority_element == nums[i] {
				majority_element_freq++
			} else {
				majority_element_freq--
			}
		}
	}

	return majority_element
}

func TitleToNumber(columnTitle string) int {
	summa := 0

	for _, letter := range columnTitle {
		summa = summa * 26
		summa += int(letter) - 64

	}

	return summa
}

func IsIsomorphic_205(s string, t string) bool {
	mapS := make([]int, 128)
	mapT := make([]int, 128)

	for index := range s {
		sCh := s[index]
		tCh := t[index]

		if mapS[sCh] == 0 && mapT[tCh] == 0 {
			mapS[sCh] = int(tCh)
			mapT[tCh] = int(sCh)
		} else if mapS[sCh] != int(tCh) || mapT[tCh] != int(sCh) {
			return false
		}
	}
	return true
}

func ContainsDuplicate_217(nums []int) bool {
	countNums := make(map[int]int, len(nums))

	for _, value := range nums {
		if _, ok := countNums[value]; ok {
			return true
		}

		countNums[value]++
	}

	return false
}

func absoluteInt(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func ContainsNearbyDuplicate_219(nums []int, k int) bool {
	countNums := make(map[int]int, len(nums))

	for index, value := range nums {
		if prevIndex, ok := countNums[value]; ok {
			if absoluteInt(index-prevIndex) <= k {
				return true
			}
		}

		countNums[value] = index
	}
	return false
}

func IsAnagram(s string, t string) bool {
	countLetter := make(map[byte]int, len(s))
	for _, letter := range s {
		countLetter[byte(letter)]++
	}

	for _, letter := range t {
		if _, ok := countLetter[byte(letter)]; !ok {
			return false
		}
		countLetter[byte(letter)]--
	}

	for _, value := range countLetter {
		if value != 0 {
			return false
		}
	}
	return true
}

func AddDigits(num int) int {
	result := num

	for result > 9 {
		summa := 0

		for result != 0 {
			summa += result % 10
			result = result / 10
		}
		result = summa
	}

	return result
}

func CanNimWim_292(n int) bool {
	return n%4 != 0
}

func DetectCapitalUse_520(word string) bool {
	firstLetter := false
	allLetter := false

	for i, s := range word {
		if i == 0 && s >= 65 && s <= 90 {
			firstLetter = true
			continue
		}

		if s >= 65 && s <= 90 {
			if i == 1 {
				if firstLetter {
					allLetter = true
				} else {
					return false
				}
			} else if !allLetter {
				return false
			} else {
				continue
			}
		} else if allLetter {
			return false
		}
	}

	return true

	// other solution
	//allCap := strings.ToUpper(word)
	//allLower := strings.ToLower(word)
	//firstCap := strings.ToUpper(string(word[0])) + strings.ToLower(word)[1:]

	//return word == allCap || word == allLower || word == firstCap
}

func RemoveDuplicates_80(nums []int) int {
	k := 1
	twiceFlag := false

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
			twiceFlag = false
		} else if !twiceFlag {
			nums[k] = nums[i]
			twiceFlag = true
			k++
		}
	}

	return k
}

func Rotate_189(nums []int, k int) {

	if k == len(nums) || len(nums) < 2 || k == 0 {
		return
	}

	if k > len(nums) {
		k = k % len(nums)
	}

	if k%2 == 0 {
		Rotate_189(nums, k-1)
		Rotate_189(nums, 1)
		return
	}

	positionNew := k
	tmpOld := nums[0]
	tmpNew := 0

	for i := 0; i < len(nums); i++ {
		if positionNew >= len(nums) {
			positionNew = positionNew - len(nums)
		}

		tmpNew = nums[positionNew]
		nums[positionNew] = tmpOld
		tmpOld = tmpNew
		positionNew += k
	}

}

func maxProfit_122(prices []int) int {
	profit := 0
	prevPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i]-prevPrice > 0 {
			profit += prices[i] - prevPrice
			prevPrice = prices[i]
		} else if prevPrice > prices[i] {
			prevPrice = prices[i]
		}
	}

	return profit
}

func IsSubsequence_392(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	si := 0
	for ti := 0; ti < len(t); ti++ {
		if s[si] == t[ti] {
			si++
			if si == len(s) {
				return true
			}
		}
	}
	return false

}

func CanConstruct_383(ransomNote string, magazine string) bool {

	counter := make(map[byte]int, len(magazine))
	for _, s := range magazine {
		counter[byte(s)]++
	}

	for _, s := range ransomNote {
		if count, ok := counter[byte(s)]; !ok || count == 0 {
			return false
		}
		counter[byte(s)]--
	}

	return true
}

func WordPattern_290(pattern string, s string) bool {
	patternS := strings.Split(s, " ")
	mapPattern := make(map[byte]string)
	mapPatternStr := make(map[string]byte)

	if len(patternS) != len(pattern) {
		return false
	}

	for i, sym := range pattern {
		checkOne, okOne := mapPattern[byte(sym)]
		checkTwo, okTwo := mapPatternStr[patternS[i]]

		if okOne != okTwo {
			return false
		} else if !okOne {
			mapPattern[byte(sym)] = patternS[i]
			mapPatternStr[patternS[i]] = byte(sym)
		} else if checkOne != patternS[i] || checkTwo != byte(sym) {
			return false
		}
	}
	return true
}

func SummaryRanges_228(nums []int) []string {
	result := make([]string, 0, len(nums))

	if len(nums) == 0 {
		return result
	}

	if len(nums) == 1 {
		return append(result, strconv.Itoa(nums[0]))
	}

	i, j := 1, 0
	for ; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			if i-j == 1 {
				result = append(result, strconv.Itoa(nums[j]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[j], nums[i-1]))
			}
			j = i
		}
	}

	if i-j == 1 {
		result = append(result, strconv.Itoa(nums[i-1]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[j], nums[i-1]))
	}

	return result
}

func HammingWeight_191(n int) int {
	bynaryStr := fmt.Sprintf("%b", n)
	count := 0

	for _, value := range bynaryStr {
		if value == '1' {
			count++
		}
	}

	return count
}

func currentArea(begin, end, long int) int {
	var min int

	if begin > end {
		min = end
	} else {
		min = begin
	}

	return min * long
}

func MaxArea(height []int) int {
	begin := 0
	end := len(height) - 1
	max := currentArea(height[begin], height[end], end-begin)

	for begin != end {
		if height[begin] > height[end] {
			end--
		} else {
			begin++
		}

		if max < currentArea(height[begin], height[end], end-begin) {
			max = currentArea(height[begin], height[end], end-begin)
		}
	}

	return max

}

func GameOfLife(board [][]int) {
	checkLive := func(x, y int) {
		countLiveSquare := 0

		if x-1 >= 0 && board[x-1][y]%2 == 1 {
			countLiveSquare++
		}

		if x-1 >= 0 && y-1 >= 0 && board[x-1][y-1]%2 == 1 {
			countLiveSquare++
		}

		if y-1 >= 0 && board[x][y-1]%2 == 1 {
			countLiveSquare++
		}

		if y-1 >= 0 && x+1 < len(board) && board[x+1][y-1]%2 == 1 {
			countLiveSquare++
		}

		if x+1 < len(board) && board[x+1][y]%2 == 1 {
			countLiveSquare++
		}

		if x+1 < len(board) && y+1 < len(board[0]) && board[x+1][y+1]%2 == 1 {
			countLiveSquare++
		}

		if y+1 < len(board[0]) && board[x][y+1]%2 == 1 {
			countLiveSquare++
		}

		if x-1 >= 0 && y+1 < len(board[0]) && board[x-1][y+1]%2 == 1 {
			countLiveSquare++
		}

		if board[x][y] == 1 {
			if countLiveSquare == 2 || countLiveSquare == 3 {
				board[x][y] = 3
			}
		} else {
			if countLiveSquare == 3 {
				board[x][y] = 2
			}
		}
	}

	for indexLine, arrayLine := range board {
		for indexColumn := range arrayLine {
			checkLive(indexLine, indexColumn)
		}
	}

	for indexLine, arrayLine := range board {
		for indexColumn := range arrayLine {
			if board[indexLine][indexColumn] > 1 {
				board[indexLine][indexColumn] = 1
			} else {
				board[indexLine][indexColumn] = 0
			}
		}
	}
}

func isValidSudoku(board [][]byte) bool {
	checkRow := func(row []byte) bool {
		array := make(map[byte]bool, 9)
		for i := byte('1'); i <= byte('9'); i++ {
			array[i] = false
		}

		for _, value := range row {
			if value == '.' {
				continue
			}

			if v, ok := array[value]; !ok || v {
				return false
			}

			array[value] = true
		}

		return true
	}

	checkColumn := func(board [][]byte, column int) bool {
		array := make(map[byte]bool, 9)
		for i := byte('1'); i <= byte('9'); i++ {
			array[i] = false
		}

		for i := 0; i < len(board); i++ {
			if board[i][column] == '.' {
				continue
			}

			if v, ok := array[board[i][column]]; !ok || v {
				return false
			}

			array[board[i][column]] = true

		}

		return true
	}

	checkBox := func(board [][]byte) bool {
		for boxRow := 0; boxRow < 3; boxRow++ {
			for boxCol := 0; boxCol < 3; boxCol++ {
				// Карта для проверки уникальности цифр в подблоке 3x3
				digits := make(map[byte]bool)
				for i := 0; i < 3; i++ {
					for j := 0; j < 3; j++ {

						num := board[boxRow*3+i][boxCol*3+j]
						if num == '.' {
							continue
						}

						// Если цифра уже есть в карте, значит есть повторение
						if digits[num] {
							return false
						}

						// Записываем цифру как встреченную
						digits[num] = true
					}
				}
			}
		}

		return true
	}

	for i, row := range board {
		if !checkRow(row) {
			return false
		}

		if !checkColumn(board, i) {
			return false
		}
	}

	return checkBox(board)
}

func spiralOrderHelper(matrix [][]int, resultOrder []int, count int) []int {
	first, second, third, four := false, false, false, false

	for i := 0 + count; i < len(matrix[0])-count; i++ {
		first = true
		resultOrder = append(resultOrder, matrix[count][i])
	}

	if first {
		for i := count + 1; i < len(matrix)-count; i++ {
			second = true
			resultOrder = append(resultOrder, matrix[i][len(matrix[i])-1-count])
		}
	}

	if second {
		for i := len(matrix[0]) - count - 2; i >= 0+count; i-- {
			third = true
			resultOrder = append(resultOrder, matrix[len(matrix)-1-count][i])
		}
	}

	if third {
		for i := len(matrix) - 2 - count; i >= 1+count; i-- {
			four = true
			resultOrder = append(resultOrder, matrix[i][count])
		}
	}

	if four {
		resultOrder = spiralOrderHelper(matrix, resultOrder, count+1)
	}

	return resultOrder
}

func spiralOrder(matrix [][]int) []int {
	resultOrder := make([]int, 0)

	resultOrder = spiralOrderHelper(matrix, resultOrder, 0)

	return resultOrder
}

func rotate_48(matrix [][]int) {
	// Транспонирование
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Отражение по вертикальной оси
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func canJump_55(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		} else if nums[i]+i > max {
			max = nums[i] + i
		}
	}

	return true
}

type kv struct {
	key int
	val int
}

func frequencySort_1636(nums []int) []int {
	freqArr := make(map[int]int)

	for _, value := range nums {
		freqArr[value]++
	}

	var sortedSlice []kv

	for k, v := range freqArr {
		sortedSlice = append(sortedSlice, kv{key: k, val: v})
	}

	sort.Slice(sortedSlice, func(i, j int) bool {
		if sortedSlice[i].val == sortedSlice[j].val {
			// When values are equal, sort by key in descending order
			return sortedSlice[i].key > sortedSlice[j].key
		}

		return sortedSlice[i].val < sortedSlice[j].val
	})

	i := 0

	for j := 0; j < len(sortedSlice); j++ {
		for sortedSlice[j].val > 0 {
			nums[i] = sortedSlice[j].key
			i++
			sortedSlice[j].val--
		}
	}

	return nums
}

func countSeniors(details []string) int {
	result := 0
	for _, val := range details {
		if val[11] > '6' || (val[11] == '6' && val[12] != '0') {
			result++
		}
	}

	return result
}

func twoOutOfThree_2032(nums1 []int, nums2 []int, nums3 []int) []int {
	result := make([]int, 0)
	resultKey := make(map[int]int)

	for _, value := range nums1 {
		resultKey[value] = 1
	}

	for _, value := range nums2 {
		if key, ok := resultKey[value]; ok && key == 1 {
			resultKey[value] = 10
		} else if key != 10 {
			resultKey[value] = 2
		}
	}

	for _, value := range nums3 {
		if key, ok := resultKey[value]; ok && (key == 1 || key == 2) {
			resultKey[value] = 10
		}
	}

	for key, value := range resultKey {
		if value == 10 {
			result = append(result, key)
		}
	}

	return result
}

func uncommonFromSentences(s1 string, s2 string) []string {
	result := make([]string, 0)

	firstWords := strings.Split(s1, " ")
	secondWords := strings.Split(s2, " ")

	mapWords := make(map[string]int)

	for _, value := range firstWords {
		if key, ok := mapWords[value]; !ok {
			mapWords[value] = 1
		} else if key == 1 {
			mapWords[value] = 0
		}
	}

	for _, value := range secondWords {
		if key, ok := mapWords[value]; !ok {
			mapWords[value] = 2
		} else if key == 1 || key == 2 {
			mapWords[value] = 0
		}
	}

	for key, value := range mapWords {
		if value == 1 || value == 2 {
			result = append(result, key)
		}
	}

	return result
}

func sortedSquares_977(nums []int) []int {
	result := make([]int, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]*nums[i] < nums[j]*nums[j]
	})
	for i := range nums {
		result[i] = nums[i] * nums[i]
	}

	return result
}

func getFinalState_3264(nums []int, k int, multiplier int) []int {
	findMinimal := func(nums []int) int {
		indexMin := 0
		valueMin := nums[indexMin]

		for i, v := range nums {
			if v < valueMin {
				indexMin = i
				valueMin = v
			}
		}
		return indexMin
	}

	for k > 0 {
		k--
		indexMin := findMinimal(nums)
		nums[indexMin] *= multiplier
	}

	return nums
}

func FindDuplicate_287(nums []int) int {
	first := nums[nums[0]]
	second := nums[nums[nums[0]]]

	for first != second {
		first = nums[first]
		second = nums[nums[second]]
	}

	first = nums[0]
	for first != second {
		first = nums[first]
		second = nums[second]
	}

	return first
}

func MissingNumber_268(nums []int) int {
	lenNums := len(nums) + 1
	summaNums := lenNums * (lenNums - 1) / 2
	summaNumsOrig := 0
	for _, value := range nums {
		summaNumsOrig += value
	}

	return summaNums - summaNumsOrig
}

func findErrorNums_645(nums []int) []int {
	freq := make([]int, len(nums)+1)
	for _, value := range nums {
		freq[value] += 1
	}

	var twiceValue, noValue int
	for index, value := range freq[1:] {
		if value == 0 {
			noValue = index + 1
		}

		if value == 2 {
			twiceValue = index + 1
		}
	}

	return []int{twiceValue, noValue}

}

func moveZeroes_283(nums []int) {

	valueIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[valueIndex] = nums[i]
			valueIndex++
		}
	}

	for i := valueIndex; i < len(nums); i++ {
		nums[i] = 0
	}

}

func increasingTriplet(nums []int) bool {
	second := math.MaxInt32
	first := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		if first >= nums[i] {
			first = nums[i]
		} else if second >= nums[i] {
			second = nums[i]
		} else {
			return true
		}
	}

	return false
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	multiplication := 1
	countZero := 0
	tmp := 1

	for _, v := range nums {
		if v == 0 {
			countZero++
		} else {
			tmp *= v
		}
		multiplication *= v
	}

	for i, v := range nums {
		if v != 0 {
			result[i] = multiplication / v
		} else if v == 0 && countZero < 2 {
			result[i] = tmp
		} else {
			result[i] = 0
		}
	}

	return result
}

func mergeAlternately_1768(word1 string, word2 string) string {
	resutl := make([]byte, len(word1)+len(word2))

	i, j, k := 0, 0, 0

	for i < len(word1) && j < len(word2) {
		resutl[k] = word1[i]
		k++
		i++
		resutl[k] = word2[j]
		k++
		j++
	}

	for i < len(word1) {
		resutl[k] = word1[i]
		k++
		i++
	}

	for j < len(word2) {
		resutl[k] = word2[j]
		k++
		j++
	}
	return string(resutl)
}

func largesAltitude(gain []int) int {
	max := 0
	tmp := max

	for _, value := range gain {
		tmp += value
		if tmp > max {
			max = tmp
		}
	}

	return max
}

func pivotIndex(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}

	summa_left := 0
	for i := 0; i < len(nums); i++ {
		summa_right := sum - summa_left - nums[i]
		if summa_right == summa_left {
			return i
		}
		summa_left += nums[i]
	}

	return -1
}

func equalPairs(grid [][]int) int {

	count := 0

	getSumCol := func(grid [][]int, arr []int) int {
		count := 0
		for j := 0; j < len(grid); j++ {
			flag := true
			for i := 0; i < len(grid); i++ {
				if grid[i][j] != arr[i] {
					flag = false
					break
				}
			}
			if flag {
				count += 1
			}
		}
		return count
	}

	for i := 0; i < len(grid); i++ {
		count += getSumCol(grid, grid[i])
	}

	return count
}

func decodeString(s string) string {
	newString := make([]byte, 0)
	myAtoi := func(s string) (int, int) {
		number, count := 0, 0
		for i := 0; s[i] < '['; i++ {
			count++
		}
		number, err := strconv.Atoi(s[:count])
		if err != nil {
			fmt.Println(err)
		}
		return number, count
	}

	mySep := func(s string) string {
		mstr := make([]byte, 0)
		for i := 0; s[i] != ']'; i++ {
			mstr = append(mstr, s[i])
		}
		return string(mstr)
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			number, count := myAtoi(s[i:])
			mstr := mySep(s[i+count:])
			for number > 0 {
				newString = append(newString, []byte(mstr)...)
				number--
			}
			i += count + len(mstr) + 1
		default:
			newString = append(newString, s[i])
		}
	}
	return string(newString)
}

func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	a := len(str1)
	b := len(str2)

	if a < b {
		a, b = b, a
	}

	for b > 0 {
		a, b = b, a%b
	}

	return str1[:a]
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
