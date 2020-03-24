//DATA-SERVER APPLICATION
package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("Skywire data-server is running")
	ln, err := net.Listen("tcp", "localhost:9020")

	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		io.WriteString(conn, fmt.Sprint("Skywire data-server Ready\n", time.Now(), "\n"))

		conn.Close()
	}
}
