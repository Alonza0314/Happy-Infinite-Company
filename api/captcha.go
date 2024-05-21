package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// 返回驗證碼圖片
func GetCaptcha(c *gin.Context) {
	id, imgData, code, err := MakeCaptcha()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"code":  code,
		"image": imgData,
	})
}

// 生成驗證碼
func MakeCaptcha() (string, string, string, error) {
	// 驗證碼配置
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	/*
		width：寬度。
		height：高度。
		noiseCount：干擾線數量。
		showLineOptions：干擾線的可見度，0是不顯示，1是顯示。
		length：驗證碼長度。
	*/

	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	// 生成驗證碼
	id, b64s, code, err := cp.Generate()
	if err != nil {
		return "", "", "", errors.New("服務器錯誤，驗證碼生成失敗")
	}

	return id, b64s, code, nil
}
