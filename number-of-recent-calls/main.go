package main

import "fmt"

func main() {
	obj := Constructor()

	input := []int{1, 100, 3001, 3002}

	for _, value := range input {
		requests := obj.Ping(value)
		fmt.Println(requests)
	}
}

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	// Add the current timestamp to the queue
	this.requests = append(this.requests, t)

	// Remove timestamps that are outside the 3000 milliseconds window
	for this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
	}

	// Return the number of requests within the window
	return len(this.requests)
}
