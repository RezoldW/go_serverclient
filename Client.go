// client.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "your_ip")
	if err != nil {
		fmt.Println("Ошибка подключения к серверу:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Введите сообщение (для выхода введите 'exit'):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "exit" {
			break
		}
		fmt.Fprintln(conn, message)
	}
}
