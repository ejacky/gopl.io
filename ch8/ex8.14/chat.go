// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 254.
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const timeout = 30 * time.Second

//!+broadcaster
type client struct {
	Out  chan<- string // an outgoing message channel
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.Out <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			cli.Out <- "Present:"
			for c := range clients {
				cli.Out <- c.Name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

//!-broadcaster

//!+handleConn
func handleConn(conn net.Conn) {
	var name string
	in := make(chan string)
	out := make(chan string) // outgoing client messages

	go clientWriter(conn, out)
	go clientReader(conn, in)

	nameTimer := time.NewTimer(timeout)
	out <- "Enter you name:"
	select {
	case name = <-in:
	case <-nameTimer.C:
		messages <- name + " has expired"
		conn.Close()
	}

	cli := client{Out: out, Name: name}
	out <- "You are " + name
	messages <- name + " has arrived"
	entering <- cli

	timer := time.NewTimer(timeout)
	for {
		select {
		case message := <-in:
			messages <- name + ":" + message
			timer.Reset(timeout)

		case <-timer.C:
			messages <- name + " has expired"
			conn.Close()
			goto LEAVE
		}
	}

LEAVE:
	leaving <- cli
	messages <- name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
