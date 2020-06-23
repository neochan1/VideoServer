package models
//requests
type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//response
type Signedup struct {
	Success bool `json:"success"`
	SessionID string `json:"session_id"`
}
//date model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string
	TTL int64
}