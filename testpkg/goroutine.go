package testpkg

import (
	"fmt"
	"runtime"
	"sync"
)

// TestGoRoutine Goroutine 예제
func TestGoRoutine() {
	// Go 는 기본적으로 하나의 cpu 만 사용하지만
	// 몇 개의 cpu 를 사용할 지 지정해줄 수 있다.
	runtime.GOMAXPROCS(4)

	// Wait group 생성
	var wait sync.WaitGroup
	wait.Add(2)

	say := func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}
	go say("1")
	go say("2")

	// 대기한다.
	wait.Wait()
}

func TestChannel() {

	// Unbuffered channel
	// 수신자가 받지 않으면 데드락에 걸린다.
	ch := make(chan int)
	go func() {
		ch <- 123
	}()

	var i int
	i = <-ch
	fmt.Println(i)

	// Buffered channel
	// 수신자가 받을 준비가 되어 있지 않아도 버퍼 크기 만큼 데이터를 보내고
	// 다른 일을 할 수 있다.
	c := make(chan int, 1)
	c <- 101
	fmt.Println(<-c)
}
