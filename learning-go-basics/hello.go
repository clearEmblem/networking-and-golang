package main

import "fmt"
// import "rsc.io/quote"

func main() {
    fmt.Println("Hello World!")
		loops()
}


func loops(){
	myword := "akash"
	for i, v := range myword{
		fmt.Println(string(v), "is letter", i+1)
	}
}
