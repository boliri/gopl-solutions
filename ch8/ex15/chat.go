// this solution is based on exercise 8.14
//
// to simulate msg drops, play with the max-wait command-line argument, and add artificial delays
// to the clientWriter function
//
// example:
//
//	func clientWriter(conn net.Conn, cli *client) {
//		for msg := range cli.msgCh {
//			cli.rlock()
//	        if cli.username == "foo" {
//	            time.Sleep(...) // set a value lower than max-wait to see the msg dropping feature in action
//	        }
//			fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
//			cli.runlock()
//		}
//	}
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
var maxWaitF = flag.Duration("max-wait", 2*time.Second, "for how long the server is willing to wait for a user to read a message before dropping the next one")

var server *client = &client{username: "server"}

// !+broadcaster
type client struct {
	username string
	msgCh    chan message // a message channel
	lastseen time.Time    // records the moment the user sent their last message to the server
	reading  bool         // whether the client is currently reading a message from the broadcaster or not
}

func (c *client) touch() {
	c.lastseen = time.Now()
}

func (c *client) isIdle() bool {
	return time.Since(c.lastseen) >= *idleF
}

func (c *client) isReading() bool {
	return c.reading
}

func (c *client) rlock() {
	c.reading = true
}

func (c *client) runlock() {
	c.reading = false
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
// messages might be dropped if a user is not ready to read it
func broadcast(clients map[*client]bool, m message) {
	for cli := range clients {
		if cli.isReading() {
			// deal with the slow-reading client in a separate goroutine so other clients
			// don't get stuck
			go func() {
				now := time.Now()
				for {
					stillReading := cli.isReading()
					waitMore := time.Since(now) < *maxWaitF
					if stillReading && waitMore { // we can wait a bit more
						continue
					} else if !stillReading && waitMore { // user is free, let's send them the msg
						cli.msgCh <- m
						break
					} else { // no more chances; msg is dropped for this user
						log.Printf(
							"broadcast: message \"%s\" from <%s> dropped for <%s>\n",
							m.msg, m.sender.username, cli.username,
						)
						break
					}
				}
			}()
			continue
		}

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
	cli := &client{msgCh: ch, lastseen: time.Now()}

	go clientWriter(conn, cli)

	var who string
	input := bufio.NewScanner(conn)

	fmt.Fprint(conn, "Welcome to the chat! Please, type your nickname: ")
	for input.Scan() {
		// CAVEAT: right now, multiple users can share the same nickname as we never check for
		// collisions
		who = input.Text()
		break
	}
	cli.username = who

	ch <- newMessage(server, "You are "+who)
	messages <- newMessage(server, who+" has arrived")

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

func clientWriter(conn net.Conn, cli *client) {
	for msg := range cli.msgCh {
		cli.rlock()
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
		cli.runlock()
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
