package main

import (
	"log"
	"net"
	"os"
	"strings"
)

var db = NewDB()

func main() {
	log.Println("logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleCommand(conn)
	}
}

func handleCommand(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 1024)
	for {
		_, err := conn.Read(b)

		if err != nil {
			log.Println("Error reading:", err.Error())
			return
		}

		log.Println("Received:", string(b))

		// Parse the RESP command
		cmd := parseRESP(b)
		log.Println("Command:", cmd)

		var res string
		switch cmd[0] {
		case "ping":
			res = pingHandler()

		case "echo":
			res = echoHandler(cmd[1])

		case "set":
			res = setHandler(cmd[1:])

		case "get":
			res = getHandler(cmd[1])

		default:
			log.Println("-ERR unknown command\r")
			break

		}

		_, err = conn.Write([]byte(res))
		if err != nil {
			log.Println("Error writing:", err.Error())
			return
		}
	}
}

func parseRESP(b []byte) []string {
	res := strings.Split(string(b), "\r\n")
	res = res[:len(res)-1]

	ret := make([]string, 0)
	for _, r := range res {
		if r[0] == '*' || r[0] == '$' {
			continue
		}

		ret = append(ret, r)
	}

	return ret
}
