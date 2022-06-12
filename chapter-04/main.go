package main

func main() {
    // 1.
    Dial("tcp", "127.0.0.1:2") // {tcp 127.0.0.1:2 haha001 haha001}

    // 2.
    Dial("tcp", "127.0.0.1:1", func(options *dailOptions) {
        options.name = "xxx"
    }) //{tcp 127.0.0.1:1 xxx haha001}
}
