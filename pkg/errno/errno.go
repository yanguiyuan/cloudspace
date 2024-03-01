package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    = 0
	ServiceErrCode = 1000 + iota
	UserAlreadyExistCode
	UsernameOrPasswordErrorCode
	NotFoundCode
	InvalidParamCode
	UserPasswordErrCode
	NoWritePermissionCode
	NoReadPermissionCode
	NoAdminPermissionCode
	NoAllPermissionCode
	FolderNotEmptyCode
)

const (
	SuccessMsg                 = "Success"
	ServiceErrMsg              = "Service is unable to start successfully"
	UserAlreadyExistMsg        = "User already exist"
	UsernameOrPasswordErrorMsg = "Username or password error"
	NotFoundMsg                = "Not found"
	InvalidParamMsg            = "Invalid param"
	UserPasswordErrMsg         = "User password error"
	NoWritePermissionMsg       = "No write permission"
	NoReadPermissionMsg        = "No read permission"
	NoAdminPermissionMsg       = "No admin permission"
	NoAllPermissionMsg         = "No all permission"
	FolderNotEmptyMsg          = "Folder is not empty"
)

var (
	Success                 = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr              = NewErrNo(ServiceErrCode, ServiceErrMsg)
	UserAlreadyExist        = NewErrNo(UserAlreadyExistCode, UserAlreadyExistMsg)
	UsernameOrPasswordError = NewErrNo(UsernameOrPasswordErrorCode, UsernameOrPasswordErrorMsg)
	NotFound                = NewErrNo(NotFoundCode, NotFoundMsg)
	InvalidParam            = NewErrNo(InvalidParamCode, InvalidParamMsg)
	UserPasswordErr         = NewErrNo(UserPasswordErrCode, UserPasswordErrMsg)
)

type ErrNo struct {
	Code int32
	Msg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("error code: %d, error message: %s", e.Code, e.Msg)
}
func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		fmt.Println("Err:", Err)
		return Err
	}
	s := ServiceErr
	s.Msg = err.Error()
	fmt.Println(" S:", s)
	return s
}
