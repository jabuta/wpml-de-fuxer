package main

import (
	"os"
	"strings"
)

func setPostList() []string {
	file, err := os.ReadFile("url_list.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}
