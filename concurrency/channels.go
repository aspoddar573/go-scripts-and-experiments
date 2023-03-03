package concurrency

import (
	"fmt"
	"sync"
)

func sumWithChannel(wg *sync.WaitGroup, s []int, c chan int) {
	for _, v := range s {
		//fmt.Printf("adding %v\n", v)
		sum := <-c
		c <- v + sum
	}
	wg.Done()
	//c <- sumWithChannel // send sumWithChannel to c
}

func FindSumWithChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 2)
	c <- 0

	var wg sync.WaitGroup
	wg.Add(2)

	go sumWithChannel(&wg, s[:len(s)/2], c)
	go sumWithChannel(&wg, s[len(s)/2:], c)

	wg.Wait()

	sum := <-c
	fmt.Println(sum)
}

func sumWithoutChannel(wg *sync.WaitGroup, s []int, c *int) {
	for _, v := range s {
		//fmt.Printf("adding %v\n", v)
		*c += v
	}
	wg.Done()
	//c <- sum // send sum to c
}

func FindSumWithoutChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	var c int = 0

	var wg sync.WaitGroup
	wg.Add(2)

	go sumWithoutChannel(&wg, s[:len(s)/2], &c)
	go sumWithoutChannel(&wg, s[len(s)/2:], &c)

	wg.Wait()

	sum := c
	fmt.Println(sum)
}
