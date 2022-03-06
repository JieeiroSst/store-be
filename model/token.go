package model

type Token struct {
	Id       string
	Username string
	RoleId   string
}

type ResultToken struct {
	Token string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}
