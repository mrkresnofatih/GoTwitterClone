package models

type ProfileFeedQueryModel struct {
	Username string `json:"username" validation:"required,max=30"`
	Limit    int    `json:"limit" validation:"required,max=20"`
	StartAt  int64  `json:"startAt" validation:"required"`
}

type ReplyFeedQueryModel struct {
	TweetId string `json:"tweetId" validation:"required,max=100"`
	Limit   int    `json:"limit" validation:"required,max=25"`
	StartAt int64  `json:"startAt" validation:"required"`
}
