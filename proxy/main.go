package main

import "fmt"

func main() {

	srv := NewRootServer()

	fmt.Println(srv.HandleRequest("admin", ""))
	fmt.Println(srv.HandleRequest("Kirill", "Unschikov"))

}
