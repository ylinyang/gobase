package main

import (
	"log"
	"net"
)

type user struct {
	name string
	addr string
	ch   chan string // channel接收服务器写入的消息
	conn net.Conn    // 与服务器的连接信息

	s *server
}

func NewUser(conn net.Conn, s *server) *user {
	u := &user{
		name: conn.RemoteAddr().String(),
		addr: conn.RemoteAddr().String(),
		ch:   make(chan string),
		conn: conn,
		s:    s,
	}

	// 启动一个goroutine 从ch中获取服务器发来的信息
	go u.listerHandler()

	return u
}

func (u *user) listerHandler() {
	for {
		msg := <-u.ch // 从ch中读取信息 写入到conn中给client展示
		if _, err := u.conn.Write([]byte(msg + "\n")); err != nil {
			log.Panicln(err)
		}
	}
}

func (u *user) online() {
	// 添加用户到map中
	u.s.mapLock.Lock()
	u.s.onLineMap[u.name] = u
	u.s.mapLock.Unlock()
}

func (u *user) offline() {
	u.s.mapLock.Lock()
	defer u.s.mapLock.Unlock()
	u.s.broadcast(u, "下线")
	delete(u.s.onLineMap, u.name)
}

func (u *user) doMessage(msg string) {
	u.s.broadcast(u, msg)
}
