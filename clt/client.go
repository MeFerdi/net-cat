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

	fmt.Print("[ENTER YOUR NAME]: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	if name == "" {
		fmt.Println("Name cannot be empty!")
		os.Exit(1)
	}

	conn.Write([]byte(name + "\n"))

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

	for {
		scanner.Scan()
		message := scanner.Text()
		if message != "" {
			conn.Write([]byte(message + "\n"))
		}
	}
}
