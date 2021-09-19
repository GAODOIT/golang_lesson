package csp

import (
	"fmt"
	"sync"
)

//3、sync.Mutex使用
//Mutex是互斥锁，用来保证在任一时刻，只能有一个例程访问某对象。Mutex 的初始值为解锁状态。Mutex 通常作为其它结构体的匿名字段使用，使该结构体具有 Lock 和 Unlock 方法。
//Mutex 可以安全的在多个例程中并行使用。

//https://blog.csdn.net/mrbuffoon/article/details/85234861


func click(c chan bool, count *int) {
	for i := 0; i < 1000; i++ {
		*count +=1
	}
	c <- true
}

func clickWithMutex(c chan bool, count *int, m *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		m.Lock()
		*count +=1
		m.Unlock()
	}
	c <- true
}

func main() {
	m := new(sync.Mutex)
	count1, count2 :=0,0
	c :=make(chan bool, 10)
	for i := 0; i < 5; i++ {
		go click(c, &count1)
	}

	for i := 0; i <5; i++ {
		go clickWithMutex(c, &count2, m)
	}

	for i := 0; i <10; i++ {
		<- c
	}
	fmt.Println("count1: ", count1)
	fmt.Println("count2: ", count2)
}