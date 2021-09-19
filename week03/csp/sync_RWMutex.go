package csp

import (
	"fmt"
	"sync"
)

//RWMutex是读写互斥锁。
//RWMutex比Mutex多了一个“读锁定”和“读解锁”，可以让多个例程同时读取某对象。
//RWMutex 的初始值为解锁状态。RWMutex 通常作为其它结构体的匿名字段使用。
//Mutex 可以安全的在多个例程中并行使用。
//https://blog.csdn.net/mrbuffoon/article/details/85234861

func clickWithRWMutex(m *sync.RWMutex, total *int, ch chan int) {
	for i := 0; i < 1000; i++ {
		m.Lock()
		*total += 1
		m.Unlock()

		if  i == 500 {
			m.RLock()
			fmt.Println("Middle Num: ", *total)
			m.RUnlock()
		}
	}
	ch <- 1
}

func main() {
	fmt.Println("hello world!")
	m :=new(sync.RWMutex)
	ch := make(chan int, 10)
	count := 0
	for i := 0; i < 10; i++  {
		go clickWithRWMutex(m,&count, ch)
	}
	for i:=0; i < 10; i++ {
		<- ch
	}
	fmt.Println("Count: ", count)
}