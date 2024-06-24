package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type newClient struct {
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	connection net.Conn
}

func (nC *newClient) Connect() error {
	connection, err := net.DialTimeout("tcp", nC.address, nC.timeout)
	if err != nil {
		return fmt.Errorf("connection error: %w", err)
	}
	nC.connection = connection
	log.Println("...Connected to", nC.address)
	return nil
}

func (nC *newClient) Close() error {
	if err := nC.connection.Close(); err != nil {
		return err
	}
	return nil
}

func (nC *newClient) Send() error {
	_, err := io.Copy(nC.connection, nC.in)
	if err != nil {
		return err
	}
	return nil
}

func (nC *newClient) Receive() error {
	_, err := io.Copy(nC.out, nC.connection)
	if err != nil {
		return err
	}
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &newClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}
