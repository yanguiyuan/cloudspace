package user

type RegisterReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type GetUsersReq struct {
	Offset int32 `json:"offset,omitempty" path:"offset"`
	Limit  int32 `json:"limit,omitempty" path:"limit"`
}
