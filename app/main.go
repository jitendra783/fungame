package main

import (
	"fungame/api"
	"fungame/pkg/config"
	"log"
	"os"
)

func main() {
	env := ""
	host := os.Getenv("SERVER_HOST")
	if host != "" {
		env = "server"
	} else {
		if len(os.Args) == 2 {
			env = os.Args[1]
		} else {
			env = os.Args[2]
		}
	}
	config.Load(env)
	if err := api.Start(); err != nil {
		log.Fatal("Failed to start server, err:", err)
		os.Exit(1)
	}
}
