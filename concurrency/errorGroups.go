package concurrency

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
)

func Task(task int) error {
	if rand.Intn(10) == task {
		return fmt.Errorf("Task %v failed\n", task)
	}
	fmt.Printf("Task %v completed\n", task)
	return nil
}
func ErrGroupExample() {
	eg := &errgroup.Group{}
	for i := 0; i < 10; i++ {
		task := i
		eg.Go(func() error {
			return Task(task)
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal("Error", err)
	}
	fmt.Println("Completed successfully!")
}
