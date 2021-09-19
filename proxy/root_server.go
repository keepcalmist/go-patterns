package main

import "strings"

type SomeRootServer struct {
	leafServer *someLeafServer
}

func NewRootServer() *SomeRootServer {
	return &SomeRootServer{
		leafServer: NewLeafServer(),
	}
}

func (s *SomeRootServer) HandleRequest(s1, s2 string) (int, string) {
	if s.checkSomething(s1) {
		return 0, "Permission denied"
	}

	return s.leafServer.HandleRequest(s1, s2)
}

func (s *SomeRootServer) checkSomething(s1 string) bool {
	return strings.Contains(s1, "admin")
}
