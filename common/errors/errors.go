package errors

import (
    "fmt"
)

var (
    ErrUnknown = Error{-1, "Unknown error"}
    ErrSuccess = Error{0, "Success"}
)

type Error struct {
    Code int `json:"code"`
    Message string `json:"message"`
}

func (e *Error) Error() string {
    return fmt.Sprintf("<code:%d message:%s>", e.Code, e.Message)
}

var errorMessage = map[int]string{
    1001: "文章不存在",
    1002: "文章发表失败",
    2001: "评论不存在",
    2002: "评论发表失败",
    3001: "分类不存在",
    3002: "分类创建失败",
    9001: "文件无效",
}

func NewError(code int) Error {

    if v, ok := errorMessage[code]; ok == true {
        return Error{code, v}
    }
    return ErrUnknown
}
