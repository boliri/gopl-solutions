package main

import (
	"bufio"
	"errors"
	"fmt"
	"ftp"
	"io"
	"io/fs"
	"log"
	"net"
	"os"
	"path"
	"strings"
	"text/tabwriter"
)

const rootPath = "./rootdir" // server's root dir path

var serverFS = ftp.FtpFS(rootPath)

// session holds hosts connected to the FTP server (i.e., clients) and their current working dir
type session struct {
	host, cwd string
}

// sessions maps hosts connected to the FTP server with their session
var sessions = make(map[string]*session)

func handleConn(c net.Conn) {
	defer c.Close()

	fmt.Fprint(c, "\nWelcome to the FTP server!\n\n")

	clientAddr := c.RemoteAddr().String()
	if _, ok := sessions[clientAddr]; !ok {
		sessions[clientAddr] = &session{host: clientAddr, cwd: "/"}
	}
	sess := sessions[clientAddr]

	s := bufio.NewScanner(c)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		var cmd, arg string
		_, err := fmt.Sscanf(s.Text(), "%s %s", &cmd, &arg)
		if err != nil && err != io.EOF {
			fmt.Fprintf(c, "error: scan: %s\n\n", err)
			continue
		}

		if strings.TrimSpace(cmd) == "" {
			continue
		}

		switch cmd {
		case "cd":
			// go to the server's root dir
			if arg == "" {
				sess.cwd = "/"
				fmt.Fprintf(c, "cwd: %s\n\n", sess.cwd)
				continue
			}

			// go to dir specified in arg, starting from the server's root dir
			if strings.HasPrefix(arg, "/") {
				_, err := serverFS.Stat(arg)
				if err != nil {
					fmt.Fprintf(c, "%s: %s\n\n", cmd, err)
					continue
				}

				sess.cwd = arg
				fmt.Fprintf(c, "cwd: %s\n\n", sess.cwd)
				continue
			}

			// noop
			if arg == "." {
				fmt.Fprintf(c, "cwd: %s\n\n", sess.cwd)
				continue
			}

			// move to dir as specified in arg
			// special case: if arg is "..", the command will go one level back in the dir tree hierarchy
			filepath := path.Join(sess.cwd, arg)
			fi, err := serverFS.Stat(filepath)
			if err != nil {
				fmt.Fprintf(c, "%s: %s\n\n", cmd, err)
				continue
			}

			if !fi.IsDir() {
				fmt.Fprintf(c, "%s: %s: not a directory\n\n", cmd, filepath)
				continue
			}

			sess.cwd = filepath
			fmt.Fprintf(c, "cwd: %s\n\n", sess.cwd)
		case "ls":
			var filepath string
			if strings.HasPrefix(arg, "/") {
				filepath = arg
			} else {
				filepath = path.Join(sess.cwd, arg)
			}

			var fis []fs.FileInfo
			if path.Ext(filepath) != "" {
				// file
				entry, err := serverFS.Stat(filepath)
				if err != nil {
					fmt.Fprintf(c, "%s: %s\n\n", cmd, err)
					continue
				}

				fis = []fs.FileInfo{entry}
			} else {
				// dir
				entries, err := serverFS.ReadFSDirectory(filepath)
				if err != nil {
					fmt.Fprintf(c, "%s: %s\n\n", cmd, err)
					continue
				}

				if len(entries) == 0 {
					fmt.Fprint(c, "<empty dir>\n\n")
					continue
				}

				for _, e := range entries {
					fi, err := e.Info()
					if errors.Is(err, os.ErrNotExist) {
						// skip files that have been renamed or removed since the dir was read
						continue
					}

					fis = append(fis, fi)
				}
			}

			tabw := tabwriter.NewWriter(c, 0, 0, 0, ' ', tabwriter.AlignRight)
			for _, fi := range fis {
				fmt.Fprintf(
					tabw,
					"%s\t   %dB\t  %s\t  %d\t %d:%d\t %s\n",
					fi.Mode(),
					fi.Size(),
					fi.ModTime().Month().String()[:3],
					fi.ModTime().Day(),
					fi.ModTime().Hour(),
					fi.ModTime().Minute(),
					fi.Name(),
				)
			}
			tabw.Flush()
			fmt.Fprintln(c)
		case "cat":
			var filepath string
			if strings.HasPrefix(arg, "/") {
				filepath = arg
			} else {
				filepath = path.Join(sess.cwd, arg)
			}

			b, err := serverFS.ReadFSFile(filepath)
			if err != nil {
				fmt.Fprintf(c, "%s: %s\n\n", cmd, err)
				continue
			}

			fmt.Fprintf(c, "%s\n\n", b)
		case "quit", "exit":
			fmt.Fprint(c, "disconnecting...\n\n")
			delete(sessions, clientAddr)
			return
		default:
			fmt.Fprintf(c, "%s: unknown command\n\n", cmd)
			continue
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:2121")
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
