package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	defer conn.Close()

	fmt.Println("Connected to server. Type messages and press ENTER (Ctrl+C to quit).")

	go readServer(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text() + "\n"
		conn.Write([]byte(text))
	}
}

func readServer(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Server disconnected.")
			os.Exit(0)
		}
		fmt.Print(message)
	}
}
