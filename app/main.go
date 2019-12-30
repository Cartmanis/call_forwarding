package main

import (
	"flag"

	"github.com/cartmanis/call_forwarding/app/config"
	"github.com/cartmanis/call_forwarding/app/forward"
	"github.com/cartmanis/call_forwarding/logger"
)

func main() {	

	c := make(chan int)
	sList, err := config.ReadConfig("config.conf")
	if err != nil {
		logger.Fatal("не удалось прочитать конфигурационный файл: ", err)
	}

	for _, s := range sList {
		f, err := forward.NewForward(s)
		if err != nil {
			logger.Error(err)
			continue
		}
		go f.StartListner()
	}
	<-c
}

func getPath() string {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return "config.conf"
	}
	return args[0]
}
