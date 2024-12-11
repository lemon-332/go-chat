package main

// Hub 负责管理所有客户端的注册、注销及消息广播
type Hub struct {
    clients    map[*Client]bool // 已连接的客户端
    broadcast  chan []byte      // 从客户端接收的广播消息
    register   chan *Client     // 注册请求
    unregister chan *Client     // 注销请求
}

func newHub() *Hub {
    return &Hub{
        clients:    make(map[*Client]bool),
        broadcast:  make(chan []byte),
        register:   make(chan *Client),
        unregister: make(chan *Client),
    }
}

func (h *Hub) run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
        case client := <-h.unregister:
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
        case message := <-h.broadcast:
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
        }
    }
}
