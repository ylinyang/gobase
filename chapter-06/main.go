package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	group := ListAddressGroup()
	address := ShowAddressGroup(group[0])
	//fmt.Println(address)

	if len(os.Args) < 2 {
		for _, v := range group {
			fmt.Printf("IP地址组ID：%s\n包含IP地址如下：%s\n", v, address[v])
		}
	} else if os.Args[1] == "add" {
		if err := UpdateAddressGroup(os.Args[2:], group[0], "add"); err != nil {
			log.Panicln(err)
		}
	} else if os.Args[1] == "del" {
		if os.Args[2] == "all" {
			UpdateAddressGroup([]string{}, group[0], "del all")
		} else if err := UpdateAddressGroup(os.Args[2:], group[0], "del"); err != nil {
			log.Panicln(err)
		}
	} else {
		fmt.Println("参数有误：add/del")
	}
}

/*
del xxx
add xxx
*/
