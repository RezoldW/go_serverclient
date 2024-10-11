package main

import (
	"fmt"
	"net"
)

func main() {
	// Запуск сервера на порту 8080
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен на порту 8080")

	// Основной цикл для принятия соединений
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения:", err)
			continue
		}
		go handleConnection(conn) // Обрабтка соединений в отдельной горутине
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close() // Закрываем соединение по завершению
	fmt.Println("Новое соединение:", conn.RemoteAddr())

	buffer := make([]byte, 1024) // Буфер для чтения данных
	for {
		n, err := conn.Read(buffer) // Чтение данных из соединения
		if err != nil {
			fmt.Println("Ошибка при чтении:", err)
			return
		}
		fmt.Println("Получено сообщение:", string(buffer[:n])) // Вывод полученного сообщения
	}
}
