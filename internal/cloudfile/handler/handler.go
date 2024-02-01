package handler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/dal"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/model"
	rpc "github.com/yanguiyuan/cloudspace/pkg/rpc"
	"github.com/yanguiyuan/yuan/pkg/gen/id"
	"path/filepath"
	"strings"
)

func LinkPath(userID int64, id, fileType string) string {
	return fmt.Sprintf("cloud-file/user/%d/%d.%s", userID, id, fileType)
}
func ParseFileType(filename string) string {
	ext := filepath.Ext(filename)
	if len(ext) == 0 {
		return ""
	}
	return ext[1:]
}

// CloudFileServiceImpl implements the last service interface defined in the IDL.
type CloudFileServiceImpl struct {
	OssBucket *oss.Bucket
}

const (
	Directory   = "directory"
	Namespace   = "namespace"
	Markdown    = "md"
	Text        = "txt"
	GoSource    = "go"
	CPPSource   = "cpp"
	Yaml        = "yml"
	Thrift      = "thrift"
	SqlScript   = "sql"
	ShellScript = "sh"
)

// Add implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) Add(ctx context.Context, req *rpc.AddRequest) (err error) {
	id1 := id.Base62()
	if err != nil {
		return err
	}
	dal.Q.Transaction(func(tx *dal.Query) error {
		fileType := ParseFileType(req.FileName)
		err := tx.FileItem.WithContext(ctx).Create(&model.FileItem{
			ID:   id1,
			Name: req.FileName,
			Type: fileType,
		})
		if err != nil {
			return err
		}
		err = tx.FileIndex.WithContext(ctx).Create(&model.FileIndex{
			ParentID: req.ParentID,
			ChildID:  id1,
		})
		if err != nil {
			return err
		}
		err = s.OssBucket.PutObject(LinkPath(req.Uid, id1, fileType), bytes.NewReader(req.FileData))
		if err != nil {
			return err
		}
		return nil
	})
	return
}

// Remove implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) Remove(ctx context.Context, req *rpc.RemoveRequest) (err error) {
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		_, err := tx.FileItem.WithContext(ctx).Where(dal.FileItem.ID.Eq(req.Id)).Delete()
		if err != nil {
			return err
		}
		_, err = tx.FileIndex.WithContext(ctx).Where(dal.FileIndex.ChildID.Eq(req.Id)).Delete()
		if err != nil {
			return err
		}
		return s.OssBucket.DeleteObject(LinkPath(req.Uid, req.Id, ParseFileType(req.Filename)))
	})
	if err != nil {
		return err
	}
	return
}

// Query implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) Query(ctx context.Context, pid string, uid int64) (resp *rpc.QueryResponse, err error) {
	find, err := dal.FileIndex.WithContext(ctx).Where(dal.FileIndex.ParentID.Eq(pid)).Find()
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, index := range find {
		ids = append(ids, index.ChildID)
	}
	items, err := dal.FileItem.WithContext(ctx).Where(dal.FileItem.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	var res []*rpc.CloudFileItem
	var urls []*rpc.UrlTable
	for _, item := range items {
		url, err := s.OssBucket.SignURL(LinkPath(uid, item.ID, item.Type), oss.HTTPGet, 60*60*24)
		if err != nil {
			return nil, err
		}
		urls = append(urls, &rpc.UrlTable{Url: url, Id: item.ID})
		res = append(res, &rpc.CloudFileItem{
			Id:         item.ID,
			FileName:   item.Name,
			FileType:   item.Type,
			UpdateTime: item.UpdateTime.String(),
			CreateTime: item.CreateTime.String(),
		})
	}
	resp.Items = res
	resp.Urlmaps = urls
	return resp, nil
}

// Update implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) Update(ctx context.Context, req *rpc.UpdateRequest) (err error) {
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		item, err := tx.FileItem.WithContext(ctx).Where(dal.FileItem.ID.Eq(req.Id)).First()
		if err != nil {
			return err
		}
		err = s.OssBucket.PutObject(LinkPath(req.Uid, req.Id, item.Type), strings.NewReader(req.Content))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// CreateDirectory implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) CreateDirectory(ctx context.Context, req *rpc.CreateDirectoryRequest) (err error) {
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		id0 := id.Base62()
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		err = tx.FileItem.WithContext(ctx).Create(&model.FileItem{
			ID:   id0,
			Name: req.DirectoryName,
			Type: Directory,
		})
		if err != nil {
			return err
		}
		err = tx.FileIndex.WithContext(ctx).Create(&model.FileIndex{
			ParentID: req.ParentID,
			ChildID:  id0,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// RemoveDirectory implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) RemoveDirectory(ctx context.Context, id string) (err error) {
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		_, err := tx.FileIndex.WithContext(ctx).Where(dal.FileIndex.ChildID.Eq(id)).Delete()
		if err != nil {
			return err
		}
		_, err = tx.FileItem.WithContext(ctx).Where(dal.FileItem.ID.Eq(id)).Delete()
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// Rename implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) Rename(ctx context.Context, id, newName string) (err error) {
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		_, err := tx.FileItem.WithContext(ctx).Where(dal.FileItem.ID.Eq(id)).Update(dal.FileItem.Name, newName)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
func (s *CloudFileServiceImpl) QueryUserFileRoot(ctx context.Context, id int64) (r string, err error) {
	res, err := dal.Q.Namespace.WithContext(ctx).
		Select(dal.Namespace.RootID).
		LeftJoin(
			dal.UserNamespace,
			dal.UserNamespace.UserID.Eq(id),
			dal.Namespace.ID.EqCol(dal.UserNamespace.NamespaceID)).
		First()
	fmt.Println("ID:", id)
	fmt.Println("Result:", res)
	fmt.Println("err:", err)
	if err != nil {
		return "", err
	}
	return res.RootID, nil
}

func (s *CloudFileServiceImpl) CreateFileItem(ctx context.Context, name string, ty string, parentID string) (r string, err error) {
	id := id.Base62()
	switch ty {
	case "namespace":
		id = "N-" + id
	case "directory":
		id = "D-" + id
	default:
		id = "F-" + id
	}
	var fileItem *model.FileItem
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		fileItem, err = tx.FileItem.WithContext(ctx).Where(tx.FileItem.ID.Eq(id), tx.FileItem.Name.Eq(name), tx.FileItem.Type.Eq(ty)).FirstOrCreate()
		if err != nil {
			return err
		}
		err = tx.FileIndex.WithContext(ctx).Create(&model.FileIndex{ParentID: parentID, ChildID: fileItem.ID})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return fileItem.ID, nil
}
func (s *CloudFileServiceImpl) CreateNamespace(ctx context.Context, name string, rootID string) (r int64, err error) {
	namespace, err := dal.Namespace.WithContext(ctx).Where(dal.Namespace.Name.Eq(name), dal.Namespace.RootID.Eq(rootID)).FirstOrCreate()
	if err != nil {
		return 0, err
	}
	return namespace.ID, nil
}
func (s *CloudFileServiceImpl) CreateUserNamespace(ctx context.Context, userID int64, namespaceID int64, authority int32) (err error) {
	return dal.UserNamespace.WithContext(ctx).Create(&model.UserNamespace{
		UserID:      userID,
		NamespaceID: namespaceID,
		Authority:   authority,
	})
}
func (s *CloudFileServiceImpl) QueryFileItemByID(ctx context.Context, id string) (resp *rpc.CloudFileItem, err error) {
	res, err := dal.FileItem.WithContext(ctx).Where(dal.FileItem.ID.Eq(id)).First()
	return &rpc.CloudFileItem{Id: res.ID, FileName: res.Name, FileType: res.Type, UpdateTime: res.UpdateTime.String(), CreateTime: res.CreateTime.String()}, err
}
