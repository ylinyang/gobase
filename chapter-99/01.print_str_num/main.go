package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	str, num := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-num:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				str <- true
				break
			default:
				break
			}
		}
	}()
	wg.Add(1)
	go func(w *sync.WaitGroup) {
		s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-str:
				// 子串统计
				if i >= strings.Count(s, "")-1 { // len(s) +1
					w.Done()
					return
				}
				fmt.Print(s[i : i+1])
				i++
				if i >= strings.Count(s, "") {
					i = 0
				}
				fmt.Print(s[i : i+1])
				i++
				num <- true
				break
			default:
				break
			}
		}
	}(&wg)
	num <- true
	wg.Wait()
}

/*
结果：12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
思路：
使⽤ channel 来控制打印的进度。使⽤两个 channel ，来分别控制数字和字⺟的打印序列，
数字打印完成后通过 channel通知字⺟打印,
字⺟打印完成后通知数字打印，然后周⽽复始的⼯作。


*/
