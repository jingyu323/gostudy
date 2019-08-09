package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	for i := 1; i < 5; i++ {
		//go starts a goroutin
		go printHelloWorld(i)
	}

	time.Sleep(time.Millisecond)

	//time.Sleep(5*time.Second)

}

func printHelloWorld(i int) {

	fmt.Printf("Hello World  %d ! \n", i)
}

type Client struct {
	conn net.Conn
	Host string
}

func (c *Client) Open() error {
	conn, err := net.Dial("tcp", c.Host)
	if err != nil {
		return err
	}
	c.conn = conn

	return nil
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}

	return nil
}
