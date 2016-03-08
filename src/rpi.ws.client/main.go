package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	ctrl "rpi.ws.client/controller"
	"rpi.ws.client/raspberry"
	"rpi.ws.client/service"

	"github.com/gorilla/websocket"
)

func main() {
	// variables
	address := "ws://srv-gkdevmaster.rhcloud.com:8000/ws"
	pingTimeout := 10
	version := "0.1.9.2"

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
		case "-p":
			err := service.CheckArgs(len(os.Args), k)
			service.CheckError(err)
			pingTimeout, err = strconv.Atoi(os.Args[k+1])
			service.CheckError(err)
		}
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	// open connection
	ws, _, err := websocket.DefaultDialer.Dial(address, nil)
	service.CheckError(err)
	defer ws.Close()

	// raspberry data
	rpi := raspberry.NewRaspberry()

	// topics pool
	topics := service.TopicPool{
		"RPI1_LED0":     rpi.Led0,
		"RPI1_CPU_TEMP": rpi.CpuTemp,
		"RPI1_CPU_MEM":  rpi.CpuMemory,
		"RPI1_SYS_MEM":  rpi.SystemMemory,
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

	ticker := time.NewTicker(time.Duration(pingTimeout) * time.Minute)
	defer ticker.Stop()

	// wait for terminating
	for {
		select {
		case <-ticker.C:
			hub.Send(ctrl.GetMessage("RPI1_PING", ""))

		case <-interrupt:
			fmt.Println("Terminating...")
			os.Exit(0)
		}
	}
}
