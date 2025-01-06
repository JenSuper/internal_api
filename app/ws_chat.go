package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var sk = "mKtWzMoPgXoK"

// 连接池，存储所有在线用户
var clients = make(map[*websocket.Conn]string) // 连接对象 -> 用户ID
var broadcast = make(chan Message)             // 广播通道
var mutex = sync.Mutex{}                       // 保护 clients 并发操作

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message 结构体
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
	Type    string `json:"type"` // "text" or "emoji"
}

// 处理 WebSocket 连接
func handleConnections(w http.ResponseWriter, r *http.Request) {
	querySk := r.URL.Query().Get("sk")
	if querySk != sk {
		http.Error(w, "Missing sk ID", http.StatusBadRequest)
		return
	}
	// 升级 HTTP 连接为 WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// 读取用户ID
	var username string
	err = ws.ReadJSON(&username)
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}

	// 记录用户连接
	mutex.Lock()
	clients[ws] = username
	mutex.Unlock()

	fmt.Println(username, "connected!")

	// 监听用户消息
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			mutex.Lock()
			delete(clients, ws)
			mutex.Unlock()
			break
		}
		// 发送消息到广播通道
		broadcast <- msg
	}
}

// 监听广播通道，转发消息给所有在线用户
func handleMessages() {
	for {
		msg := <-broadcast

		// 发送消息给所有客户端
		mutex.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("WebSocket error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func InitSocket() {
	// 启动消息处理
	go handleMessages()

	// 设置 WebSocket 端点
	http.HandleFunc("/ws", handleConnections)

	// 启动服务器
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
