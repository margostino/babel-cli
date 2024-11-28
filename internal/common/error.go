package common

import (
	"fmt"
	"log"
	"os"
)

func Check(err error, message string) {
	if err != nil {
		fmt.Printf("Error: %s - %s\n", err.Error(), message)
		os.Exit(1)
	}
}

func CheckPanic(err error, message string) {
	if err != nil {
		fmt.Printf("Error: %s - %s\n", err.Error(), message)
		panic(err)
	}
}

//func Check(err error, message string) {
//	if err != nil {
//		log.Fatalf("Error: %s - %s\n", err.Error(), message)
//	}
//}

func SilentCheck(err error, message string) {
	if err != nil {
		log.Printf("Error: %s - %s\n", err.Error(), message)
	}
}

func IsError(err error, message string) bool {
	if err != nil {
		log.Printf("Error: %s - %s\n", err.Error(), message)
		return true
	}
	return false
}
