package h2m

import (
	"errors"
)

type H2m struct {
	targets  []*TargetChannel
	Interval int
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
	return NewH2m([]*TargetChannel{target}, config.CheckNewInterval)
}

func NewH2m(targets []*TargetChannel, interval int) (*H2m, error) {
	if len(targets) < 1 {
		return nil, errors.New("Target channel cannot be empty.")
	}

	if interval < 1 {
		return nil, errors.New("Interval seconds must over 1 sec.")
	}
	h2m := &H2m{targets: targets, Interval: interval}
	return h2m, nil
}

func (h2m *H2m) Start() {

}
