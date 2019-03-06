package main

import "fmt"

// https://gist.github.com/bilxio/4968875f7d93fd07e6521f0afb4899e1
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
