package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

// 使用 defer 函数，调用 Recover 来完成 Panic 中断程序导致退出
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
