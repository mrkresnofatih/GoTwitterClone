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

type PlayerCreateRequestModel struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	FullName string `json:"fullName" validate:"required"`
	Password string `json:"password" validate:"required"`
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
