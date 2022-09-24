package models

type Player struct {
	Username  string `json:"username" firestore:"username,omitempty"`
	Email     string `json:"email" firestore:"email,omitempty"`
	FullName  string `json:"fullName" firestore:"fullName,omitempty"`
	AvatarURL string `json:"avatarURL" firestore:"avatarURL,omitempty"`
	Bio       string `json:"bio" firestore:"bio,omitempty"`
	CreatedAt int64  `json:"createdAt" firestore:"createdAt,omitempty"`
}

func (s *Player) ToCreateResponseModel() PlayerCreateResponseModel {
	return PlayerCreateResponseModel{
		Username:  s.Username,
		FullName:  s.FullName,
		AvatarURL: s.AvatarURL,
		Bio:       s.Bio,
		CreatedAt: s.CreatedAt,
	}
}

func (s *Player) ToGetMinimumProfileResponseModel() PlayerGetMinimumProfileResponseModel {
	return PlayerGetMinimumProfileResponseModel{
		Username:  s.Username,
		FullName:  s.FullName,
		AvatarURL: s.AvatarURL,
		Bio:       s.Bio,
		CreatedAt: s.CreatedAt,
	}
}

type PlayerCredentials struct {
	Username string `json:"username" firestore:"username,omitempty"`
	Email    string `json:"email" firestore:"email,omitempty"`
	Password string `json:"password" firestore:"password,omitempty"`
}

type PlayerSocialStats struct {
	Username        string `json:"username" firestore:"username,omitempty"`
	NumOfFollowers  int64  `json:"numOfFollowers" firestore:"numOfFollowers"`
	NumOfFollowings int64  `json:"numOfFollowings" firestore:"numOfFollowings"`
}

type PlayerUpdateSocialStatsType string

const (
	IncrementFollowerUpdateSocialStatsType  PlayerUpdateSocialStatsType = "INCREMENT_FOLLOWER"
	DecrementFollowerUpdateSocialStatsType  PlayerUpdateSocialStatsType = "DECREMENT_FOLLOWER"
	IncrementFollowingUpdateSocialStatsType PlayerUpdateSocialStatsType = "INCREMENT_FOLLOWING"
	DecrementFollowingUpdateSocialStatsType PlayerUpdateSocialStatsType = "DECREMENT_FOLLOWING"
)

type PlayerUpdateSocialStatsRequestModel struct {
	Username   string
	UpdateType PlayerUpdateSocialStatsType
}

type PlayerCreateRequestModel struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,alphanum,min=6,max=20"`
	FullName string `json:"fullName" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,containsany=abscdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=1234567890,containsany=!@#$%^&*(),min=6,max=20"`
}

type PlayerCreateResponseModel struct {
	Username  string `json:"username"`
	FullName  string `json:"fullName"`
	AvatarURL string `json:"avatarURL"`
	Bio       string `json:"bio"`
	CreatedAt int64  `json:"createdAt"`
}

type PlayerGetMinimumProfileResponseModel struct {
	Username  string `json:"username"`
	FullName  string `json:"fullName"`
	AvatarURL string `json:"avatarURL"`
	Bio       string `json:"bio"`
	CreatedAt int64  `json:"createdAt"`
}

type PlayerLoginRequestModel struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type PlayerLoginResponseModel struct {
	Token   string                               `json:"token"`
	Profile PlayerGetMinimumProfileResponseModel `json:"profile"`
}
