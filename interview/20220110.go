/*
2022-01-10 深信服 技术一面 算法题1
golang两个协程交替打印1-100的奇数偶数
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//solution1()
	solution2()
}

// 方法一：CPU调度
func solution1() {
	//设置可同时使用的CPU核数为1
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 1; i <= 100; i++ {
			//奇数
			if i%2 == 1 {
				fmt.Println("goroutine 1 :", i)
			}
			//让出cpu
			runtime.Gosched()
		}
	}()
	go func() {
		for i := 1; i <= 100; i++ {
			//偶数
			if i%2 == 0 {
				fmt.Println("goroutine 2 :", i)
			}
			//让出cpu
			runtime.Gosched()
		}
	}()
	// 睡眠，防止 main 直接结束
	time.Sleep(3 * time.Second)
	//select {}
	fmt.Println("end")
}

// 方法二：channel
func solution2() {
	c := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			// 发送数据，发完一个阻塞
			c <- 1
			// 奇数
			if i%2 == 1 {
				fmt.Println("goroutine 1 :", i)
			}
		}
	}()
	go func() {
		for i := 1; i <= 100; i++ {
			// 接收数据，接到一个阻塞
			<-c
			// 偶数
			if i%2 == 0 {
				fmt.Println("goroutine 2 :", i)
			}
		}
	}()
	// 睡眠，防止 main 直接结束
	time.Sleep(3 * time.Second)
	//select {}
	fmt.Println("end")
}
