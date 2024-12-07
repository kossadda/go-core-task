package main

import (
	"fmt"
	"github.com/kossadda/go-core-task/8/pkg/waitgroup"
	"time"
)

func main() {
	wg := waitgroup.New()

	tasks := 3

	wg.Add(int64(tasks))

	for i := 1; i <= tasks; i++ {
		go func(taskID int) {
			defer wg.Done()
			fmt.Printf("Task %d started\n", taskID)
			time.Sleep(time.Duration(taskID) * time.Second)
			fmt.Printf("Task %d finished\n", taskID)
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks completed")
}
