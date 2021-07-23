/*
 * @Description:
 * @Author: Caoxiang
 * @Date: 2021-05-27 18:08:47
 * @LastEditors: Caoxiang
 */
package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Counter struct {
	count int
	// mu    sync.Mutex
	mu sync.RWMutex
}

var wg sync.WaitGroup

func addSafety() {
	wg.Add(50)
	c := &Counter{}
	for i := 1; i <= 50; i++ {
		go func() {
			c.mu.Lock()
			c.count++
			c.mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.count)
}

func doOnce() {
	var once sync.Once
	var wg sync.WaitGroup
	body := func() {
		fmt.Println("println once")
	}

	wg.Add(11)
	for i := 0; i <= 10; i++ {
		go func() {
			once.Do(body)
			wg.Done()
		}()
	}
	wg.Wait()
}

func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d号种子准备完毕\n", i)
			cond.L.Lock()
			cond.Wait()
			fmt.Printf("%d号种子选手开始起跑...\n", i)
			cond.L.Unlock()
		}(i)
	}

	//协程
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经准备就位,准备发令枪")
		fmt.Println("比赛开始")
		cond.Broadcast()
	}()

	wg.Wait()
}

//10个人赛跑，1个裁判发号施令
func Race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock()
			cond.Wait() //等待发令枪响
			fmt.Println(num, "号开始跑……")
			cond.L.Unlock()
		}(i)
	}
	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始～")
		cond.Broadcast() //发令枪响
	}()
	//防止函数提前返回退出
	wg.Wait()
}

func mapTest() {
	// var m map[string]map[string]string
	m := make(map[string]map[string]string)
	m["person"] = map[string]string{"name": "caoxiang", "age": "29"}
	m["person1"] = map[string]string{"name": "caoxiang", "age": "29"}
	for k, v := range m {
		name, _ := v["name"]
		age, _ := v["age"]
		fmt.Printf("%s`s name:%s age:%s\n", k, name, age)
	}

}

func sliceCap() {
	var arr []int
	fmt.Printf("begin cap:%d\n", cap(arr))
	for i := 0; i < 1024; i++ {
		arr = append(arr, i)
	}
	fmt.Printf("old cap:%d\n", cap(arr))
	arr = append(arr, 1)
	fmt.Printf("except cap:%d\n", 1280)
	fmt.Printf("real cap:%d\n", cap(arr))
}

func channelClose() {
	// var ch chan int
	ch := make(chan string, 2)
	go func(ch chan string) {
		ch <- "1"
		ch <- "2"
	}(ch)
	<-ch
	close(ch)
	fmt.Println("channel closed")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func stringBuilder() {
	var s strings.Builder
	for i := 0; i <= 10; i++ {
		s.WriteString("a")
	}
	fmt.Println(s.String())
}

func main() {
	// doOnce()
	// race()
	// mapTest()
	// sliceCap()
	// channelClose()
	//grep
	stringBuilder()
}
