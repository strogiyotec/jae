package internal

import (
	"syscall"
	"unsafe"
)

type Pool struct {
	fd    int //file descriptor
	notes noteQueue
}

func OpenPool() *Pool {
	poll := new(Pool)
	epoll, err := syscall.EpollCreate1(0)
	if err != nil {
		panic(err)
	}
	poll.fd = epoll
	r0, _, e0 := syscall.Syscall(syscall.SYS_EVENTFD2, 0, 0, 0)
	if e0 != 0 {
		syscall.Close(epoll)
		panic(err)
	}
}
