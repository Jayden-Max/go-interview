package main

import "fmt"

func main() {
	ch1Toch2 := make(chan byte, 1)
	ch2Toch1 := make(chan byte, 1)

	// quit := make(chan int,1)
	go func() {
		for {
			char := <-ch2Toch1
			if char > byte('z') {
				close(ch2Toch1)
				break
			}

			fmt.Println("goroutine 1:\t", string(char))

			ch1Toch2 <- byte(char) + 1
		}
	}()

	go func() {
		for {
			char := <-ch1Toch2
			if char > byte('z') {
				close(ch2Toch1)
				break
			}
			fmt.Println("goroutine 2:\t", string(char))
			ch2Toch1 <- byte(char) + 1
		}

	}()

	ch2Toch1 <- 'a'

	for {
		switch {

		}
	}
}
