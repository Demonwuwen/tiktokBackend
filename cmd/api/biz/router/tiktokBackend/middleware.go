// Code generated by hertz generator.

package TiktokBackend

import (
	"Demonwuwen/tiktokBackend/cmd/api/biz/mw"
	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _usergetMw() []app.HandlerFunc {
	// your code...

	return []app.HandlerFunc{
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _comment_ctMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoriteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favorite_ctMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _favoritelistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishvideoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relationMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _relation_ctMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userloginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userregisterMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _friendlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followlistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _followerlistMw() []app.HandlerFunc {
	// your code...
	return nil
}
