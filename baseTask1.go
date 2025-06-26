package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1.1 只出现一次的数字
	arr := []int{4, 2, 3, 1, 2, 3, 4}
	number := onlyNumber(arr)
	fmt.Println("only number:", number)

	// 1.2 回文数
	palindromeNum := 129
	fmt.Println("palindrome is ture?", isPalindrome(palindromeNum))

	//1.3 有效的括号
	symbol := "{}]}()"
	vaild := isValid(symbol)
	fmt.Println("brackets is valid?:", vaild)

	// 1.4 最长公共前缀
	str := []string{"flower", "flow", "flight"}
	commonPrefix := longestCommonPrefix(str)
	fmt.Println("the longest common prefix:", commonPrefix)

	// 1.5加1
	plus := []int{9, 9}
	fmt.Println("plus one result:：", plusOne(plus))

	// 1.6 移除重复项
	nums := []int{1, 2, 2}
	duplicate := removeDuplicates(nums)
	fmt.Println("del duplicate options result:", duplicate)

	// 1.7 合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge := merge(intervals)
	fmt.Println("merge result:", merge)

	// 1.8 两数相加
	arr2 := []int{2, 7, 11, 15}
	ret := twoSum(arr2, 9)
	fmt.Println("two sum result:", ret)

}

/**
 *	136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
 *  找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
 *  结合 if 条件判断和 map 数据结构来解决
 */
func onlyNumber(arr []int) int {
	m := make(map[int]int)
	for _, val := range arr {
		m[val]++
	}

	for key := range m {
		if m[key] == 1 {
			return key
		}
	}
	return -1
}

// 两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if p, ok := m[target-val]; ok {
			return []int{p, index}
		}
		m[val] = index
	}
	return []int{}
}

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 数字操作、条件判断
// 判断一个整数是否是回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	ret := 0
	for x > ret {
		ret = ret*10 + x%10
		x = x / 10
	}
	return x == ret || x == ret/10
}

// 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}
	m := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	stack := []byte{}
	for i := range length {
		if m[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 字符串处理、循环嵌套
// 查找字符串数组中的最长公共前缀
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < n; i++ {
		prefix = commonPrefix(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func commonPrefix(s1 string, s2 string) string {
	min := min(len(s1), len(s2))
	index := 0
	for index < min && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

// 考察：数组操作、进位处理
// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
// 这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。将大整数加 1，并返回结果的数字数组。
func plusOne(digits []int) []int {
	len := len(digits)
	for i := len - 1; i >= 0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}

// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
// 返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1)
// 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，
// 一个快指针 j 用于遍历数组，
// 当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	fast, slow := 1, 1
	for fast < n {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			merged = append(merged, interval)
		}
	}
	return merged
}
