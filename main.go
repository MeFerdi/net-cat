package main

import (
	"fmt"
	"os"

	"net-cat/clt"
	"net-cat/serv"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(1)
	}

	port := ""
	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) == 1 {
		port = "8989"
	}

	mode := "server"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	if len(os.Args) > 2 && os.Args[2] == "client" {
		mode = "client"
	}

	switch mode {
	case "client":
		if len(os.Args) != 4 {
			fmt.Println("[USAGE]: ./TCPChat $port $host")
			os.Exit(1)
		}
		host := os.Args[2]
		clt.Client(host, port)
	case "server":
		if len(os.Args) > 2 {
			fmt.Println("[USAGE]: ./TCPChat $port")
			return
		}
		serv.Server(port)
	default:
		fmt.Println("[USAGE]: ./TCPChat $port")
	}
}
