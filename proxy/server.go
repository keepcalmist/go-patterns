package main

type server interface {
	HandleRequest(s1, s2 string) (int, string)
}
