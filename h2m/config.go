package h2m

import (
	"time"
)

type MMConfig struct {
	WebhookUrl       string
	Channel          string
	Username         string
	CheckNewInterval time.Duration
}
