package main

import (
	"os"

	"github.com/nokamoto/egosla/internal/cmd"
	slackhandler "github.com/nokamoto/egosla/internal/slack"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"go.uber.org/zap"
)

const (
	slackAppToken = "SLACK_APP_TOKEN"
	slackBotToken = "SLACK_BOT_TOKEN"
)

type loggerWrapper struct {
	logger *zap.Logger
	client string
}

func (l loggerWrapper) Output(i int, s string) error {
	l.logger.Debug("slack", zap.Int("i", i), zap.String("s", s), zap.String("client", l.client))
	return nil
}

func main() {
	logger := cmd.NewLogger(cmd.GetenvDebug())
	defer logger.Sync()

	appToken := os.Getenv(slackAppToken)
	botToken := os.Getenv(slackBotToken)

	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionLog(loggerWrapper{logger: logger, client: "api"}),
		slack.OptionAppLevelToken(appToken),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(loggerWrapper{logger: logger, client: "socketmode"}),
	)

	handler := slackhandler.NewMessageEventHandler(logger)

	go func() {
		for evt := range client.Events {
			switch evt.Type {
			case socketmode.EventTypeConnecting:
				logger.Info("connecting to slack with socketmode")

			case socketmode.EventTypeConnectionError:
				logger.Info("connection failed")

			case socketmode.EventTypeConnected:
				logger.Info("connected to slack with socketmode")

			case socketmode.EventTypeEventsAPI:
				eventsAPIEvent, ok := evt.Data.(slackevents.EventsAPIEvent)
				if !ok {
					client.Debugf("unsupported data")
					continue
				}

				client.Ack(*evt.Request)

				switch eventsAPIEvent.Type {
				case slackevents.CallbackEvent:
					innerEvent := eventsAPIEvent.InnerEvent
					switch ev := innerEvent.Data.(type) {
					case *slackevents.MessageEvent:
						err := handler.Receive(ev)
						if err != nil {
							logger.Error("failed to handle slackevents.MessageEvent", zap.Error(err), zap.Any("event", ev))
						}

					default:
						logger.Debug("ignore innerEvent.Data.type", zap.String("type", innerEvent.Type))
					}
				default:
					client.Debugf("unsupported Events API event received")
				}

			default:
				logger.Debug("ignore evt.Type", zap.String("type", string(evt.Type)))
			}
		}
	}()

	client.Run()
}
