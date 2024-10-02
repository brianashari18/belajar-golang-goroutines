package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
	fmt.Println(counter)
}

func DisplayProses(i int, once *sync.Once, group *sync.WaitGroup) {
	group.Add(1)
	once.Do(OnlyOnce)
	fmt.Println("...Sedang proses...", i)
	group.Done()
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go DisplayProses(i, once, group)
	}

	group.Wait()
	fmt.Println("Selesai")
}
