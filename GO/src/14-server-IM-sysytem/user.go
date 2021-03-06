package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	go user.ListenMessage()

	return user
}

//监听当前User chanel 的方法,一旦有消息,就直接发送给对端客户端
func (t *User) ListenMessage() {
	for {
		msg := <-t.C

		t.conn.Write([]byte(msg + "\n"))
	}
}
