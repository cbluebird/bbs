include "base.thrift"
namespace go user

struct User {
    1: required string ID,
    2: required string UserName,
    3: required string Password,
    4: required string StudentID,
}

struct CreateUserResponse {
    1:base.BaseResp base_resp
}

struct CreateUserRequest {
    1:string StudentID
    2:string Password
    3:string IID
    4:string UserName
}

struct LoginRequest {
    1:string StudentID
    2:string Password
}

struct LoginResponse {
    1:base.BaseResp base_resp
    255: User user
}

service UserService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    LoginResponse Login(1:LoginRequest req)
}
