package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var timeoutStr string
	flag.StringVar(&timeoutStr, "timeout", "10s", "Timeout for connection in seconds")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: go-telnet [options] host port")
		os.Exit(1)
	}

	host, port := args[0], args[1]

	duration, err := time.ParseDuration(timeoutStr)
	if err != nil {
		fmt.Println("Invalid timeout")
		os.Exit(1)
	}

	conn, err := net.DialTimeout("tcp", host+":"+port, duration)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT)
		<-signals
		fmt.Println("Ctrl+C detected, closing connection.")
		conn.Close()
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			fmt.Fprintln(conn, input)
		}
	}()

	select {}
}
