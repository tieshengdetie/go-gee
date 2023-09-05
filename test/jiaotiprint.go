package main

import (
	"fmt"
)

/*问题描述

使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

func main() {

	letter, number, done := make(chan bool), make(chan bool), make(chan bool)
	//wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-letter:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				number <- true
			}
		}
	}()
	go func() {
		j := 'A'
		for {
			select {
			case <-number:
				if j >= 'Z' {
					done <- true

				} else {
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					letter <- true
				}

			}
		}
	}()
	letter <- true
	for {
		select {
		case <-done:
			return
		}
	}

}
