package main

import (
	"github.com/cartmanis/call_forwarding/app/forward"
	"github.com/cartmanis/call_forwarding/app/models"
	"github.com/cartmanis/call_forwarding/logger"
)

func main() {
	c := make(chan int)
	s := &models.Settings{
		ListnerIP:   "172.22.2.60",
		ListnerPort: 7371,
		ForwardIP:   "192.168.41.26",
		ForwardPort: 3050,
		Comment:     "Some Comment",
	}
	f, err := forward.NewForward(s)
	if err != nil {
		logger.Error(err)
		return
	}
	go f.StartListner()
	<-c
}
