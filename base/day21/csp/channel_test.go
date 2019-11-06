package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 100)
	return "Done"
}

func otherTask() {
	fmt.Println("working other task start")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("other task worked")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// channel 阻塞 等到消息从 channel 中取出之后才向下执行
	channel := make(chan string)
	go func() {
		fmt.Println("launch a channel")
		channel <- service()
		fmt.Println("channel executed")
	}()
	return channel
}

func AsyncBufService() chan string {
	// buf channel 直接返回结果
	channel := make(chan string, 1)
	go func() {
		fmt.Println("launch a channel")
		channel <- service()
		fmt.Println("channel executed")
	}()
	return channel
}

func TestChannel(t *testing.T) {
	channel := AsyncService()

	otherTask()
	fmt.Println(<-channel)
	time.Sleep(time.Second * 1)
}

func TestBufChannel(t *testing.T) {
	channel := AsyncBufService()
	otherTask()
	fmt.Println(<-channel)
	time.Sleep(time.Second * 1)
}
