package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Length, Width int
}

func main() {
	result := 0
	cl, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	if err = cl.Call("Rectangle.Area", Params{
		Length: 100,
		Width:  20,
	}, &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("面积是", result)
	if err = cl.Call("Rectangle.Perimeter", Params{
		Length: 10,
		Width:  20,
	}, &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("周长是", result)
}
