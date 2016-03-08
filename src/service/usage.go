package service

import (
	"fmt"
	"os"
	"path/filepath"
)

// Usage menu
func Usage() {
	a := filepath.Base(os.Args[0])
	fmt.Println()
	fmt.Println("Usage:", a, "[OPTIONS]")
	fmt.Println()
	fmt.Println("    Websocket client.")
	fmt.Println()
	fmt.Println("    Websocket client for Raspberry Pi.")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("    -a    STR        Server address. [ws://127.0.0.1:8000/ws]")
	fmt.Println("    -p    INT        Ping timeout.   [10 min]")
	fmt.Println()
	fmt.Println("    -h               This help.")
	fmt.Println("    -v               Print version.")
}
