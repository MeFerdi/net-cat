package clt

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client(host, port string) {
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Welcome to TCP-Chat!\n")
	fmt.Printf(" |    `.       | `' \\Zq\n     `-'       `--'")
	fmt.Print("[ENTER YOUR NAME]: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	// Validate non-empty name
	if name == "" {
		fmt.Println("Name cannot be empty!")
		os.Exit(1)
	}

	// Send name to server
	conn.Write([]byte(name + "\n"))

	// Receive previous messages
	go func() {
		for {
			message := make([]byte, 1024)
			_, err := conn.Read(message)
			if err != nil {
				fmt.Println("Error reading from server:", err)
				break
			}
			fmt.Print(string(message))
		}
	}()

	// Send messages to server
	for {
		fmt.Print("Enter message: ")
		scanner.Scan()
		message := scanner.Text()
		if message != "" {
			conn.Write([]byte(message + "\n"))
		}
	}
}
