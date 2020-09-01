package queue

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	queue := &Queue{}
	queue.Enqueue(1).
		Enqueue(2).
		Enqueue(3)
	for !queue.Empty() {
		fmt.Print(queue.Dequeue(), " - ")
	}
	fmt.Println()
}

func Print(waitGroup *sync.WaitGroup, c chan string, index int) {
	defer waitGroup.Done()
	fmt.Println("index ", index, " print", <-c)
	time.Sleep(100*time.Millisecond)
}

func TestGoroutine(t *testing.T) {
	bufferChn := make(chan string, 10)
	w := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		w.Add(1)
		go Print(&w, bufferChn, i)
	}
	for i := 0; i < 5; i++ {
		bufferChn <- fmt.Sprint("Task ", i)
	}
	fmt.Println("goroutine count", runtime.NumGoroutine())
	close(bufferChn)
	w.Wait()
	fmt.Println("goroutine count", runtime.NumGoroutine())
}
