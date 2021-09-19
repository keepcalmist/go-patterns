package main

import "fmt"

type someLeafServer struct {
}

func NewLeafServer() *someLeafServer {
	return &someLeafServer{}
}

func (s *someLeafServer) HandleRequest(s1, s2 string) (int, string) {
	return 1488, fmt.Sprintf("%s🤪🤪🤪%s🤪🤪🤪KEK", s1, s2)
}
