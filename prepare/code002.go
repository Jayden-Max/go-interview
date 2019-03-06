package main

import "fmt"

// https://github.com/stong1994/interview/blob/master/Golang-2019-2-20/6-print-a-z/main.go
func main() {
	min := int('a')
	max := int('z')
	endFlag := make(chan struct{})
	oddFlag := make(chan int)
	evenFlag := make(chan int)

	end := func(i int) bool {
		if i > max {
			endFlag <- struct{}{}
			return true
		}
		return false
	}

	// 奇数
	odd := func(i int) {
		if !end(i) {
			fmt.Println("goroutine 1:\t", string(i))
			evenFlag <- i + 1
		}
	}

	// 偶数
	even := func(i int) {
		if !end(i) {
			fmt.Println("goroutine 2:\t", string(i))
			oddFlag <- i + 1
		}
	}

	// 分配任务
	go func() {
		for {
			select {
			case o := <-oddFlag:
				go odd(o)
			case e := <-evenFlag:
				go even(e)
			}
		}
	}()

	// 入口
	oddFlag <- min

	// <- endFlag
	for {
		select {
		case <-endFlag:
			return
		}
	}
}
