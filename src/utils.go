package main

import (
	"fmt"
	"os"
)

// NULL - Created this just for me, for better readability.
// This is the closest thing, you have to Enums in Go
const (
	NULL = ""
)

func checkError(err error, message string, important bool) {
	if err != nil {
		fmt.Printf("[-] ERROR! Details:\n")
		fmt.Printf("[-]  - %s failed!\n", message)
		fmt.Printf("[-]  - %s\n", err)
		if important {
			os.Exit(1)
		}
	}
}

func createDirectory(directoryName string) error {
	permission := os.FileMode(777) //The folder can be read and written by EVERYONE

	err := os.MkdirAll(directoryName, permission)
	if err != nil && os.IsNotExist(err) {
		return err
	}
	return nil
}
