package main

import (
	"github.com/cartmanis/call_forwarding/app/config"
	"github.com/cartmanis/call_forwarding/app/forward"
	"github.com/cartmanis/call_forwarding/logger"
)

func main() {
	c := make(chan int)
	sList, err := config.ReadConfig()	
	if err != nil {
		logger.Fatal("не удалось прочитать конфигурационный файл: ", err)
	}
		
	for _, s := range sList {
		f, err := forward.NewForward(s)
		if err != nil {
			logger.Error(err)
			return
		}
		go f.StartListner()
	}
	<-c
}
