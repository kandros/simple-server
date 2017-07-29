package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
)

func main() {
	log.Printf("Server Listening on port 4000")
	if ln, err := net.Listen("tcp", ":4000"); err == nil {
		for {
			if conn, err := ln.Accept(); err == nil {
				reader := bufio.NewReader(conn)
				if resp, err := http.ReadRequest(reader); err == nil {
					if err := resp.Write(conn); err == nil {
						conn.Close()
					}
				}
			}
		}
	}
}
