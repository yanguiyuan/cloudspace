package file

type DeleteFileItemReq struct {
	ID       string `json:"ID,omitempty" path:"id"`
	FileName string `json:"fileName,omitempty" path:"name"`
}
type RenameReq struct {
	ID      string `json:"ID,omitempty"`
	NewName string `json:"newName,omitempty"`
}
