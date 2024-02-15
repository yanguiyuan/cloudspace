namespace go rpc

struct LoginRequest{
    1:string username
    2:string password
}
struct RegisterRequest{
    1:string username
    2:string password
}
struct User{
    1:i64 id
    2:string username
    3:string password
    4:string email
    5:string gender
    6:string role
    7:string createTime
    8:string updateTime
}
service UserService{
   i64 UserLogin(1:string username,2:string password)
   i64 UserRegister(1:string username,2:string password)
   User GetUser(1:i64 id)
   void UpdateUser(1:User user)
   list<User> GetUsers(1:i32 offset,2:i32 limit)
}
