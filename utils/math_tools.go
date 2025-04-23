package utils

// 数学工具类
type MathUtil struct{}

func (m *MathUtil) Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m *MathUtil) Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m *MathUtil) GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 使用示例
//mathUtil := MathUtil{}
//fmt.Println(mathUtil.GCD(48, 18)) // 输出 6
