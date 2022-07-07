package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type server struct {
	ip   string
	port int
	// 存储用信息的 name + user
	onLineMap map[string]*user
	mapLock   sync.RWMutex
	// 广播消息
	message chan string
}

func NewServer(ip string, port int) *server {
	return &server{ip: ip, port: port, onLineMap: make(map[string]*user), message: make(chan string)}
}

func (s *server) Start() {
	// 启动一个tcp server
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.ip, s.port))
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()

	//  监听广播的message
	go s.listenMessages()

	//  事件处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.handler(conn)
	}

}

func (s *server) handler(conn net.Conn) {
	// 实例化上线用户
	u := NewUser(conn, s)
	u.online()
	// 广播消息
	s.broadcast(u, "已上线")

	// 接受用户的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				u.offline()
				return
			}
			if err != nil && err != io.EOF {
				log.Panicln("read error", err)
				return
			}
			// 提取用户的消息
			msg := string([]byte(buf[:n-1]))
			u.doMessage(msg)
		}
	}()

	// 阻塞
	select {}
}

// 将上线消息写入到message中
func (s *server) broadcast(u *user, msg string) {
	sendMsg := "[" + u.name + "]" + u.addr + ":" + msg
	s.message <- sendMsg
}

func (s *server) listenMessages() {
	for {
		msg := <-s.message
		s.mapLock.Lock()
		for _, ch := range s.onLineMap {
			// client里面的channel 写入msg
			ch.ch <- msg
		}
		s.mapLock.Unlock()
	}

}

/*
1. server里面新添加两个字段 一个存储在线用户新的(需要有锁的操作)  一个用于广播上线消息的  同时修改start方法
2. 服务器监听到用户上线了，需要将用添加到map中 同时广播消息给所有的客户端
3. 通过一个监听所有的message消息的函数 将消息发送给所有的客户端

1. 接收用户的消息 并且广播出去
*/
