package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	// channel 阻塞 等到消息从 channel 中取出之后才向下执行
	channel := make(chan string, 1)
	go func() {
		fmt.Println("launch a channel")
		channel <- service()
		fmt.Println("channel executed")
	}()
	return channel
}

func TestService(t *testing.T) {
	select {
	case ref := <-AsyncService():
		t.Log(ref)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
