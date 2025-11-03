package main

import "testing"

func TestControl(t *testing.T) {
	test := []struct {
		input    []int
		expected map[int]int
	}{
		{
			[]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5},
			map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 1},
		},
		{
			[]int{},
			map[int]int{},
		},
	}

	for i, tt := range test {
		result := Control(tt.input)

		for k, v := range tt.expected {
			if result[k] != v {
				t.Errorf("测试用例%d: 键%d的值不匹配，期望%d，得到%d", i, k, v, result[k])
			}
		}

	}
}

func TestCalculationAdd(t *testing.T) {
	test := []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 2, 1},
	}

	for _, tt := range test {
		result := Calculation("add")(tt.a, tt.b)

		if result != tt.want {
			t.Errorf("加法 %d + %d = %d, 期望 %d", tt.a, tt.b, result, tt.want)
		}
	}
}

//其他的差不多
