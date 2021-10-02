package model

type Page struct {
	Rss Rss `json:"rss"`
}

type Rss struct {
	Version string  `json:"-version"`
	Channel Channel `json:"channel"`
}

type Channel struct {
	Link        string `json:"link"`
	Description string `json:"description"`
	Copyright   string `json:"copyright"`
	Language    string `json:"language"`
	PubDate     string `json:"pubDate"`
	Item        []Item `json:"item"`
}

type Item struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Guid        string `json:"guid"`
}
