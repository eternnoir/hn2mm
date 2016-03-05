package payloads

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestToJson(t *testing.T) {
	asst := assert.New(t)
	text := "Test"
	mmp := &MatterMostPayload{Text: &text}
	blob, err := mmp.Serialize()
	if err != nil {
		asst.Fail("ToJson Fail.")
	}
	jsonStr := string(blob[:])
	asst.False(strings.Contains(jsonStr, "channel"))
	asst.True(strings.Contains(jsonStr, "text"))
}
