namespace go rpc

struct LoginRequest{
    1:string username
    2:string password
}
struct RegisterRequest{
    1:string username
    2:string password
}

service UserService{
   i64 UserLogin(1:string username,2:string password)
   i64 UserRegister(1:string username,2:string password)
}
