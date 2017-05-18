/*
 * 并发的求1~n的和
 */
package main

import "os"
import "fmt"
import "strconv"

func sum(a, b int, done chan int) {
	s := 0
	for i := a; i <= b; i++ {
		s += i
	}
	done <- s
}
func add(done chan int) int {
	s := 0
	a := <- done
	s += a
	b := <- done
	s += b
	return s
}
func main() {
	n, _ := strconv.Atoi(os.Args[1])
	done := make(chan int, 2)
	go sum(1, n / 2, done)
	go sum(n / 2 + 1, n, done)
	fmt.Println(add(done))
}
