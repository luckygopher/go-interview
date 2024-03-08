package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Rectangle struct {
}

type Params struct {
	Length, Width int
}

func (r Rectangle) Area(params Params, result *int) error {
	*result = params.Width * params.Length
	return nil
}

func (r Rectangle) Perimeter(params Params, result *int) error {
	*result = (params.Width + params.Length) * 2
	return nil
}

func main() {
	rectangle := new(Rectangle)
	if err := rpc.Register(rectangle); err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	http.ListenAndServe(":8080", nil)
}
