package model

// CodeInfo 实体
type CodeInfo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Timestamp int    `json:"timestamp"`
	Token     string `json:"token"`
}
