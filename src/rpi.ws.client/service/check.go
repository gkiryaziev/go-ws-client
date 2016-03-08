package service

import (
	"errors"
	"fmt"
	"os"
)

// Check key value
func CheckArgs(args_length, arg_index int) error {
	if args_length == (arg_index + 1) {
		return errors.New("Not specified key value.")
	}
	return nil
}

// check error
func CheckError(err error) {
	if err != nil {
		Usage()
		fmt.Println()
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
