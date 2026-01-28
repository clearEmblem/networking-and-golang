package main

import (
	"fmt"
	"net"
	"os"
)

// map that holds the connections
var connections = make(map[net.Conn]string)

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error occured in listening")
		os.Exit(1)
	}
	fmt.Println("TCP Server is listening on :8888...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error occured in accepting")
			os.Exit(1)
		}

		num, err := conn.Write([]byte("Enter your name: "))
		name_buf := make([]byte, 64)
		num, err = conn.Read(name_buf)

		//redundant for having a map...
		connections[conn] = string(name_buf)

		go client_func(conn, string(name_buf[0:num-1]))
	}

}

func client_func(conn net.Conn, name string) {
	addr := conn.RemoteAddr().String()

	fmt.Println("Connection accepted from", addr)

	buff := make([]byte, 1024)

	num, err := conn.Write([]byte("\n"))

	for {
		// Read data from connection to buffer
		num, err = conn.Read(buff)
		if err != nil {
			fmt.Println("Error in reading data from the connection...")
			os.Exit(1)
		}

		// Write data back to client
		write_to_all(name, num, buff, conn)
		// Print data locally
		fmt.Printf("Read %d bytes:\n%s", num, string(buff[:num]))

	}
}

func write_to_all(name string, num int, buff []byte, connection net.Conn) {
	message := fmt.Sprintf("%s: %s\n", name, string(buff[0:num]))
	for conn, _ := range connections {
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error in writing data back to clients...")
		}

	}

}
