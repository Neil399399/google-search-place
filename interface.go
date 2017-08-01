package main

import (
	"fmt"
	"google_place/datamodel"
	"net"
)

type Storage interface {
	Read(id string) (datamodel.Coffee, error)
	Write(data datamodel.Coffee) error
}

type StoppableListener struct {
	*net.TCPListener          //Wrapped listener
	stop             chan int //Channel used only to indicate listener should shutdown
}
type StoreImpl struct {
}

func (s StoreImpl) Read(id string) (datamodel.Coffee, error) {
	fmt.Println("READ?")

	return datamodel.Coffee{}, nil
}

func (s StoreImpl) Write(data datamodel.Coffee) error {
	fmt.Println("WRITE?")

	return nil
}
