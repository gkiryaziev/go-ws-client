package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	ctrl "./controller"
	rpi "./raspberry"
	"./service"

	"github.com/gorilla/websocket"
)

func main() {
	// variables
	address := "ws://127.0.0.1:8000/ws"
	version := "0.1.7"

	// args
	for k, arg := range os.Args {
		switch arg {
		case "-h":
			service.Usage()
			return
		case "-v":
			fmt.Println(version)
			return
		case "-a":
			err := service.CheckArgs(len(os.Args), k)
			service.CheckError(err)
			address = os.Args[k+1]
		}
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	// open connection
	ws, _, err := websocket.DefaultDialer.Dial(address, nil)
	service.CheckError(err)
	defer ws.Close()

	// topics pool
	topics := service.TopicPool{
		"RPI1_LED0":     rpi.Led0,
		"RPI1_CPU_TEMP": rpi.CpuTemp,
		"RPI1_CPU_MEM":  rpi.CpuMemory,
	}

	log.Println("Connected to", address)
	time.Sleep(time.Second)

	// main circle
	hub := ctrl.NewHub(ws, topics)
	go hub.Run()
	time.Sleep(time.Second)
	go hub.Writer()
	go hub.Reader()
	time.Sleep(time.Second)

	// subscribe
	ctrl.NewSubscribe(hub).Subscribe(topics)

	//ticker := time.NewTicker(time.Second * 3)
	//defer ticker.Stop()

	// publish
	//list := []func() []byte{rpi.CpuTemp, rpi.CpuMemory}

	// wait for terminating
	for {
		select {
		//case <-ticker.C:
		//	for _, f := range list {
		//		hub.Send(f())
		//		//time.Sleep(time.Millisecond * 500)
		//	}

		case <-interrupt:
			fmt.Println("Stoping...")
			os.Exit(0)
		}
	}
}
