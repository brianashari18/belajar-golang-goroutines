package golanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespone(channel1)
	go GiveMeRespone(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2: ", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}

		if counter == 2 {
			break
		}
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeRespone(channel1)
	go GiveMeRespone(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1: ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2: ", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		fmt.Println("Memulai loop...")
		for i := 0; i < 10; i++ {
			fmt.Println("Mengirim data...")
			channel <- "Perulangan ke-" + strconv.Itoa(i)
			fmt.Println("Data berhasil dikirim --")
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data: ", data)
	}
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	fmt.Println(cap(channel))
	defer close(channel)

	go func() {
		channel <- "Brian"
		channel <- "Anashari"
		channel <- "Puyol"

	}()

	time.Sleep(2 * time.Second)
	fmt.Println(len(channel))

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(3 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Brian Anashari"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func GiveMeRespone(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Brian Anashari"
	fmt.Println("Selesai mengirim data")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeRespone(channel)

	data := <-channel
	time.Sleep(2 * time.Second)
	fmt.Println(data)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Brian Anashari"
		fmt.Println("Selesai mengirim data")
	}()

	data := <-channel
	time.Sleep(2 * time.Second)
	fmt.Println(data)
}
