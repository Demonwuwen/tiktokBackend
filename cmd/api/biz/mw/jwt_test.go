package mw

import (
	"fmt"
	"testing"
)

type Int64 = int64
type Int32 = int32
type UserRegisterResponse struct {
	StatusCode Int32  `thrift:"status_code,1,required" frugal:"1,required,i32" json:"status_code"`
	StatusMsg  string `thrift:"status_msg,2" frugal:"2,default,string" json:"status_msg"`
	UserId     Int64  `thrift:"user_id,3,required" frugal:"3,required,i64" json:"user_id"`
	Token      string `thrift:"token,4,required" frugal:"4,required,string" json:"token"`
}

func TestGenToken(t *testing.T) {
	InitJWT()
	id := int64(23)

	resp := &UserRegisterResponse{}
	resp.UserId = id
	resp.StatusMsg = "asdasd"
	token, expire, err := JwtMiddleware.TokenGenerator(resp.UserId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(expire)
	fmt.Println(token)

}
