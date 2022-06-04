package api

import "net/http"

var Api = &apiHandler{}

type apiHandler struct {
}

func (*apiHandler) Two(w http.ResponseWriter, r *http.Request) {
    s := "这是一个简单地api接口"
    w.Write([]byte(s))
}
