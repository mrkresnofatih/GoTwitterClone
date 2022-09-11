package models

type Tweet struct {
	TweetId      string    `json:"tweetId" firestore:"tweetId,omitempty"`
	Username     string    `json:"username" firestore:"username,omitempty"`
	Message      string    `json:"message" firestore:"message,omitempty"`
	ImageURL     string    `json:"imageURL" firestore:"imageURL,omitempty"`
	TweetType    TweetType `json:"tweetType" firestore:"tweetType,omitempty"`
	ParentTweet  *Tweet    `json:"parentTweet" firestore:"parentTweet,omitempty"`
	CreatedAt    string    `json:"createdAt" firestore:"createdAt,omitempty"`
	AvatarURL    string    `json:"avatarURL" firestore:"avatarURL,omitempty"`
	ReplyCount   int16     `json:"replyCount" firestore:"replyCount"`
	RetweetCount int16     `json:"retweetCount" firestore:"retweetCount"`
	QuoteCount   int16     `json:"quoteCount" firestore:"quoteCount"`
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
	TweetId string `json:"tweetId" validate:"required,max=100"`
}

type TweetRetweetActorModel struct {
	TweetId            string `json:"tweetId" firestore:"tweetId,omitempty"`
	ActorUsername      string `json:"actorUsername" firestore:"actorUsername,omitempty"`
	TweetOwnerUsername string `json:"tweetOwnerUsername" firestore:"tweetOwnerUsername,omitempty"`
}

type TweetReplyRecordModel struct {
	UsernameStartsWith string            `json:"usernameStartsWith" firestore:"usernameStartsWith,omitempty"`
	Replies            map[string]string `json:"replies" firestore:"replies"`
}
