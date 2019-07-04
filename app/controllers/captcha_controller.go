package controllers

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	xctx "quarxlab/lib/context"
	xerrors "quarxlab/lib/errors"
	"time"
)

type CaptchaResponse struct {
	CaptchaID string `json:"captcha_id"` //验证码Id
	ImageURL  string `json:"image_url"`  //验证码图片url
}

type captchaController int

const CaptchaController = captchaController(0)

func (this captchaController) Refresh(c *gin.Context) {
	captchaLen := captcha.DefaultLen
	captchaID := captcha.NewLen(captchaLen)
	log.Println("captchaID", captchaID)

	c.Set(xctx.CID, captchaID)

	ctxCID, ok := c.Get(xctx.CID)
	if !ok {
		err := xerrors.NewError(4105)
		panic(err)
	}
	captchaID1 := ctxCID.(string)
	log.Println("captchaID1", captchaID1)

	err := this.serve(c.Writer, c.Request, captchaID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "无法获取验证码", "data": gin.H{}})
		return
	}
}

func (this captchaController) Verify(c *gin.Context) {

	ctxCID, ok := c.Get(xctx.CID)
	if !ok {
		err := xerrors.NewError(4105)
		panic(err)
	}
	captchaID := ctxCID.(string)
	log.Println("captchaID", captchaID)

	value := c.PostForm("value")
	if captchaID == "" || value == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "参数错误", "data": gin.H{}})
		return
	}

	if captcha.VerifyString(captchaID, value) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "验证成功", "data": gin.H{}})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "验证失败", "data": gin.H{}})
	}
}

func (this captchaController) serve(w http.ResponseWriter, r *http.Request, captchaID string) error {

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")

	var content bytes.Buffer
	err := captcha.WriteImage(&content, captchaID, captcha.StdWidth, captcha.StdHeight)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	name := captchaID + ".png"
	http.ServeContent(w, r, name, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
