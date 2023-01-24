namespace go tiktokBackend

typedef i32 int32
typedef i64 int64



struct FeedRequest {
  optional int64 latest_time//可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
  optional string token // 可选参数，登录用户设置
}

struct FeedResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg //返回状态描述
  3: list<Video>  video_list //视频列表
  4: optional int64 next_time  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct Video{
  1: required int64 id //视频唯一标识
  2: required User author //视频作者信息
  3: required string play_url // 视频播放地址
  4: required string cover_url // 视频封面地址
  5: required int64 favorite_count //视频的点赞数
  6: required int64 comment_count //视频的总评论数
  7: required bool is_favorite  //true-已点赞，false-未点赞
  8: required string title //视频标题
}

struct User{
  1: required int64 id //用户id
  2: required string name // 用户名
  3: optional int64 follow_count//总关注数
  4: optional int64 follower_count //粉丝总数
  5: required bool is_follow//true-已关注，false-未关注
}

struct Comment{
    1: required int64 id //用户id
    2: required User user // 评论用户信息
    3: required string content//评论内容
    4: required string create_date //评论发布日期，格式mm-dd
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

struct PublishActionRequest {
  1: required string token //注册用户名，最长32个字符
  2: required binary data //视频数据
  3: required string title //视频标题
}

struct PublishActionResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述

}

struct PublishListRequest {
  1: required int64 user_id //注册用户名，最长32个字符
  2: required string token //用户鉴权token
}

struct PublishListResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: list<Video> video_list //用户发布的视频列表
}


struct FavoriteActionRequest {
  1: required string token //用户鉴权token
  2: required int64 video_id //注册用户名，最长32个字符
  3: required int32 action_type //用户鉴权token
}

struct FavoriteActionResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
}

struct FavoriteListRequest {
  1: required int64 user_id //用户id
  2: required string token //用户鉴权token
}

struct FavoriteListResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: list<Video>  video_list //点赞视频列表
}

struct CommentActionRequest {
  1: required string token //用户鉴权token
  2: required int64 video_id //注册用户名，最长32个字符
  3: required int32 action_type //1-发布评论，2-删除评论
  4: required string comment_text //用户填写评论内容，在action_type=1的时候使用
  5: required int64 comment_id //要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: optional Comment comment //评论成功返回内容，不需要重新拉去整个列表
}

struct CommentListRequest {
  1: required string token //用户鉴权token
  2: required int64 video_id //用户id
}

struct CommentListResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: list<Video>  video_list //点赞视频列表
}


struct RelationActionRequest {
  1: required string token //用户鉴权token
  2: required int64 to_user_id //注册用户名，最长32个字符
  3: required int32 action_type //1-关注，2-取消关注
}

struct RelationActionResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
}


struct FollowListRequest {
  1: required int64 user_id //用户id
  2: required string token //用户鉴权token
}

struct FollowListResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: list<User>  user_list //用户信息列表
}

struct FollowerListRequest {
  1: required int64 user_id //用户id
  2: required string token //用户鉴权token
}

struct FollowerListResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: list<User>  user_list //用户信息列表
}

struct FriendListRequest {
  1: required int64 user_id //用户id
  2: required string token //用户鉴权token
}

struct FriendListResponse {
  1: required int32  status_code //状态吗，0-成功，其他失败
  2: optional  string status_msg//返回状态描述
  3: list<User>  user_list //用户信息列表
}

service ApiService{
    UserRegisterResponse UserRegister(1: UserRegisterRequest req) (api.post="/douyin/user/register")
    UserLoginResponse UserLogin(1: UserLoginRequest req) (api.post="/douyin/user/login")
    UserResponse UserGet(1: UserRequest req) (api.get="/douyin/user")
    PublishActionResponse PublishVideo(1: PublishActionRequest req) (api.post="/douyin/publish/action")
    PublishListResponse PublishList(1: PublishListRequest req) (api.get="/douyin/publish/list")
    FeedResponse Feed(1: FeedRequest req) (api.get="/douyin/feed")

    //interact
    FavoriteActionResponse FavoriteAct(1:FavoriteActionRequest req) (api.post="/douyin/favorite/action")
    FavoriteListResponse  FavoriteList(1:FavoriteListRequest req) (api.get="/douyin/favorite/list")
    CommentActionResponse CommentAct(1:CommentActionRequest req) (api.post="/douyin/comment/action")
    CommentListResponse CommentList(1:CommentListRequest req) (api.get="/douyin/comment/list")

    //relation
     RelationActionResponse RelationAct(1:RelationActionRequest req) (api.post="/douyin/relation/action")
     FollowListResponse FowllowList(1:FollowListRequest req) (api.get="/douyin/relation/follow/list")
     FollowerListResponse FowllowerList(1:FollowerListRequest req) (api.get="/douyin/relation/follower/list")
     FriendListResponse FrientList(1:FriendListRequest req) (api.get="/douyin/relation/friend/list")
}