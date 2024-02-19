package user

type RegisterReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type GetUsersReq struct {
	Offset int32 `json:"offset,omitempty" path:"offset"`
	Limit  int32 `json:"limit,omitempty" path:"limit"`
}

type ResetPasswordReq struct {
	ID          int64  `json:"id,omitempty" path:"id"`
	OldPassword string `json:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}
type AdminResetPasswordReq struct {
	ID          int64  `json:"id,omitempty" path:"id"`
	NewPassword string `json:"newPassword,omitempty"`
}
