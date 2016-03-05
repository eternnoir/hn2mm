package payloads

import (
	"time"
)

type HNItem struct {
	Id      int64      `json:"id"`
	Deleted bool       `json:"deleted"`
	Type    *string    `json:"type"`
	By      *string    `json:"by"`
	Time    *time.Time `json:"time"`
	Text    *string    `json:"text"`
	Dead    bool       `json:"dead"`
	Parent  *int64     `json:"parent"`
	Kids    []*int64   `json:"kids"`
	Url     *string    `json:"url"`
	Score   *int64     `json:"score"`
	Title   *string    `json:"title"`
}
