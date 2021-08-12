package my_routine

// @Author: lvxiaozheng
// @Date: 2021/7/14 17:01
// @Description:

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestMyRoutine(t *testing.T) {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}

func TestRuntime(t *testing.T) {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func TestRuntime2(t *testing.T) {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}

func a() {
	for i := 1; i < 20; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 20; i++ {
		fmt.Println("B:", i)
	}
}

func TestRuntime3(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}

func TestRuntime4(t *testing.T) {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
