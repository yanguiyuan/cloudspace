package file

type RenameReq struct {
	ID      string `json:"ID,omitempty"`
	NewName string `json:"newName,omitempty"`
}
type ModifyFileContentReq struct {
	ID      string `json:"ID,omitempty" path:"id"`
	Content string `json:"content,omitempty"`
}
