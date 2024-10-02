package golanggoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println(totalGoroutine)

	group.Wait()
}

func TestGomaxprocsGetThread(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println(totalCpu)

	runtime.GOMAXPROCS(10)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println(totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println(totalGoroutine)

	group.Wait()
}
