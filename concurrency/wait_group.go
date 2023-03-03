package concurrency

import (
	"fmt"
	"sync"
)

func evaluateParams(x int) int {
	fmt.Printf("Evaluating parameter: %v\n", x)
	return x
}

func waitAndPrintParam(x int) {
	fmt.Printf("Started waiting for: %v\n", x)
	fmt.Printf("Completed wait for: %v\n", x)
}

func UseWaitGroup(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			waitAndPrintParam(evaluateParams(i))
		}()
	}
	wg.Wait()
}

func OnlyDefer(n int) {
	for i := 0; i < n; i++ {
		defer waitAndPrintParam(evaluateParams(i))
	}
}
