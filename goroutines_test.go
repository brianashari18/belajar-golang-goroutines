package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(n int) {
	fmt.Println(n)
}
func TestHellowWorld(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}
