package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type student struct {
	Name string
	Age  int
}

func main() {
	const (
		a = iota
		b = iota
	)
	const (
		name = "menglu"
		c    = iota
		d    = iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
func pase_student() {
	m := make(map[string]student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = stu
	}
	b, _ := json.Marshal(m)
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(m)
		return
	}

	fmt.Println(out.String())
}

func testTicker() {
	timer := time.NewTicker(time.Second * 1)
	go func() {
		for {
			select {
			case <-timer.C:
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()
				proc()
			}
		}
	}()
}
func proc() {
	panic("ok")
}
func testContext() {
	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	chan1 := make(chan int)
	go func(ctx context.Context) {
		for {
			select {
			case value := <-chan1:

				fmt.Print(value)

			case <-ctx.Done():
				fmt.Print("子协程结束了")
				return

			}
		}

	}(ctx)
	chan1 <- 1
	time.Sleep(time.Second * 4)
	cancel()
	time.Sleep(time.Second * 2)
}
func testTimer() {
	//设置定时器为3秒
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("当前时间为:", time.Now())

	t := <-timer.C //从定时器拿数据
	fmt.Println("当前时间为:", t)
}
func testChan() {
	ch := make(chan int)
	//quitChan := make(chan bool)

	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println(v)
			case <-time.After(time.Second * time.Duration(3)):
				//quitChan <- true
				fmt.Println("timeout, send notice")
				return
			}
		}
	}()

	for i := 0; i < 4; i++ {
		ch <- i
	}

	//<-quitChan // 输出值，相当于收到通知，解除主程阻塞
	fmt.Println("main quit out")
}
func testGoroutine() {
	intChan := make(chan int)
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			//n, _ := rand.Int(rand.Reader, big.NewInt(100))
			n := rand.Intn(5)
			intChan <- n
		}
		close(intChan)
	}()

	go func() {
		defer wg.Done()
		for {
			if v, ok := <-intChan; ok {
				fmt.Print(v)
			} else {
				break
			}
		}

	}()

	wg.Wait()
}
func syncWait() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("A: ", i)
			wg.Done()
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}
func (t *Teacher) ShowA() {
	fmt.Println("teacher showA")
	t.ShowB()
}
