## code001. 
- 题目
```
设计两个goroutine，一个打印a,c,e...，另一个打印b,d,f...
输出结构为a,b,c,d,e...z
```
- 实现（对应code001.go, code002.go)
```
# 修改一下输出即可
package main

import (
	"fmt"
)

func main() {
	ch1Toch2 := make(chan byte, 1)
	ch2Toch1 := make(chan byte, 1)
	quit := make(chan bool)

	go func() {
		for {
			if char, ok := <-ch2Toch1; ok {
				fmt.Println("goroutine 1:\t", string(char))
				if char+1 > byte('z') {
					quit <- true
					break
				}
				ch1Toch2 <- byte(char) + 1
			}
		}
	}()

	go func() {
		for {
			if char, ok := <-ch1Toch2; ok {
				fmt.Println("goroutine 2:\t", string(char))
				if char+1 > byte('z') {
					quit <- true
					break
				}
				ch2Toch1 <- byte(char) + 1
			}
		}
	}()

	ch2Toch1 <- 'a'
	<-quit
	fmt.Println("print finish")
}

```
- 结果截图
![code001](../img/code001.png)


### code001的变种
- 题目
```
使用两个 goroutine 交替打印序列，
一个 goroutinue 打印数字， 另外一个goroutine打印字母， 最终效果如下:

$ go run main.go

12AB34CD56EF78GH910IJ
```
- 解答
```
package main

import "fmt"

func main() {
	numCh := make(chan bool, 1)
	charCh := make(chan bool, 1)
	done := make(chan bool, 1)

	go func() {
		for i := 1; i < 11; i += 2 {
			<-charCh
			fmt.Println("num:\t", i, i+1)
			numCh <- true
		}
	}()

	go func() {
		for i := 'A'; i < 'K'; i += 2 {
			<-numCh
			fmt.Println("char:\t", string(i), string(i+1))
			charCh <- true
		}
		done <- true
	}()

	charCh <- true
	<-done
}

```