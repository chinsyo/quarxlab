package errors

import (
	"fmt"
)

var (
	ErrUnknown = Error{-1, "Unknown error"}
	ErrSuccess = Error{0, "Success"}

	ErrTokenRequired = Error{101, "Token required"}
	ErrTokenFailure  = Error{102, "Token failure"}
	ErrTokenExpired  = Error{103, "Token expired"}

	ErrArticleNotExist  = Error{1001, "Article not exist"}
	ErrArticlePubFailed = Error{1002, "Article pub failed"}

	ErrCommentNotExist  = Error{2001, "Comment not exist"}
	ErrCommentPubFailed = Error{2002, "Comment pub failed"}

	ErrCategoryNotExist  = Error{3001, "Category not exist"}
	ErrCategoryPubFailed = Error{3002, "Category pub failed"}

	ErrUserNotExist     = Error{4001, "User not exist"}
	ErrUserCreateFailed = Error{4002, "User create failed"}
	ErrUserAlreadyExist = Error{4003, "User already exist"}
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("<code:%d message:%s>", e.Code, e.Message)
}

var errorMessage = map[int]string{
	// global
	0: "Success",
	1: "Unkown error",

	// token
	101: "Token required",
	102: "Token failure",
	103: "Token expired",

	// articles
	1001: "文章不存在",
	1002: "文章发表失败",
	// comments
	2001: "评论不存在",
	2002: "评论发表失败",
	// categories
	3001: "分类不存在",
	3002: "分类创建失败",
	// users: register
	4001: "用户不存在",
	4002: "用户创建失败",
	4003: "用户名已占用",
	// users: login
	4101: "用户名或密码不能为空",
	4102: "密码错误",
	4103: "用户名格式错误",
	4104: "密码格式错误",
	4105: "没有验证码需要校验",
	// files
	9001: "文件无效",
	9900: "Redis连接失败",
}

func NewError(code int) Error {

	if v, ok := errorMessage[code]; ok == true {
		return Error{code, v}
	}
	return ErrUnknown
}
