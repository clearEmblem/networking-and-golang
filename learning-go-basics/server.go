package main

import (
	"fmt"
	"os"
	"net"
)

func main (){
	listener, err := net.Listen("tcp", "localhost:8888")
		if err != nil{
			fmt.Println("Error occured in listening")
			os.Exit(1)
	}
	
	fmt.Println("TCP Server is listening on localhost:8888...")

	conn, err := listener.Accept()
	if err != nil{
		fmt.Println("Error occured in accepting")
		os.Exit(1)
	}

	addr := conn.RemoteAddr().String()

	fmt.Println("Local address of accepted connection is", addr)


}



func accept(listener net.Listener){
}
