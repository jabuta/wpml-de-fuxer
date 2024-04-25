package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, postSlug := range cfg.postList {
		postSlugProc := strings.Split(postSlug, "/")
		//fmt.Println(postSlugProc[1], postSlugProc[2])
		post, err := cfg.client.GetPostBySlug(postSlugProc[1], postSlugProc[2])
		if err != nil {
			fmt.Fprintf(file, "%s	%s\n", postSlug, err)
			continue
		}
		if post.EnglishTranslationID == 0 || post.EnglishTranslationID == post.ID {
			fmt.Fprintf(file, "%s	%v	No Translation\n", postSlug, post.ID)
			continue
		}
		fmt.Fprintf(file, "%s	%v	%s	%v\n", postSlug, post.ID, post.EnglishTranslationSlug, post.EnglishTranslationID)
	}
}
