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

type HomeFeedQueryRecord struct {
	Username      string `json:"username" firestore:"username,omitempty"`
	LastQueryTime int64  `json:"lastQueryTime" firestore:"lastQueryTime,omitempty"`
}

type HomeFeedBatchTweets struct {
	CollectedAt string           `json:"collectedAt" firestore:"collectedAt,omitempty"`
	Tweets      map[string]Tweet `json:"tweets" firestore:"tweets"`
	Username    string           `json:"username" firestore:"username,omitempty"`
}

type HomeFeedQueryModel struct {
	Limit       int   `json:"limit" validation:"required,max=25"`
	CollectedAt int64 `json:"collectedAt" validation:"required"`
}
