package handler

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/dal"
	"github.com/yanguiyuan/cloudspace/internal/cloudfile/model"
	rpc "github.com/yanguiyuan/cloudspace/pkg/rpc"
	"github.com/yanguiyuan/yuan/pkg/gen/id"
	"io"
	"path/filepath"
	"strings"
)

func LinkPath(userID int64, id, fileType string) string {
	return fmt.Sprintf("cloud-file/user/%d/%s.%s", userID, id, fileType)
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
func (s *CloudFileServiceImpl) Add(ctx context.Context, req *rpc.AddRequest) (resp *rpc.CloudFileItem, err error) {
	id1 := "F-" + id.Base62()
	if err != nil {
		return nil, err
	}
	var f *model.FileItem
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		fileType := ParseFileType(req.FileName)
		client, err := cloudfile.NewFileServiceClient()
		if err != nil {
			return err
		}
		r, err := client.QueryFileItemByID(ctx, req.ParentID)
		if err != nil {
			return err
		}
		f, err = tx.FileItem.WithContext(ctx).Where(
			tx.FileItem.NamespaceID.Eq(r.NamespaceID),
			tx.FileItem.ID.Eq(id1),
			tx.FileItem.Name.Eq(req.FileName),
			tx.FileItem.Type.Eq(fileType)).FirstOrCreate()
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
			fmt.Println("oss err:", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &rpc.CloudFileItem{
		Id:          f.ID,
		FileName:    f.Name,
		NamespaceID: f.NamespaceID,
		FileType:    f.Type,
		CreateTime:  f.CreateTime.String(),
		UpdateTime:  f.UpdateTime.String(),
	}, nil
}

// Remove implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) Remove(ctx context.Context, id string) (err error) {
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		first, err := tx.FileItem.WithContext(ctx).Where(tx.FileItem.ID.Eq(id)).First()
		if err != nil {
			return err
		}
		_, err = tx.FileItem.WithContext(ctx).Where(dal.FileItem.ID.Eq(id)).Delete()
		if err != nil {
			return err
		}
		_, err = tx.FileIndex.WithContext(ctx).Where(dal.FileIndex.ChildID.Eq(id)).Delete()
		if err != nil {
			return err
		}
		return s.OssBucket.DeleteObject(fmt.Sprintf("cloud-file/namespace/%d/%s.%s", first.NamespaceID, first.ID, first.Type))
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
		var url string
		if item.Type != Directory && item.Type != Namespace {
			url, err = s.OssBucket.SignURL(LinkPath(uid, item.ID, item.Type), oss.HTTPGet, 60*60*24)
			if err != nil {
				return nil, err
			}
			urls = append(urls, &rpc.UrlTable{Url: url, Id: item.ID})
		}
		res = append(res, &rpc.CloudFileItem{
			Id:         item.ID,
			FileName:   item.Name,
			FileType:   item.Type,
			UpdateTime: item.UpdateTime.String(),
			CreateTime: item.CreateTime.String(),
		})
	}
	var resp0 = &rpc.QueryResponse{}
	resp0.Items = res
	resp0.Urlmaps = urls
	resp = resp0
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
func (s *CloudFileServiceImpl) CreateDirectory(ctx context.Context, req *rpc.CreateDirectoryRequest) (resp *rpc.CloudFileItem, err error) {
	var file *model.FileItem
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		id0 := "D-" + id.Base62()
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		client, err := cloudfile.NewFileServiceClient()
		if err != nil {
			return err
		}
		r, err := client.QueryFileItemByID(ctx, req.ParentID)
		if err != nil {
			return err
		}
		file, err = tx.FileItem.WithContext(ctx).Where(
			tx.FileItem.NamespaceID.Eq(r.NamespaceID),
			tx.FileItem.ID.Eq(id0),
			tx.FileItem.Name.Eq(req.DirectoryName),
			tx.FileItem.Type.Eq(Directory)).
			FirstOrCreate()
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
	return &rpc.CloudFileItem{
		Id:          file.ID,
		FileName:    file.Name,
		FileType:    file.Type,
		NamespaceID: file.NamespaceID,
		UpdateTime:  file.UpdateTime.String(),
		CreateTime:  file.CreateTime.String(),
	}, err
}

// RemoveDirectory implements the CloudFileServiceImpl interface.
func (s *CloudFileServiceImpl) RemoveDirectory(ctx context.Context, id string) (err error) {

	err = dal.Q.Transaction(func(tx *dal.Query) error {
		first, err := tx.FileIndex.WithContext(ctx).Where(dal.FileIndex.ParentID.Eq(id)).First()
		if err != nil {
			return err
		}
		if first != nil {
			return errors.New("目录不为空，无法进行删除")
		}
		_, err = tx.FileIndex.WithContext(ctx).Where(dal.FileIndex.ChildID.Eq(id)).Delete()
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

func (s *CloudFileServiceImpl) CreateFileItem(ctx context.Context, name string, ty string, parentID string, namespaceID int64) (*rpc.CloudFileItem, error) {
	id := id.Base62()
	switch ty {
	case "namespace":
		id = "N-" + id
	case "directory":
		id = "D-" + id
	default:
		id = "F-" + id
	}
	var err error
	var fileItem *model.FileItem
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		fileItem, err = tx.FileItem.WithContext(ctx).Where(tx.FileItem.ID.Eq(id), tx.FileItem.Name.Eq(name), tx.FileItem.Type.Eq(ty), tx.FileItem.NamespaceID.Eq(namespaceID)).FirstOrCreate()
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
		return nil, err
	}
	return &rpc.CloudFileItem{
		Id:          fileItem.ID,
		FileName:    fileItem.Name,
		FileType:    fileItem.Type,
		NamespaceID: fileItem.NamespaceID,
		UpdateTime:  fileItem.UpdateTime.String(),
		CreateTime:  fileItem.CreateTime.String(),
	}, nil
}
func (s *CloudFileServiceImpl) CreateNamespace(ctx context.Context, name string) (r int64, err error) {
	namespace, err := dal.Namespace.WithContext(ctx).Where(dal.Namespace.Name.Eq(name)).FirstOrCreate()
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
	if err != nil {
		return nil, err
	}
	return &rpc.CloudFileItem{Id: res.ID, FileName: res.Name, FileType: res.Type, UpdateTime: res.UpdateTime.String(), CreateTime: res.CreateTime.String(), NamespaceID: res.NamespaceID}, err
}
func (s *CloudFileServiceImpl) GetFileURL(ctx context.Context, id string) (url string, err error) {
	item, err := s.QueryFileItemByID(ctx, id)
	if err != nil {
		return "", err
	}
	url, err = s.OssBucket.SignURL(fmt.Sprintf("cloud-file/namespace/%d/%s.%s", item.NamespaceID, id, item.FileType), oss.HTTPGet, 60*60*24)
	if err != nil {
		return "", err
	}
	return url, err
}
func (s *CloudFileServiceImpl) QueryUserNamespaces(ctx context.Context, uid int64) (r []*rpc.Namespace, err error) {
	var res []*rpc.Namespace
	err = dal.Namespace.WithContext(ctx).
		Select(dal.Namespace.ALL, dal.UserNamespace.Authority, dal.FileItem.ID.As("root_id")).
		LeftJoin(dal.UserNamespace, dal.Namespace.ID.EqCol(dal.UserNamespace.NamespaceID)).
		LeftJoin(dal.FileItem, dal.FileItem.NamespaceID.EqCol(dal.UserNamespace.NamespaceID), dal.FileItem.Type.Eq(Namespace)).
		Where(dal.UserNamespace.UserID.Eq(uid)).
		Debug().
		Scan(&res)
	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *CloudFileServiceImpl) GetUserIDByFileID(ctx context.Context, id string) (r int64, err error) {
	err = dal.UserNamespace.WithContext(ctx).
		Select(dal.UserNamespace.UserID).
		LeftJoin(dal.FileItem, dal.FileItem.NamespaceID.EqCol(dal.UserNamespace.NamespaceID), dal.UserNamespace.Authority.Eq(0)).
		Where(dal.FileItem.ID.Eq(id)).Debug().Scan(&r)
	//dal.Namespace.
	return r, err
}
func (s *CloudFileServiceImpl) LinkNamespace(ctx context.Context, namespaceId int64, uid int64, authority int32) error {
	err := dal.UserNamespace.WithContext(ctx).Create(&model.UserNamespace{
		UserID:      uid,
		NamespaceID: namespaceId,
		Authority:   authority,
	})
	return err
}
func (s *CloudFileServiceImpl) UploadFile(ctx context.Context, req *rpc.UploadFileRequest) (resp *rpc.CloudFileItem, err error) {
	id1 := "F-" + id.Base62()
	if err != nil {
		return nil, err
	}
	var f *model.FileItem
	err = dal.Q.Transaction(func(tx *dal.Query) error {
		fileType := ParseFileType(req.FileName)
		f, err = tx.FileItem.WithContext(ctx).Where(
			tx.FileItem.NamespaceID.Eq(req.NamespaceID),
			tx.FileItem.ID.Eq(id1),
			tx.FileItem.Name.Eq(req.FileName),
			tx.FileItem.Type.Eq(fileType)).FirstOrCreate()
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
		err = s.OssBucket.PutObject(fmt.Sprintf("cloud-file/namespace/%d/%s.%s", req.NamespaceID, id1, fileType), bytes.NewReader(req.FileData))
		if err != nil {
			fmt.Println("oss err:", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &rpc.CloudFileItem{
		Id:          f.ID,
		FileName:    f.Name,
		NamespaceID: f.NamespaceID,
		FileType:    f.Type,
		CreateTime:  f.CreateTime.String(),
		UpdateTime:  f.UpdateTime.String(),
	}, nil
}

func (s *CloudFileServiceImpl) FetchFileData(ctx context.Context, id string) ([]byte, error) {
	first, err := dal.FileItem.WithContext(ctx).Where(dal.FileItem.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	object, err := s.OssBucket.GetObject(fmt.Sprintf("cloud-file/namespace/%d/%s.%s", first.NamespaceID, first.ID, first.Type))
	if err != nil {
		return nil, err
	}
	all, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}
	return all, err
}
func (s *CloudFileServiceImpl) ModifyFileContent(ctx context.Context, id string, content string) error {
	first, err := dal.FileItem.WithContext(ctx).Where(dal.FileItem.ID.Eq(id)).First()
	if err != nil {
		return err
	}
	return s.OssBucket.PutObject(fmt.Sprintf("cloud-file/namespace/%d/%s.%s", first.NamespaceID, first.ID, first.Type), bytes.NewReader([]byte(content)))
}
func (s *CloudFileServiceImpl) CreateTextFile(ctx context.Context, name string, parentID string, content string, namespaceID int64) (*rpc.CloudFileItem, error) {
	t := ParseFileType(name)
	item, err := s.CreateFileItem(ctx, name, t, parentID, namespaceID)
	if err != nil {
		fmt.Println("create:", err)
		return nil, err
	}
	err = s.OssBucket.PutObject(fmt.Sprintf("cloud-file/namespace/%d/%s.%s", namespaceID, item.Id, t), bytes.NewReader([]byte(content)))
	if err != nil {
		fmt.Println("CreateTextFile:", err)
		err = s.Remove(ctx, item.Id)
		return nil, err
	}
	return item, nil
}
