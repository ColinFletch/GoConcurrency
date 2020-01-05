package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for n := range jobs {
		fmt.Println("worker", id, "started job", n)
		results <- fib(n)
		fmt.Println("worker", id, "finished job", n)
	}
}

//recursive fib
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	jobs := make(chan int, 100)    // jobs come in
	results := make(chan int, 100) // results go out

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results) // each one will put in work
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 5; j++ {
		fmt.Println(<-results)
	}
}
