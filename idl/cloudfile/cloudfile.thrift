namespace go rpc

struct AddRequest{
    1:string fileName
    2:string parentID
    3:binary fileData
    4:i64 uid
}
struct UploadFileRequest{
    1:string fileName
    2:string parentID
    3:binary fileData
    4:i64 namespaceID
}
struct CloudFileItem{
    1:string id
    2:string fileName
    3:string fileType
    4:string updateTime
    5:string createTime
    6:i64 namespaceID
}
struct Namespace{
    1:i64 id
    2:string name
    3:string rootID
    4:i32 authority
    5:string createTime
    6:string updateTime
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
    CloudFileItem uploadFile(1:UploadFileRequest req)
    CloudFileItem add(1:AddRequest req)
    CloudFileItem createDirectory(1:CreateDirectoryRequest req)
    void remove(1:string id)
    void removeDirectory(1:string id)
    QueryResponse query(1:string pid,2:i64 uid )
    CloudFileItem queryFileItemByID(1:string id)
    void update(1:UpdateRequest req)
    void rename(1:string id,2:string newName)
    CloudFileItem createFileItem(1:string name,2:string ty,3:string parentID,4:i64 namespaceID)
    i64 createNamespace(1:string name)
    void createUserNamespace(1:i64 userID,2:i64 namespaceID,3:i32 authority)
    string getFileURL(1:string id)
    list<Namespace> queryUserNamespaces(1:i64 userID)
    void LinkNamespace(1:i64 userID,2:i64 namespaceID,3:i32 authority)
    i64 getUserIDByFileID(1:string id)
    binary fetchFileData(1:string id)
    void modifyFileContent(1:string id,2:string content)
    CloudFileItem createTextFile(1:string name,2:string parentID,3:string content,4:i64 namespaceID)
}
