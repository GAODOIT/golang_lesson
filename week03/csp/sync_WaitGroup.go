package csp

import (
	"fmt"
	"sync"
	"time"
)


//golang 上下文context用法详解
//https://www.cnblogs.com/sunlong88/p/11272559.html

//WaitGroup是组等待。
//WaitGroup 用于等待一组例程的结束。主例程在创建每个子例程的时候先调用 Add 增加等待计数，
//每个子例程在结束时调用 Done 减少例程计数。之后，主例程通过 Wait 方法开始等待，直到计数器归零才继续执行。

func wgProcess(wg *sync.WaitGroup, index int) {
	fmt.Println("WgProcess ", index, "is going!")
	time.Sleep(1*time.Second)
	wg.Done()
}

func main() {
	wg := new(sync.WaitGroup)
	for i :=0; i < 20; i++ {
		wg.Add(1)
		go wgProcess(wg, i)
	}

	wg.Wait()
	fmt.Println("The game is over!")
}
