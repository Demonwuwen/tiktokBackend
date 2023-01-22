// Code generated by hertz generator. DO NOT EDIT.

package TiktokBackendBase

import (
	tiktokBackendBase "Demonwuwen/tiktokBackend/cmd/api/biz/handler/tiktokBackendBase"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		_douyin.GET("/feed", append(_feedMw(), tiktokBackendBase.Feed)...)
		_douyin.GET("/user", append(_usergetMw(), tiktokBackendBase.UserGet)...)
		{
			_comment := _douyin.Group("/comment", _commentMw()...)
			_comment.POST("/action", append(_comment_ctMw(), tiktokBackendBase.CommentAct)...)
			_comment.GET("/list", append(_commentlistMw(), tiktokBackendBase.CommentList)...)
		}
		{
			_favorite := _douyin.Group("/favorite", _favoriteMw()...)
			_favorite.POST("/action", append(_favorite_ctMw(), tiktokBackendBase.FavoriteAct)...)
			_favorite.GET("/list", append(_favoritelistMw(), tiktokBackendBase.FavoriteList)...)
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			_publish.POST("/action", append(_publishvideoMw(), tiktokBackendBase.PublishVideo)...)
			_publish.GET("/list", append(_publishlistMw(), tiktokBackendBase.PublishList)...)
		}
		{
			_relation := _douyin.Group("/relation", _relationMw()...)
			_relation.POST("/action", append(_relation_ctMw(), tiktokBackendBase.RelationAct)...)
			{
				_follower := _relation.Group("/follower", _followerMw()...)
				_follower.GET("/list", append(_fowllowlistMw(), tiktokBackendBase.FowllowList)...)
				_follower.GET("/list", append(_fowllowerlistMw(), tiktokBackendBase.FowllowerList)...)
				_follower.GET("/list", append(_frientlistMw(), tiktokBackendBase.FrientList)...)
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.POST("/login", append(_userloginMw(), tiktokBackendBase.UserLogin)...)
			_user.POST("/register", append(_userregisterMw(), tiktokBackendBase.UserRegister)...)
		}
	}
}
