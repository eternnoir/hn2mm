package payloads

import (
	"fmt"
	"strconv"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil

}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err

	}

	*t = Timestamp(time.Unix(int64(ts), 0))

	return nil

}

type HNItem struct {
	Id      int64      `json:"id"`
	Deleted bool       `json:"deleted"`
	Type    string     `json:"type"`
	By      string     `json:"by"`
	Time    *Timestamp `json:"time"`
	Text    string     `json:"text"`
	Dead    bool       `json:"dead"`
	Parent  int64      `json:"parent"`
	Kids    []int64    `json:"kids"`
	Url     string     `json:"url"`
	Score   int64      `json:"score"`
	Title   string     `json:"title"`
}
