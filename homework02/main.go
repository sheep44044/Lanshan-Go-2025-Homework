package main

import "fmt"

func main() {
	//1
	slice := []int{1, 1, 2, 3}
	result := Control(slice)
	fmt.Println(result)

	//2
	fmt.Println("10 + 5 =", Calculation("add")(10, 5))
	fmt.Println("10 - 5 =", Calculation("subtract")(10, 5))
	fmt.Println("10 * 5 =", Calculation("multiply")(10, 5))
	fmt.Println("10 / 5 =", Calculation("divide")(10, 5))
}

// 1 函数
func Control(slice []int) map[int]int {
	result := make(map[int]int)
	for _, v := range slice {
		result[v] += 1
	}
	return result
}

// 2 函数
func Calculation(operation string) func(int, int) int {
	switch operation {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "subtract":
		return func(a, b int) int {
			return a - b
		}
	case "multiply":
		return func(a, b int) int {
			return a * b
		}
	case "divide":
		return func(a, b int) int {
			if b != 0 {
				return a / b
			}
			return 0
		}
	default:
		return nil
	}
}
