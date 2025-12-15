package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

var count int64

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: %s <listen_addr> <sleep_ms>", os.Args[0])
	}

	addr := os.Args[1]
	sleepMs, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Mock TCP server listening on %s, sleep=%dms", addr, sleepMs)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		go handleConn(conn, time.Duration(sleepMs)*time.Millisecond)
	}
}

func handleConn(conn net.Conn, sleep time.Duration) {
	defer conn.Close()

	curCount := atomic.AddInt64(&count, 1)
	fmt.Printf("Receive connection %d\n", curCount)
	// Optional read (forces real I/O)
	reader := bufio.NewReader(conn)
	_, _ = reader.ReadString('\n')

	// Simulate I/O wait
	jitter := time.Duration(rand.Intn(200)) * time.Millisecond
	time.Sleep(sleep + jitter)

	// Small response
	_, _ = fmt.Fprintln(conn, "OK")
}
