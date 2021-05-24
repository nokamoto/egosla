package slack

import (
	"github.com/slack-go/slack/slackevents"
	"go.uber.org/zap"
)

// MessageEventHandler implements a callback for slackevents.MessageEvent.
type MessageEventHandler struct {
	logger *zap.Logger
}

// NewMessageEventHandler returns a new MessageEventHandler.
func NewMessageEventHandler(logger *zap.Logger) *MessageEventHandler {
	return &MessageEventHandler{logger: logger}
}

func (h *MessageEventHandler) Receive(ev *slackevents.MessageEvent) error {
	h.logger.Info("todo")
	return nil
}
