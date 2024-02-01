namespace go rpc

struct AddRequest{
    1:string fileName
    2:string parentID
    3:binary fileData
    4:i64 uid
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
struct UpdateRequest{
    1:string id
    2:string content
    3:i64 uid
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
    2:i64 uid
    3:string filename
}
service CloudFileService{
    void add(1:AddRequest req)
    void createDirectory(1:CreateDirectoryRequest req)
    void remove(1:RemoveRequest req)
    void removeDirectory(1:string id)
    QueryResponse query(1:string pid,2:i64 uid )
    CloudFileItem queryFileItemByID(1:string id)
    void update(1:UpdateRequest req)
    void rename(1:string id,2:string newName)
    string queryUserFileRoot(1:i64 userID)
    string createFileItem(1:string name,2:string ty,3:string parentID)
    i64 createNamespace(1:string name,2:string rootID)
    void createUserNamespace(1:i64 userID,2:i64 namespaceID,3:i32 authority)
}
