package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("ddd")
	task := make([]string,10)
	task = append(task, "ddddd")
	task = append(task, "ddddd2")
	task = append(task, "dddddr")
	Process1(task)
	Process2(task)

}

func Process1(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func() {
			fmt.Printf("Worker start process task1: %s\n", task)
		}()
	}
}

func Process2(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func(t string) {
			fmt.Printf("Worker start process task2: %s\n", t)
		}(task)
	}
}

func Double(a int) int {
	return a * 2
}

func TestDouble(t *testing.T) {
	var tests = []struct {
		name         string
		input        int
		expectOutput int
	}{
		{
			name:         "double 1 should got 2",
			input:        1,
			expectOutput: 2,
		},
		{
			name:         "double 2 should got 4",
			input:        2,
			expectOutput: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.expectOutput != Double(test.input) {
				t.Fatalf("expect: %d, but got: %d", test.input, test.expectOutput)
			}
		})
	}
}