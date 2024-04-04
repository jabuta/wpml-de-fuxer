package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type config struct {
	baseURL  string
	client   *http.Client
	postList []string
	//apiKEY  string
}

func main() {
	cfg := &config{
		baseURL:  "https://www.bizlatinhub.com",
		client:   &http.Client{},
		postList: setPostList(),
	}

	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, postSlug := range cfg.postList {
		postSlugProc := strings.Split(postSlug, "/")
		//fmt.Println(postSlugProc[1], postSlugProc[2])
		post, err := cfg.getPostBySlug(postSlugProc[1], postSlugProc[2])
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
