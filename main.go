package main

import "fmt"

func main() {

	//lv0
	fmt.Println("Hello 康桥")

	//lv1
	const pi = 3.14
	r := 5.0
	fmt.Println(pi * r * r)

	//lv2
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)

	//lv3
	var num3 int
	fmt.Print("请输入一个数字: ")
	_, err := fmt.Scan(&num3)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}
	result := factorial(num3)
	fmt.Printf("%d的阶乘是%d\n", num3, result)

	//lvx
	var num4 int
	sum2, count := 0, 0
	for {
		fmt.Println("请输入一个整数(输入0结束)")
		_, err := fmt.Scan(&num4)
		if err != nil {
			fmt.Println("输入错误:", err)
			break
		}
		if num4 == 0 {
			break
		}
		sum2 += num4
		count++
	}

	if average := average(sum2, count); average >= 60 {
		fmt.Printf("平均成绩为%2f ，成绩合格", average)
	} else {
		fmt.Printf("平均成绩为%2f ，成绩不合格", average)
	}
}

func factorial(n int) int {
	factorial := 1
	if n < 0 {
		fmt.Println("n不能为负数")
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	for ; n > 1; n-- {
		factorial *= n
	}
	return factorial
}

func average(sum int, count int) float64 {
	return float64(sum) / float64(count)
}
