package websockets

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"testing"
)

// 使用字典结构更容易追加和删除内容
var clients = make(map[*websocket.Conn]bool) // 保存连接的客户端
var broadcast = make(chan Message)           // 消息广播通道

// 配置 upgrader
// 其中包含获取正常HTTP连接并将其升级到WebSocket的方法
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 定义 message 结构
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// 将初始GET请求升级到websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// 确保在函数返回时关闭连接
	defer ws.Close()
	// 将新的客户端连接添加到字典中
	clients[ws] = true

	for {
		var msg Message
		// 反序列化
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("read error: %v", err)
			delete(clients, ws)
			break
		}
		// 将新接收的消息发送到广播通道
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// 从广播通道中获取消息
		msg := <-broadcast
		// 将其发送给当前连接的每个客户端
		for ws := range clients {
			err := ws.WriteJSON(msg)
			if err != nil {
				log.Printf("write error: %v", err)
				ws.Close()
				delete(clients, ws)
			}
		}
	}
}

func TestGorillaMain(t *testing.T) {
	// 创建简单的静态文件服务器
	http.Handle("/", http.FileServer(http.Dir("./public")))
	// 配置 websocket route
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}