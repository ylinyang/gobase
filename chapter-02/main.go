package main

import (
    "log"
    "net/http"

    "github.com/ylnyang/gobase/chapter-02/route"
)

func main() {
    server := http.Server{Addr: ":9099"}

    http.HandleFunc("/index.html", index)

    // 路由
    route.Router()

    // 启动服务
    if err := server.ListenAndServe(); err != nil {
        log.Panicln(err)
    }
}

func index(w http.ResponseWriter, r *http.Request) {
    s := "这是一个首页"
    w.Write([]byte(s))
}
