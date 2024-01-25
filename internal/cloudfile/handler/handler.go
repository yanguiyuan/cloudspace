package handler

import (
	"bytes"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/dal"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/model"
	rpc "github.com/yanguiyuan/cloudspace/pkg/rpc"
	"github.com/yanguiyuan/yuan/pkg/gen/id"
	"path/filepath"
	"strings"
)

func LinkPath(username, id, fileType string) string {
	return "cloud-file/user/" + username + "/" + id + "." + fileType
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
	id2 := id.Base62()
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
			ID:       id2,
			ParentID: req.ParentID,
			ChildID:  id1,
		})
		if err != nil {
			return err
		}
		err = s.OssBucket.PutObject(LinkPath(req.Username, id1, fileType), bytes.NewReader(req.FileData))
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
		return s.OssBucket.DeleteObject(LinkPath(req.Username, req.Id, ParseFileType(req.Filename)))
	})
	if err != nil {
		return err
	}
	return
}

// Query implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) Query(ctx context.Context, req *rpc.QueryRequest) (resp *rpc.QueryResponse, err error) {
	find, err := dal.FileIndex.WithContext(ctx).Where(dal.FileIndex.ParentID.Eq(req.Pid)).Find()
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
		url, err := s.OssBucket.SignURL(LinkPath(req.Username, item.ID, item.Type), oss.HTTPGet, 60*60*24)
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
		err = s.OssBucket.PutObject(LinkPath(req.Username, req.Id, item.Type), strings.NewReader(req.Content))
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
		id1 := id.Base62()
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
			ID:       id1,
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
