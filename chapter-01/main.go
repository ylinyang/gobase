package main

import (
    "fmt"
    "time"

    "github.com/BurntSushi/toml"
)

//总的配置信息
type tomlConfig struct {
    Title   string
    Owner   owner
    DB      database `toml:"database"`
    Servers map[string]server
    Clients clients
}

//用户信息
type owner struct {
    Name string
    Org  string `toml:"organization"`
    Bio  string
    DOB  time.Time
}

//数据库相关信息
type database struct {
    Server  string
    Ports   []int
    ConnMax int `toml:"connection_max"`
    Enabled bool
}

//服务器信息
type server struct {
    IP string
    DC string
}

//客户端信息
type clients struct {
    Data  [][]interface{}
    Hosts []string
}

func main() {
    //设置结构体对象
    var config tomlConfig
    //对结构体对象填充信息，文件信息在example.toml上
    if _, err := toml.DecodeFile("./config.toml", &config); err != nil {
        fmt.Println(err)
        return
    }
    //输出标题，全局信息
    fmt.Printf("Title: %s\n", config.Title)
    //输出owner相关信息
    fmt.Printf("Owner: %s (%s, %s), Born: %s\n",
        config.Owner.Name,
        config.Owner.Org,
        config.Owner.Bio,
        config.Owner.DOB)
    //输出数据库相关信息
    fmt.Printf("Database: %s %v (Max conn. %d), Enabled? %v\n",
        config.DB.Server,
        config.DB.Ports,
        config.DB.ConnMax,
        config.DB.Enabled)
    //输出服务器相关信息
    for serverName, server := range config.Servers {
        fmt.Printf("Server: %s (%s, %s)\n",
            serverName,
            server.IP,
            server.DC)
    }
    //输出客户端相关信息
    fmt.Printf("Client data: %v\n", config.Clients.Data)
    fmt.Printf("Client hosts: %v\n", config.Clients.Hosts)
}
