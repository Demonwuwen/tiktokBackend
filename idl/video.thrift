namespace go tiktokBackendBase

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


service VideoService{
    PublishActionResponse PublishVideo(1: PublishActionRequest req)
    PublishListResponse PublishList(1: PublishListRequest req)
    FeedResponse Feed(1: FeedRequest req)
}