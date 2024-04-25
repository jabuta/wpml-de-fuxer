package main

import (
	"flag"
	"log"
	"time"

	"github.com/jabuta/wpml-de-fuxer/internal/wpAPI"
)

func main() {
	siteURL := flag.String("url", "noURL", "write out the full wordpress url")
	flag.Parse()
	if *siteURL == "noURL" {
		log.Panic("input a url")
	}
	cfg := &config{
		client:   wpAPI.NewClient(5*time.Second, *siteURL),
		postList: setPostList(),
	}

	startREPL(cfg)

}
