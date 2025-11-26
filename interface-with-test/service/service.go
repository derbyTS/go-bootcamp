package service

import (
	"example.com/project/client"
)

type Service struct {
	C client.Client
}

func (s Service) Process() string {
	data := s.C.FetchData()
	return "Process: " + data
}
