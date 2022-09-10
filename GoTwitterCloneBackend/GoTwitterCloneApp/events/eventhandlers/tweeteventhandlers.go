package eventhandlers

import (
	"context"
	"log"
	"mrkresnofatihdev/apps/gotwittercloneapp/application/tweetservice"
	"mrkresnofatihdev/apps/gotwittercloneapp/models"
)

func IncrementRetweetCountEventHandler(ctx context.Context, retweetReq models.TweetRetweetRequestModel) {
	err := tweetservice.IncrementRetweetCountOfTargetTweet(ctx, retweetReq)
	if err != nil {
		log.Println("increment retweet count event handler failed")
	}
}

func IncrementReplyCountEventHandler(ctx context.Context, replyReq models.TweetReplyRequestModel) {
	err := tweetservice.IncrementReplyCountOfTargetTweet(ctx, replyReq)
	if err != nil {
		log.Println("increment reply count event handler failed")
	}
}

func IncrementQuoteCountEventHandler(ctx context.Context, quoteReq models.TweetReplyRequestModel) {
	err := tweetservice.IncrementQuoteCountOfTargetTweet(ctx, quoteReq)
	if err != nil {
		log.Println("increment quote count event handler failed")
	}
}

func RecordReplyTweetEventHandler(ctx context.Context, replyTweet models.Tweet) {
	tweetservice.RecordReplyTweet(ctx, replyTweet)
}
