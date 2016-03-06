package h2m

import (
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/hn2mm/h2m/payloads"
	"github.com/eternnoir/hn2mm/h2m/utils"
	"github.com/eternnoir/hn2mm/h2m/utils/db"
	"time"
)

type H2m struct {
	targets         []*TargetChannel
	Interval        time.Duration
	checker         utils.NewStoryChecker
	topStoryChannel chan int64
	PostedStories   []int64
}

type TargetChannel struct {
	WebhookUrl string
	Channel    *string
	Username   *string
}

func NewH2mByConfig(config *MMConfig) (*H2m, error) {
	target := new(TargetChannel)
	if config.Channel != "" {
		target.Channel = &config.Channel
	}
	if config.Username != "" {
		target.Username = &config.Username
	}
	target.WebhookUrl = config.WebhookUrl
	return NewH2m([]*TargetChannel{target}, config.CheckNewInterval.Duration, config.DbString)
}

func NewH2m(targets []*TargetChannel, interval time.Duration, dbc string) (*H2m, error) {
	if len(targets) < 1 {
		return nil, errors.New("Target channel cannot be empty.")
	}

	if interval < 1 {
		return nil, errors.New("Interval seconds must over 1 sec.")
	}

	checker, err := InitChecker(dbc)
	if err != nil {
		return nil, err
	}
	h2m := &H2m{targets: targets, Interval: interval, checker: checker, topStoryChannel: make(chan int64, 100), PostedStories: []int64{}}

	return h2m, nil
}

func InitChecker(dbc string) (utils.NewStoryChecker, error) {
	// Only support sqlite3 now.
	return db.NewSqliteChecker(dbc)
}

func (h2m *H2m) Start() {
	go h2m.processNewTopStories()
	for {
		h2m.GetAndCheckTopStory()
		time.Sleep(h2m.Interval)
	}
}

func (h2m *H2m) GetAndCheckTopStory() {
	topStoryIds, err := utils.GetTopStories()
	if err != nil {
		log.Errorf("Get top story ids error. %s", err)
	}
	for _, id := range topStoryIds {
		isnew, err := h2m.checker.IsNewSotry(id)
		if err != nil {
			log.Errorf("Check story id error. %s", err)
		}
		if isnew {
			err = h2m.checker.AddPostedStory(id)
			if err != nil {
				log.Errorf("Add new story id error. %s", err)
			}
			log.Infof("Get new story id.%d", id)
			h2m.topStoryChannel <- id
		}
	}
}

func (h2m *H2m) processNewTopStories() {
	for {
		newStoryId := <-h2m.topStoryChannel
		item, err := utils.GetStory(newStoryId)
		if err != nil {
			log.Errorf("Get story item fail. %s", err)
			continue
		}
		log.Infof("Get new story item.%#v", item)
		go h2m.fireToMatterMost(item)
	}
}

func (h2m *H2m) fireToMatterMost(story *payloads.HNItem) {
	for _, target := range h2m.targets {
		mmpayload := &payloads.MatterMostPayload{Channel: target.Channel, Username: target.Username}
		text := fmt.Sprintf("New Hakcer News: [%s](%s) by %s", story.Title, story.Url, story.By)
		log.Debugf("Send payload to mattermost.%s", text)
		mmpayload.Text = &text
		json, err := mmpayload.Serialize()
		if err != nil {
			log.Errorf("Send to targer: %#v payload %#v tojson fail.%s", target, mmpayload, err)
			continue
		}
		log.Debugf("Send to targer: %#v payload %s", target, string(json))
		_, err = utils.SendRequest(target.WebhookUrl, string(json))
		if err != nil {
			log.Errorf("Send to targer: %#v payload %#v fail. Error: %s.", target, mmpayload, err)
			continue
		}
	}
}
