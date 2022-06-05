package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/ylnyang/gobase/chapter-03/controller"
)

func main() {
    server := gin.Default()

    // 1. /ping
    server.GET("/ping", func(c *gin.Context) {
        c.String(200, "%s", "pong")
    })

    // 2. 静态资源文件，文件存放在resource
    server.Static("/resources", "./resources")         // 文件 访问时：/resources/back.png
    server.StaticFile("/file", "./resources/back.png") // 文件夹 访问时：/file

    // 3. 通过从controller来实现的
    imgController := controller.New()
    /*
       server.GET("/image", imgController.GetAll)
       server.PUT("/image/:id", imgController.Update) // id作为变量传进来
       server.POST("/image", imgController.Create)
       server.DELETE("/image:id", imgController.Delete)
    */

    // 如上写法 都有一个共同的/image，可以通过groups来简化代码
    imagGroup := server.Group("/image")
    // get image
    imagGroup.GET("/", imgController.GetAll)

    // put image
    imagGroup.PUT("/:id", imgController.Update) // id作为变量传进来

    // create image
    imagGroup.POST("/", imgController.Create)

    // delete image
    imagGroup.DELETE("/:id", imgController.Delete)

    // 启动服务
    if err := server.Run(":9097"); err != nil {
        log.Panicln("gin启动：", err)
    }
}
