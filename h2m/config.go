package h2m

import (
	"time"
)

type MMConfig struct {
	WebhookUrl       string
	Channel          string
	Username         string
	CheckNewInterval duration
	DbString         string
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
