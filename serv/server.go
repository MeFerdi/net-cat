package serv

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	clients    = make(map[net.Conn]string)
	messages   []string
	mu         sync.Mutex
	maxClients = 10
)

func Server(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Server listening on port %s...\n", port)

	for {
		if len(clients) >= maxClients {
			fmt.Println("Maximum number of clients reached.")
			break
		}

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	welcomeMessage(conn)
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	name := scanner.Text()

	// Validate the uniqueness of the name
	mu.Lock()
	for _, existingName := range clients {
		if existingName == name {
			conn.Write([]byte("This name is already in use. Please try again with a different name.\n"))
			conn.Close()
			mu.Unlock()
			return
		}
	}
	mu.Unlock()

	if name == "" {
		conn.Write([]byte("Name cannot be empty!\n"))
		return
	}

	mu.Lock()
	clients[conn] = name
	mu.Unlock()

	broadcast(fmt.Sprintf("[%s] has joined the chat...\n", name))
	logActivity(fmt.Sprintf("Client [%s] joined the chat.", name))

	mu.Lock()
	for _, msg := range messages {
		conn.Write([]byte(msg))
	}
	mu.Unlock()

	for {
		scanner.Scan()
		message := scanner.Text()

		if message == "" {
			break
		}

		if strings.HasPrefix(message, "/name ") {
			newName := strings.TrimSpace(strings.TrimPrefix(message, "/name "))
			if newName != "" && newName != name {
				mu.Lock()
				for _, existingName := range clients {
					if existingName == newName {
						conn.Write([]byte("This name is already taken. Choose a different one.\n"))
						mu.Unlock()
						continue
					}
				}
				oldName := name
				clients[conn] = newName
				mu.Unlock()
				broadcast(fmt.Sprintf("[%s] has changed their name to [%s]\n", oldName, newName))
				logActivity(fmt.Sprintf("Client [%s] changed their name to [%s].", oldName, newName))
				name = newName
				continue
			}
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		messageWithTime := fmt.Sprintf("[%s][%s]: %s\n", timestamp, name, message)

		mu.Lock()
		messages = append(messages, messageWithTime)
		mu.Unlock()

		broadcast(messageWithTime)
	}

	mu.Lock()
	delete(clients, conn)
	mu.Unlock()

	broadcast(fmt.Sprintf("[%s] has left the chat.\n", name))
	logActivity(fmt.Sprintf("Client [%s] left the chat.", name))
}

func broadcast(message string) {
	mu.Lock()
	defer mu.Unlock()
	for conn := range clients {
		conn.Write([]byte(message))
	}
	logActivity(message)
}

func logActivity(activity string) {
	f, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%s: %s\n", time.Now().Format(time.RFC3339), activity)); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

func welcomeMessage(conn net.Conn) {
	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	conn.Write([]byte("         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n"))
	conn.Write([]byte(`__| ".        |\dS"qML`))
	conn.Write([]byte("\n |    .       | ' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     -'       --'"))
	conn.Write([]byte("\n[ENTER YOUR NAME]: "))
}
