package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Ошибка обновления соединения:", err)
        return
    }
    defer conn.Close()

    for {
        var msg string
        err := conn.ReadJSON(&msg)
        if err != nil {
            fmt.Println("Ошибка чтения:", err)
            break
        }
        fmt.Println("Получено сообщение:", msg) 

        response := fmt.Sprintf("Вы сказали: %s", msg)
        if err := conn.WriteJSON(response); err != nil {
            fmt.Println("Ошибка отправки:", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleConnections)
    fmt.Println("Сервер слушает на порту :8080...")
    http.ListenAndServe(":8080", nil)
}