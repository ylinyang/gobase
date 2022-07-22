package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZWXYZ"
	fmt.Println(strings.Count(s, "WXYZ")) //2
	fmt.Println(len(s))
}
