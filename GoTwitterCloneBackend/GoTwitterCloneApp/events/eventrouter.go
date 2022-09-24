package events

import (
	"github.com/rabbitmq/amqp091-go"
	"mrkresnofatihdev/apps/gotwittercloneapp/events/eventhandlers"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func GetEventHandlers() map[string]func(delivery amqp091.Delivery) {
	followEvtHandlerInstance := BaseEventHandler[models.FollowRequestModel]{
		ExecutorFunc: eventhandlers.FollowEventHandler,
	}
	unfollowEvtHandlerInstance := BaseEventHandler[models.FollowRequestModel]{
		ExecutorFunc: eventhandlers.UnfollowEventHandler,
	}
	incrementRetweetCountEvtHandlerInstance := BaseEventHandler[models.TweetRetweetRequestModel]{
		ExecutorFunc: eventhandlers.IncrementRetweetCountEventHandler,
	}
	incrementReplyCountEvtHandlerInstance := BaseEventHandler[models.TweetReplyRequestModel]{
		ExecutorFunc: eventhandlers.IncrementReplyCountEventHandler,
	}
	incrementQuoteCountEvtHandlerInstance := BaseEventHandler[models.TweetReplyRequestModel]{
		ExecutorFunc: eventhandlers.IncrementQuoteCountEventHandler,
	}
	recordReplyTweetEvtHandlerInstance := BaseEventHandler[models.Tweet]{
		ExecutorFunc: eventhandlers.RecordReplyTweetEventHandler,
	}
	runHomeFeedQueryEvtHandlerInstance := BaseEventHandler[string]{
		ExecutorFunc: eventhandlers.RunGetHomeQueryEventHandler,
	}

	return map[string]func(delivery amqp091.Delivery){
		FollowEventHandlerName:                followEvtHandlerInstance.GetHandler(),
		UnfollowEventHandlerName:              unfollowEvtHandlerInstance.GetHandler(),
		IncrementRetweetCountEventHandlerName: incrementRetweetCountEvtHandlerInstance.GetHandler(),
		IncrementReplyCountEventHandlerName:   incrementReplyCountEvtHandlerInstance.GetHandler(),
		IncrementQuoteCountEventHandlerName:   incrementQuoteCountEvtHandlerInstance.GetHandler(),
		RecordReplyTweetEventHandlerName:      recordReplyTweetEvtHandlerInstance.GetHandler(),
		RunHomeFeedQueryEventHandlerName:      runHomeFeedQueryEvtHandlerInstance.GetHandler(),
	}
}

const FollowEventHandlerName = "FollowEvtHandler"
const UnfollowEventHandlerName = "UnfollowEvtHandler"

const IncrementRetweetCountEventHandlerName = "IncrementRetweetCountEvtHandler"
const IncrementReplyCountEventHandlerName = "IncrementReplyCountEvtHandler"
const IncrementQuoteCountEventHandlerName = "IncrementQuoteCountEvtHandler"
const RecordReplyTweetEventHandlerName = "RecordReplyTweetEvtHandler"

const RunHomeFeedQueryEventHandlerName = "RunHomeFeedQueryEvtHandlerName"
