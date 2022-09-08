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

	return map[string]func(delivery amqp091.Delivery){
		FollowEventHandlerName:   followEvtHandlerInstance.GetHandler(),
		UnfollowEventHandlerName: unfollowEvtHandlerInstance.GetHandler(),
	}
}

const FollowEventHandlerName = "FollowEvtHandler"
const UnfollowEventHandlerName = "UnfollowEvtHandler"
