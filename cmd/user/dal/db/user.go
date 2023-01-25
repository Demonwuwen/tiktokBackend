package db

import (
	"Demonwuwen/tiktokBackend/kitex_gen/user"
	"Demonwuwen/tiktokBackend/pkg/consts"
	"Demonwuwen/tiktokBackend/pkg/util"
	"context"
	"errors"
	"fmt"
	"time"
)

type UserLogin struct {
	Id       int64 `gorm:"primary_key"`
	UserId   int64
	Username string `gorm:"primary_key"`
	Password string `gorm:"size:200;notnull"`
}

func (l *UserLogin) TableName() string {
	return consts.UserLoginTableName
}

var (
	ErrUsrExist = errors.New("Username already exits ")
)

// CreateUserLogin create userLogin info
func CreateUserLogin(ctx context.Context, userLogin []*UserLogin) error {
	return DB.WithContext(ctx).Create(userLogin).Error
}

func QueryUserLogin(cxt context.Context, username string) ([]*UserLogin, error) {
	res := make([]*UserLogin, 0)
	err := DB.WithContext(cxt).Where("username = ?", username).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

type User struct {
	Id            int64      `json:"id" gorm:"id,omitempty"`                         //用户id
	Name          string     `json:"name" gorm:"name,omitempty"`                     // 用户名
	FollowCount   int64      `json:"follow_count" gorm:"follow_count,omitempty"`     //总关注数
	FollowerCount int64      `json:"follower_count" gorm:"follower_count,omitempty"` //粉丝总数
	IsFollow      bool       `json:"is_follow" gorm:"is_follow,omitempty"`           //true-已关注，false-未关注
	User          *UserLogin `json:"-"`                                              //用户与账号密码之间的一对一
	Videos        []*Video   `json:"-"`                                              //用户与投稿视频的一对多
	Follows       []*User    `json:"-" gorm:"many2many:user_relations;"`             //用户之间的多对多
	FavorVideos   []*Video   `json:"-" gorm:"many2many:user_favor_videos;"`          //用户与点赞视频之间的多对多
	Comments      []*Comment `json:"-"`                                              //用户与评论的一对多
}

func CreateUser(ctx context.Context, user *User) (int64, error) {
	fmt.Println("createUser")
	if user.Name == "" {
		user.Name = util.RandNameString(consts.RandNameLen)
	}
	err := DB.WithContext(ctx).Create(user).Error
	return user.Id, err
}

func QueryUserById(cxt context.Context, userId int64) (*user.User, error) {
	res := &user.User{}
	err := DB.WithContext(cxt).Where("id = ?", userId).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

type Video struct {
	Id            int64      `json:"id,omitempty"`              //视频唯一标识
	UserId        int64      `json:"-"`                         //
	Author        User       `json:"author,omitempty" gorm:"-"` //视频作者信息
	PlayUrl       string     `json:"play_url,omitempty"`        // 视频播放地址
	CoverUrl      string     `json:"cover_url,omitempty"`       // 视频封面地址
	FavoriteCount int64      `json:"favorite_count,omitempty"`  //视频的点赞数
	CommentCount  int64      `json:"comment_count,omitempty"`   //视频的总评论数
	IsFavorite    bool       `json:"is_favorite,omitempty"`     //true-已点赞，false-未点赞
	Title         string     `json:"title,omitempty"`           //视频标题
	Users         []*User    `json:"-" gorm:"many2many:user_favor_videos;"`
	Comments      []*Comment `json:"-"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     time.Time  `json:"-"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`            //视频评论id
	UserId     int64  `json:"-"`                       //用于一对多关系的id
	User       User   `json:"user,omitempty" gorm:"-"` //视频作者信息
	Content    string `json:"content,omitempty"`       // 评论内容
	VideoId    int64  `json:"-"`                       //一对多，视频对评论
	CreateData string `json:"create_data,omitempty"`   // 评论发布日期，格式mm-dd
}
