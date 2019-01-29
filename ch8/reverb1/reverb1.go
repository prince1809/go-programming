package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration)  {
	fmt.Fprintln(c, "From Server:", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "From Server:", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "From Server:", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main()  {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}

}