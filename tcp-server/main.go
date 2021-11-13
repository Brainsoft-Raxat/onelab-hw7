package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"tcp-server/server"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	//ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		//cancel()
		oscall := <-c
		log.Printf("system call:%+v", oscall)
	}()

	addr := "127.0.0.1:8081"

	srv := server.NewServer(addr)
	wg.Wait()
	srv.Stop()
}
