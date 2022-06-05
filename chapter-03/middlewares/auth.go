package middlewares

import "github.com/gin-gonic/gin"

// 自定义一个中间件

func MyAuth() gin.HandlerFunc {
    return gin.BasicAuth(gin.Accounts{
        "yang": "yang",
    })
}
