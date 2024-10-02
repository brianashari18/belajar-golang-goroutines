package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "Habis coy"
		},
	}
	group := &sync.WaitGroup{}

	pool.Put("Brian")
	pool.Put("Anashari")
	pool.Put("Puyol")

	group.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(1 * time.Second)
			defer group.Done()
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Wait()
	fmt.Println("Selesai")
}
