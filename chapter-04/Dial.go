package main

import (
    "fmt"
)

type dailOptions struct {
    //dial     func(network string, address string) (net.Conn, error)
    network  string
    address  string
    name     string
    password string
}

type DialOption func(*dailOptions)

func Dial(network, address string, options ...DialOption) {
    do := dailOptions{
        network:  network,
        address:  address,
        name:     "haha001", // default value
        password: "haha001",
    }
    for _, option := range options {
        option(&do)
    }
    fmt.Println(do)
}
