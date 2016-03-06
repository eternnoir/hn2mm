package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/hn2mm/h2m/payloads"
	"io/ioutil"
	"net/http"
)

func GetTopStories() ([]int64, error) {
	url := "https://hacker-news.firebaseio.com/v0/topstories.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Get top stories Fail.%s", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Get Top Stories fail. %#v", resp))
	}
	jsonStr, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Errorf("Get top stories response body fail.%s", err)
		return nil, err
	}

	var storyids []int64
	err = json.Unmarshal(jsonStr, &storyids)
	if err != nil {
		log.Errorf("Get Tos Stories dejson fail. %s", err)
		return nil, err
	}
	log.Infof("Get %d stories", len(storyids))

	return storyids, nil
}

func GetStory(id int64) (*payloads.HNItem, error) {
	url := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	log.Debugf("Try to get story by request %s.", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Get story item fail. %s", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Errorf("Get stroy iteam fail. %#v", resp)
		return nil, errors.New(fmt.Sprintf("Get story item fall, %#v", resp))
	}
	jsonStr, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var item payloads.HNItem
	err = json.Unmarshal(jsonStr, &item)
	if err != nil {
		log.Errorf("Get story item fail. %s", err)
		return nil, err
	}
	log.Infof("Get story item: %#v", item)
	return &item, nil
}
