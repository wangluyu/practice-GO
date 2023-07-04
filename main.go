package main

import (
	"fmt"
	"os"
	log_wire "practice/log-wire"
)

func main() {
	logger, err := log_wire.InitLogger()
	if err != nil {
		fmt.Printf("failed to create logger: %s\n", err)
		os.Exit(2)
	}
	logger.Info("hello", "name", "wangluyu")
}
