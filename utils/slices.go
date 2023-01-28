package utils

import "sync"

func SumElements(slice []int, out chan<- int, wg *sync.WaitGroup) {
	var sum int
	for _, value := range slice {
		sum += value
	}
	out <- sum
	defer wg.Done()

}
