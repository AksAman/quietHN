package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type HNItem struct {
	ID          int64         `json:"id"`          // The item's unique id.
	Deleted     bool          `json:"deleted"`     // `true` if the item is deleted.
	Type        string        `json:"type"`        // One of "job", "story", "comment", "poll", or "pollopt"
	By          string        `json:"by"`          // The username of the item's author
	Time        int64         `json:"time"`        // Creation date of the item, in [Unix Time]
	Text        string        `json:"text"`        // The comment, story or poll text. HTML.
	Dead        bool          `json:"dead"`        // `true` if the item is dead.
	Kids        []int64       `json:"kids"`        // The ids of the item's comments, in ranked display order.
	URL         string        `json:"url"`         // The URL of the story.
	Score       int64         `json:"score"`       // The story's score, or the votes for a pollopt.
	Title       string        `json:"title"`       // The title of the story, poll or job. HTML.
	Descendants int64         `json:"descendants"` // In the case of stories or polls, the total comment count.
	Latency     time.Duration `json:"-"`           // Time taken to fetch the item
}

func (item *HNItem) ZeroLatency() {
	item.Latency = 0
}

func (item *HNItem) Host() string {
	parsedURL, err := url.Parse(item.URL)
	if err != nil {
		return ""
	}
	return parsedURL.Host
}

func (item *HNItem) IsStory() bool {
	return item.Type == "story" && item.URL != "" && item.Score > 100
}

func UnmarshalHNItem(data []byte) (HNItem, error) {
	var r HNItem
	err := json.Unmarshal(data, &r)
	return r, err
}

func (item *HNItem) Marshal() ([]byte, error) {
	return json.Marshal(item)
}

func (item *HNItem) String() string {
	return fmt.Sprintf("ID: %d, Type: %s, Title: %s, URL: %s", item.ID, item.Type, item.Title, item.URL)
}
