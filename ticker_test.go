package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	stopCh := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(stopCh)
	}()

	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())
		case <-stopCh:
			ticker.Stop()
			return
		}
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	stopCh := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(stopCh)
	}()

	for {
		select {
		case <-channel:
			fmt.Println(time.Now())
		case <-stopCh:
			return
		}
	}
}
