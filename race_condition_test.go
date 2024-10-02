package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				x++
				// bisa terjadi karena pada laptop yang multi core goroutines terjadi secara serentak pada waktu tertentu
				// x = 1000 + 1
				// x = 1000 + 1
				// x = 1000 + 1
				// x = 1000 + 1
				// x = 1000 + 1
				// x = 1000 + 1
				// x = 1000 + 1
				// x = 1000 + 1
				// x = 1000 + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
