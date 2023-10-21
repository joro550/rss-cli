package data

import (
	"errors"

	"github.com/mmcdole/gofeed"
)

var (
  InvalidFeedUrl = errors.New("Invalid feed url")
)

type Feed struct {
	Id  uint `gorm:"primaryKey"`
	Url string
  Title string 

}

func (f Feed) Validate() bool {
	fp := gofeed.NewParser()
	_, err := fp.ParseURL(f.Url)

	if err == nil {
		return true
	}
	return false
}

func FeedFromUrl(url string) (Feed, error) {
  fp := gofeed.NewParser()
  feed, err := fp.ParseURL(url)
  if err != nil {
    return Feed{}, InvalidFeedUrl
  }

  rssFeed := Feed{Url: url, Title: feed.Title}
  return rssFeed, nil
}
