package models

type FollowerList struct {
	Username     string          `json:"username" firestore:"username,omitempty"`
	StartsWith   string          `json:"startsWith" firestore:"startsWith,omitempty"`
	FollowerList map[string]bool `json:"followerList" firestore:"followerList"`
}

type FollowingList struct {
	Username      string          `json:"username" firestore:"username,omitempty"`
	StartsWith    string          `json:"startsWith" firestore:"startsWith,omitempty"`
	FollowingList map[string]bool `json:"followingList" firestore:"followingList"`
}

type FollowRequestModel struct {
	FollowerUsername string
	Username string
}
