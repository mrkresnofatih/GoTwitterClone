package controller

import (
	"github.com/gorilla/mux"
	"mrkresnofatihdev/apps/gotwittercloneapp/controller/tweethandlers"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
	"mrkresnofatihdev/apps/gotwittercloneapp/utils"
	"net/http"
)

type TweetController struct {}

func (t *TweetController) AddControllerTo(router *mux.Router) {
	subRouter := &utils.ApplicationRouter{
		Parent: router,
		PathPrefix: "/tweet",
	}

	postTweetEndpoint := &utils.ApplicationEndpoint{
		Handler: tweethandlers.PostTweetHandler,
		Path: "/post",
		Method: http.MethodPost,
	}
	postTweetEndpointWithValidation := &utils.RequireValidation[models.TweetCreateRequestModel]{
		Endpoint: postTweetEndpoint,
	}
	postTweetEndpointWithAuth := &utils.RequireAuthentication{
		Endpoint: postTweetEndpointWithValidation,
	}
	subRouter.AddEndpoint(postTweetEndpointWithAuth)

	postReplyEndpoint := &utils.ApplicationEndpoint{
		Handler: tweethandlers.PostReplyHandler,
		Path: "/reply",
		Method: http.MethodPost,
	}
	postReplyWithValidation := &utils.RequireValidation[models.TweetReplyRequestModel]{
		Endpoint: postReplyEndpoint,
	}
	postReplyWithAuth := &utils.RequireAuthentication{
		Endpoint: postReplyWithValidation,
	}
	subRouter.AddEndpoint(postReplyWithAuth)

	postRetweetEndpoint := &utils.ApplicationEndpoint{
		Handler: tweethandlers.PostRetweetHandler,
		Path: "/retweet",
		Method: http.MethodPost,
	}
	postRetweetWithValidation := &utils.RequireValidation[models.TweetRetweetRequestModel]{
		Endpoint: postRetweetEndpoint,
	}
	postRetweetWithAuth := &utils.RequireAuthentication{
		Endpoint: postRetweetWithValidation,
	}
	subRouter.AddEndpoint(postRetweetWithAuth)

	postQuoteEndpoint := &utils.ApplicationEndpoint{
		Handler: tweethandlers.PostQuoteHandler,
		Path: "/quote",
		Method: http.MethodPost,
	}
	postQuoteWithValidation := &utils.RequireValidation[models.TweetReplyRequestModel]{
		Endpoint: postQuoteEndpoint,
	}
	postQuoteWithAuth := &utils.RequireAuthentication{
		Endpoint: postQuoteWithValidation,
	}
	subRouter.AddEndpoint(postQuoteWithAuth)

	subRouter.Init()
}
