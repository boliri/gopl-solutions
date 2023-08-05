// this solution is based on exercise 8.13 (chatv2.go)
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
}

func (c *client) touch() {
	c.lastseen = time.Now()
}

func (c *client) isIdle() bool {
	return time.Since(c.lastseen) >= *idleF
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

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan message) // outgoing client messages
	go clientWriter(conn, ch)

	var who string
	input := bufio.NewScanner(conn)

	fmt.Fprint(conn, "Welcome to the chat! Please, type your nickname: ")
	for input.Scan() {
		// CAVEAT: right now, multiple users can share the same nickname as we never check for
		// collisions
		who = input.Text()
		break
	}

	ch <- newMessage(server, "You are "+who)
	messages <- newMessage(server, who+" has arrived")

	cli := &client{username: who, msgCh: ch, lastseen: time.Now()}

	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			<-ticker.C
			if cli.isIdle() {
				conn.(*net.TCPConn).CloseRead() // don't let the user type anything else
				fmt.Fprintln(conn, "You've been idle for too long. Disconnecting...")
				close(done)
				break
			}
		}
		ticker.Stop()
	}()

	entering <- cli

	for input.Scan() {
		messages <- newMessage(cli, input.Text())
	}
	// NOTE: ignoring potential errors from input.Err()

	select {
	case <-done: // wait for the server to warn an idle user that their connection is gonna be dropped

	default: // either the user disconnected manually, or there was a connection error

	}

	leaving <- cli
	messages <- newMessage(server, who+" has left")

	// two things can happen at this point:
	//   a. if the user was idle, then conn.Close() closes the writing side of the connection
	//      only, since the reading one is already closed
	//   b. if the user disconnected from the server manually, or there was a connection error,
	//      conn.Close() closes the connection entirely (writing + reading sides)
	conn.Close()
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
