package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	servAddr := "localhost:8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	start := time.Now()

	for i := 0; i < 1000; i++ {
		_, err = conn.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("write to server = ", strconv.Itoa(i))

		reply := make([]byte, 1024)

		_, err = conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("reply from server=", string(reply))
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)

	conn.Close()
}
