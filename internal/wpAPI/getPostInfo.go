package wpAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (client Client) getPostByID(id int) (SinglePost, error) {
	endpoint := fmt.Sprintf("%s/wp-json/wp/v2/posts/%v", client.baseURL, id)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return SinglePost{}, err
	}
	setNoChacheHeaders(req)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return SinglePost{}, err
	}
	posts := Posts{}
	if err := json.NewDecoder(res.Body).Decode(&posts); err != nil {
		return SinglePost{}, err
	}
	return posts[0], nil
}

func (client Client) GetPostBySlug(lang, slug string) (SinglePost, error) {
	endpoint := fmt.Sprintf("%s/%s/wp-json/wp/v2/posts?slug=%s", client.baseURL, lang, slug)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return SinglePost{}, err
	}
	setNoChacheHeaders(req)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return SinglePost{}, err
	}
	posts := Posts{}
	if err := json.NewDecoder(res.Body).Decode(&posts); err != nil {
		return SinglePost{}, err
	}
	if len(posts) == 0 {
		return SinglePost{}, errors.New("no posts fetched")
	}
	return posts[0], nil
}

type Posts []SinglePost

type SinglePost struct {
	ID      int    `json:"id"`
	Date    string `json:"date"`
	DateGmt string `json:"date_gmt"`
	Guid    struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	Modified    string `json:"modified"`
	ModifiedGmt string `json:"modified_gmt"`
	Slug        string `json:"slug"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Link        string `json:"link"`
	Title       struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered  string `json:"rendered"`
		Protected bool   `json:"protected"`
	} `json:"content"`
	Excerpt struct {
		Rendered  string `json:"rendered"`
		Protected bool   `json:"protected"`
	} `json:"excerpt"`
	Author        int    `json:"author"`
	FeaturedMedia int    `json:"featured_media"`
	CommentStatus string `json:"comment_status"`
	PingStatus    string `json:"ping_status"`
	Sticky        bool   `json:"sticky"`
	Template      string `json:"template"`
	Format        string `json:"format"`
	Meta          struct {
		ACFChanged         bool   `json:"_acf_changed"`
		StopModifiedUpdate bool   `json:"_stopmodifiedupdate"`
		ModifiedDate       string `json:"_modified_date"`
		Footnotes          string `json:"footnotes"`
	} `json:"meta"`
	Categories             []int    `json:"categories"`
	Tags                   []int    `json:"tags"`
	BlocksyMeta            []string `json:"blocksy_meta"`
	ACF                    []string `json:"acf"`
	EnglishTranslationID   int      `json:"english_translation_id"`
	EnglishTranslationSlug string   `json:"english_translation_slug"`
	Links                  struct {
		Self []struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection []struct {
			Href string `json:"href"`
		} `json:"collection"`
		About []struct {
			Href string `json:"href"`
		} `json:"about"`
		Author []struct {
			Href string `json:"href"`
		} `json:"author"`
		Replies []struct {
			Href string `json:"href"`
		} `json:"replies"`
		VersionHistory []struct {
			Href string `json:"href"`
		} `json:"version-history"`
		FeaturedMedia []struct {
			Href string `json:"href"`
		} `json:"wp:featuredmedia"`
		Attachment []struct {
			Href string `json:"href"`
		} `json:"wp:attachment"`
		Term []struct {
			Href string `json:"href"`
		} `json:"wp:term"`
		Curies []struct {
			Name      string `json:"name"`
			Href      string `json:"href"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
	} `json:"_links"`
}
