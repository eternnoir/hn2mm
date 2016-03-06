package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTopStoriesIds(t *testing.T) {
	asst := assert.New(t)
	ids, err := GetTopStories()
	if err != nil {
		asst.Fail(fmt.Sprint(err))
	}
	asst.True(len(ids) > 0)
}

func TestGetStoryItem(t *testing.T) {
	asst := assert.New(t)
	item, err := GetStory(int64(8863))
	if err != nil {
		asst.Fail(fmt.Sprintf("Get story item fail.%s", err))
		return
	}
	asst.Equal(int64(8863), item.Id)
	asst.Equal("http://www.getdropbox.com/u/2/screencast.html", item.Url)
}
