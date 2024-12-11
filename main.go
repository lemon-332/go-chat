package main

import (
    "log"
    "net/http"
)

func main() {
	hub := newHub()
    go hub.run()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(hub, w, r)
    })

    log.Println("服务器启动，监听端口 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("监听失败:", err)
    }
}
