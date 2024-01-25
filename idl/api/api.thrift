
namespace go api
struct AddRequest{
    1:string fileName
    2:string parentID
    3:binary fileData
    4:string username
    5:string path
}
struct CloudFileItem{
    1:string id
    2:string fileName
    3:string fileType
    4:string updateTime
    5:string createTime
}
struct QueryResponse{
    1:list<CloudFileItem> items
    2:list<UrlTable> urlmaps
}
struct QueryRequest{
    1:string pid
    2:string username
}
struct UpdateRequest{
    1:string id
    2:string content
    3:string username
}
struct UrlTable{
    1:string id
    2:string url
}
struct CreateDirectoryRequest{
    1:string directoryName
    2:string parentID
}
struct RemoveRequest{
    1:string id
    2:string username
    3:string filename
}
service CloudFileService{
    void add(1:AddRequest req)
    void createDirectory(1:CreateDirectoryRequest req)
    void remove(1:RemoveRequest req)
    void removeDirectory(1:string id)
    QueryResponse query(1:QueryRequest req)
    void update(1:UpdateRequest req)
    void rename(1:string id,2:string newName)
}
