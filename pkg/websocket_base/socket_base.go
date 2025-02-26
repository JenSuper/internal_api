package websocket_base

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool) // 当前连接的客户端集合
var broadcast = make(chan Message)           // 广播通道

// Message 结构表示要发送的消息
type Message struct {
	Message string `json:"message"`
}

// 处理每个 WebSocket 连接
func handleConnection(w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 请求为 WebSocket 协议
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有来源连接
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 将连接加入到客户端集合中
	clients[conn] = true

	// 不断读取客户端发送的消息
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			delete(clients, conn)
			break
		}

		// 创建一个 Message 对象并发送到广播通道
		message := Message{Message: string(msg)}
		broadcast <- message
	}
}

// 广播消息给所有连接的客户端
func handleMessages() {
	for {
		// 从广播通道接收消息
		message := <-broadcast
		// 向所有客户端发送消息
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(message.Message))
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func Init() {
	// 设置路由和处理函数
	http.HandleFunc("/ws", handleConnection)

	// 启动消息广播处理器
	go handleMessages()

	// 启动 WebSocket 服务器
	fmt.Println("WebSocket server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
