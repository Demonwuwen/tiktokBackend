namespace go user

typedef i32 int32
typedef i64 int64

struct User{
  1: required int64 id //用户id
  2: required string name // 用户名
  3: optional int64 follow_count//总关注数
  4: optional int64 follower_count //粉丝总数
  5: required bool is_follow//true-已关注，false-未关注
}


struct UserRegisterRequest {
  1: required string username //注册用户名，最长32个字符
  2: required string password //密码，最长32个字符
}

struct UserRegisterResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: required int64 user_id // 用户id
  4: required string token //用户鉴权token
}

struct UserLoginRequest {
  1: required string username //注册用户名，最长32个字符
  2: required string password //密码，最长32个字符
}

struct UserLoginResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: required int64 user_id // 用户id
  4: required string token //用户鉴权token
}


struct UserRequest {
  1: required int64 user_id //注册用户名，最长32个字符
  2: required string token //用户鉴权token
}

struct UserResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: required User user // 用户信息

}

service UserService{
    UserRegisterResponse UserRegister(1: UserRegisterRequest req)
    UserLoginResponse UserLogin(1: UserLoginRequest req)
    UserResponse UserGet(1: UserRequest req)

}