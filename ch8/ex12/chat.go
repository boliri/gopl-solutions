package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// !+broadcaster
type client struct {
	username string
	ch       chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[chan<- string]string) // all connected clients, mapped to their name
	for {
		select {
		case msg := <-messages:
			broadcast(clients, msg)

		case cli := <-entering:
			clients[cli.ch] = cli.username

			u := getUsersOnline(clients)
			msg := buildUsersOnlineMsg(u)
			broadcast(clients, msg)

		case cli := <-leaving:
			delete(clients, cli.ch)
			close(cli.ch)
		}
	}
}

// sends a message to all users in the chat
func broadcast(clients map[chan<- string]string, msg string) {
	for cli := range clients {
		cli <- msg
	}
}

// returns a slice holding the users currently online
func getUsersOnline(clients map[chan<- string]string) []string {
	users := make([]string, 0, len(clients))
	for _, username := range clients {
		users = append(users, username)
	}
	return users
}

func buildUsersOnlineMsg(users []string) string {
	return "\nUser(s) online: " + strings.Join(users, ", ") + "\n"
}

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"

	cli := client{username: who, ch: ch}
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// !+main
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
