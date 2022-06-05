package controller

import (
    "sync"

    "github.com/gin-gonic/gin"
    "github.com/ylnyang/gobase/chapter-03/models"
)

type ImageController interface {
    GetAll(c *gin.Context)
    Update(c *gin.Context)
    Create(c *gin.Context)
    Delete(c *gin.Context)
}

// 该controller内部可见 
type controller struct {
    img []models.Image
}

// New 与client-go里面的controller定义使用方式类似,外部想要使用通过该方法获取controller
func New() ImageController {
    return &controller{img: make([]models.Image, 0)}
}

func (c *controller) GetAll(context *gin.Context) {
    // 返回json格式
    context.JSON(200, c.img)
}

func (c *controller) Update(context *gin.Context) {
    var image models.Image
    // 只读uri
    if err := context.ShouldBindUri(&image); err != nil {
        context.String(400, "bad request %s", err)
        return
    }
    // 读里面的body
    if err := context.ShouldBindJSON(&image); err != nil {
        context.String(400, "bad request %s", err)
        return
    }
    for idx, v := range c.img {
        if v.Id == image.Id {
            c.img[idx] = image
            context.String(200, "update image id %v success", image.Id)
            return
        }
    }
    context.String(400, "bad request cannot find image with %v to update", image.Id)
}

func (c *controller) Create(context *gin.Context) {
    // 可能出现并发操作create 需要mx 获取数据添加到切片中
    image := models.Image{Id: g.getNextId()}
    if err := context.BindJSON(&image); err != nil {
        context.String(400, "bad request %s", err)
        return
    }
    c.img = append(c.img, image)
    context.String(200, "success, new image id is %v", image.Id)
}

func (c *controller) Delete(context *gin.Context) {
    var image models.Image
    if err := context.ShouldBindUri(&image); err != nil {
        context.String(400, "bad request %s", err)
        return
    }
    for idx, value := range c.img {
        if value.Id == image.Id {
            c.img = append(c.img[:idx], c.img[idx+1:len(c.img)]...)
            context.String(200, "success, image id %v has been deleted", image.Id)
            return
        }
    }
    context.String(400, "bad request cannot find image with %v to delete", image.Id)
}

// 结构体+锁实现 添加时增加ID
type generator struct {
    counter int
    mtx     sync.Mutex
}

func (g *generator) getNextId() int {
    g.mtx.Lock()
    defer g.mtx.Unlock()

    g.counter++
    return g.counter
}

var g *generator = &generator{}
