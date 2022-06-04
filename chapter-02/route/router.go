package route

import (
    "encoding/json"
    "html/template"
    "log"
    "net/http"
    "os"

    "github.com/ylnyang/gobase/chapter-02/api"
)

type testInfo struct {
    Name string `json:"name"`
    Age  string `json:"age"`
}

func Router() {
    // 1. 返回一个json
    http.HandleFunc("/test.html", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        s := testInfo{
            Name: "haha001",
            Age:  "001",
        }
        bytes, _ := json.Marshal(s)
        _, _ = w.Write(bytes)
    })

    // 2. 前端显示
    http.HandleFunc("/one.html", one)

    // 3. api接口 只是将这个处理func剥离出来了 单独实现
    http.HandleFunc("/api/v1/post", api.Api.Two)
}

func one(w http.ResponseWriter, r *http.Request) {
    s := testInfo{
        Name: "haha002",
        Age:  "002",
    }
    // 模板标识
    t := template.New("one.html")

    // 解析文件
    path, _ := os.Getwd() // 获取当前项目的绝对路径
    t, err := t.ParseFiles(path + "/template/one.html")
    if err != nil {
        log.Panicln("解析:", err)
    }
    // 渲染
    if err := t.Execute(w, s); err != nil {
        log.Panicln("渲染:", err)
    }
}
