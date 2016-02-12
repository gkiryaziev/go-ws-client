package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"errors"

	ctrl "./controller"
	srvc "./service"
	rpi "./raspberry"

	"github.com/gorilla/websocket"
)

// Check key value
func checkArgs(args_length, arg_index int) error {
	if args_length == (arg_index + 1) {
		return errors.New("Not specified key value.")
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		srvc.Usage()
		fmt.Println()
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func main() {
	// variables
	address := "ws://127.0.0.1:8000/ws"
	version := "0.1.7"

	// args
	for k, arg := range os.Args {
		switch arg {
		case "-h":
			srvc.Usage()
			return
		case "-v":
			fmt.Println(version)
			return
		case "-a":
			err := checkArgs(len(os.Args), k)
			checkError(err)
			address = os.Args[k + 1]
		}
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)

	// open connection
	ws, _, err := websocket.DefaultDialer.Dial(address, nil)
	checkError(err)
	defer ws.Close()

	// topics pool
	topics := map[string]func(...string)[]byte{
		"RPI1_LED0": rpi.Led0,
		"RPI1_CPU_TEMP": rpi.CpuTemp,
		"RPI1_CPU_MEM": rpi.CpuMemory,
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
	srvc.NewSubscribe(hub).Subscribe(topics)

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
