package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Worker Pool")
	start := time.Now()
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 1; i <= 9; i++ {
		go consumer(i, jobs, results)
	}

	go producer(jobs, 10)
	for i := 1; i <= 10; i++ {
		res := <-results
		fmt.Println("Hasil ke", res)
	}

	fmt.Println("==========================")
	elapsed := time.Since(start)
	fmt.Println("Total waktu :", elapsed)
}

func FakeHTTPRequest(x int) int {
	return x * 10
}

// perbedaan jobs chan <- int dengan jobs <- chan int adalah
//
// jobs chan <- int		=> memasukkan nilai kedalam jobs
// jobs <- chan int		=> menerima channel

func producer(jobs chan<- int, size int) {
	for i := 1; i <= size; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("Consumer ke ", id, "Mulai")
		time.Sleep(time.Second * 2)

		results <- FakeHTTPRequest(job)
	}
}
