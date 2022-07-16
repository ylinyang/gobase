package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

func main() {
	sum := 1
	a := "netstat -tulnp |grep 8090"
	b := "ps -aux |grep '/usr/sbin/sshd -D好'|grep -v grep"
	for {
		fmt.Println(sum)
		time.Sleep(3 * time.Second)
		if check(a) && sum < 3 {
			fmt.Println("端口正常")
			if check(b) {
				fmt.Println("进程正常")
			} else {
				if check("touch q.txt") {
					sum += 1
					fmt.Println("启动进程")
				}
			}
		}
	}
}

func check(s string) bool {
	cmd := exec.Command("/bin/bash", "-c", s)
	output, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("无法获取命令的标准输出管道", err.Error())
		return false
	}

	// 执行Linux命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Linux命令执行失败，请检查命令输入是否有误", err.Error())
		return false
	}

	// 读取所有输出
	_, err = ioutil.ReadAll(output)
	if err != nil {
		fmt.Println("打印异常，请检查")
		return false
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait", err.Error())
		return false
	}

	//fmt.Printf("%s", bytes)
	return true
}
