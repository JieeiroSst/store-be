package server

import (

)

type Server interface {
	AppServerAPI() error
}

type server struct {

}

func NewSserver() Server {
	return &server{}
}

func (s *server) AppServerAPI() error {
	
	return nil 
}