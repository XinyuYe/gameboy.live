package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{1, 1, 1, 2, 2, 2}

	c := make(chan int)
	go sum(s[:len(s)/3], c)
	go sum(s[len(s)/3*2:], c)
	go sum(s[len(s)/3:len(s)/3*2], c)
	x := <-c
	y := <-c
	// y := <-c
	// y, x := <-c, <-c // receive from c

	fmt.Println(x, y)
}
