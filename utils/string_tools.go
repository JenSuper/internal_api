package utils

// 字符串处理工具
type StringUtil struct{}

// KMP算法实现字符串查找
func (s *StringUtil) KMP(text, pattern string) int {
	// 实现KMP算法
	return -1 // 返回匹配位置或-1
}

// 判断是否为回文
func (s *StringUtil) IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
