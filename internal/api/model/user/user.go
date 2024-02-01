package user

type RegisterReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
