package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"

	"github.com/gkiryaziev/go-ws-client/conf"
	ctrl "github.com/gkiryaziev/go-ws-client/controller"
	"github.com/gkiryaziev/go-ws-client/raspberry"
	"github.com/gkiryaziev/go-ws-client/service"
)

// checkError check errors
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// config
	config, err := conf.NewConfig("config.yaml").Load()
	checkError(err)

	// interrupt
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// open connection
	ws, _, err := websocket.DefaultDialer.Dial(config.Server.Address, nil)
	checkError(err)
	defer ws.Close()

	// raspberry data
	rpi := raspberry.NewRaspberry()

	// topics pool
	topics := service.TopicPool{
		"RPI1_LED0":          rpi.Led0,
		"RPI1_CPU_TEMP":      rpi.CpuTemp,
		"RPI1_CPU_MEM":       rpi.CpuMemory,
		"RPI1_CPU_CORE_VOLT": rpi.CpuCoreVolt,
		"RPI1_SYS_MEM":       rpi.SystemMemory,
	}

	log.Println("Connected to", config.Server.Address)
	time.Sleep(time.Second)

	// main circle
	hub := ctrl.NewHub(ws, topics, config.Debug)
	go hub.Run()
	time.Sleep(time.Second)
	go hub.Writer()
	go hub.Reader()
	time.Sleep(time.Second)

	// subscribe
	ctrl.NewSubscribe(hub).Subscribe(topics)

	ticker := time.NewTicker(time.Duration(config.Server.PingTimeout) * time.Minute)
	defer ticker.Stop()

	// wait for terminating
	for {
		select {
		case <-ticker.C:
			hub.Send(ctrl.GetMessage("RPI1_PING", ""))

		case <-interrupt:
			fmt.Println("Clean and terminating...")
			os.Exit(0)
		}
	}
}
