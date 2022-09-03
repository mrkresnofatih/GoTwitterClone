package models

import "time"

type Tweet struct {
	TweetId     string    `json:"tweetId" firestore:"tweetId,omitempty"`
	Username    string    `json:"username" firestore:"username,omitempty"`
	Message     string    `json:"message" firestore:"message,omitempty"`
	ImageURL    string    `json:"imageURL" firestore:"imageURL,omitempty"`
	TweetType   TweetType `json:"tweetType" firestore:"tweetType,omitempty"`
	ParentTweet *Tweet    `json:"parentTweet" firestore:"parentTweet,omitempty"`
	CreatedAt   time.Time `json:"createdAt" firestore:"createdAt,omitempty"`
}

type TweetType string

const (
	SimpleTweetType  TweetType = "SIMPLE"
	ReplyTweetType   TweetType = "REPLY"
	RetweetTweetType TweetType = "RETWEET"
	QuoteTweetType   TweetType = "QUOTE"
)

type TweetCreateRequestModel struct {
	Message  string `json:"message" validate:"required,max=200"`
	ImageURL string `json:"imageURL" validate:"omitempty,max=300,url"`
}

type TweetReplyRequestModel struct {
	TweetId  string `json:"tweetId" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=30"`
	Message  string `json:"message" validate:"max=200,required_without=ImageURL"`
	ImageURL string `json:"imageURL" validate:"max=300,required_without=Message"`
}

type TweetRetweetRequestModel struct {
	TweetId  string `json:"tweetId" validate:"required,max=100"`
	Username string `json:"username" validate:"required,max=30"`
}

type TweetGetRequestModel struct {
	Username string `json:"username" validate:"required,max=30"`
	TweetId  string `json:"tweetId" validate:"required,max=100"`
}

type TweetRetweetActorModel struct {
	TweetId            string `json:"tweetId" firestore:"tweetId,omitempty"`
	ActorUsername      string `json:"actorUsername" firestore:"actorUsername,omitempty"`
	TweetOwnerUsername string `json:"tweetOwnerUsername" firestore:"tweetOwnerUsername,omitempty"`
}
