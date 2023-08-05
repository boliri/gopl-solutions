// this solution is based on exercise 8.12
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

var idleF = flag.Duration("idle", 300*time.Second, "for how long a user can stay idle before dropping their connection")

var server *client = &client{username: "server"}

// !+broadcaster
type client struct {
	username string
	msgCh    chan<- message // an outgoing message channel
	lastseen time.Time      // records the moment the user sent their last message to the server
	idleCh   chan struct{}  // acts as a cancellation channel when the user stays idle for too long
}

func (c *client) touch() {
	c.lastseen = time.Now()
}

type message struct {
	sender *client
	msg    string
	when   time.Time
}

func newMessage(sender *client, msg string) message {
	return message{sender: sender, msg: msg, when: time.Now()}
}

func (m message) String() string {
	return fmt.Sprintf("%s: <%s>: %s", m.when.Format(time.Kitchen), m.sender.username, m.msg)
}

var (
	entering = make(chan *client)
	leaving  = make(chan *client)
	messages = make(chan message) // all incoming client messages
)

func broadcaster() {
	clients := make(map[*client]bool) // all connected clients
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case msg := <-messages:
			broadcast(clients, msg)
			msg.sender.touch()

		case cli := <-entering:
			clients[cli] = true

			u := getUsersOnline(clients)
			msg := buildUsersOnlineMsg(u)
			broadcast(clients, newMessage(server, msg))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msgCh)

		case <-ticker.C:
			disconnectIdleUsers(clients)
		}
	}
}

// sends a message to all users in the chat
func broadcast(clients map[*client]bool, m message) {
	for cli := range clients {
		cli.msgCh <- m
	}
}

// returns a slice holding the users currently online
func getUsersOnline(clients map[*client]bool) []string {
	users := make([]string, 0, len(clients))
	for cli := range clients {
		users = append(users, cli.username)
	}
	return users
}

func buildUsersOnlineMsg(users []string) string {
	return "User(s) online: " + strings.Join(users, ", ")
}

func disconnectIdleUsers(clients map[*client]bool) {
	for cli := range clients {
		if time.Since(cli.lastseen) >= *idleF {
			close(cli.idleCh)
		}
	}
}

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan message) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- newMessage(server, "You are "+who)
	messages <- newMessage(server, who+" has arrived")

	cli := &client{username: who, msgCh: ch, lastseen: time.Now(), idleCh: make(chan struct{})}
	go func() {
		<-cli.idleCh

		// close only the reading side of the TCP connection, just so the server can
		// inform the user their connection was closed due to inactivity
		conn.(*net.TCPConn).CloseRead()
	}()

	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- newMessage(cli, input.Text())
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- newMessage(server, who+" has left")

	select {
	case <-cli.idleCh:
		// idle user spotted => send one last message to warn them, and close the writing
		// side of the connection afterwards
		fmt.Fprintln(conn, "You've been idle for too long. Disconnecting...")
		conn.(*net.TCPConn).CloseWrite()

	default:
		// manual disconnect => close the connection entirely (reading & writing sides), as well as
		// the client's idleCh to prevent goroutine leaks
		close(cli.idleCh)
		conn.Close()
	}
}

func clientWriter(conn net.Conn, ch <-chan message) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// !+main
func main() {
	flag.Parse()

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
