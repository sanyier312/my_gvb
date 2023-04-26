package v1

import (
	"github.com/gin-gonic/gin"
	"my_gvb/service"
	"my_gvb/utils/errmsg"
	"net/http"
)

// UpLoad 上传图片接口
func UpLoad(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR_UPLOAD_FAIL),
			"url":     "",
		})
		return
	}
	fileSize := fileHeader.Size

	url, code := service.UpLoadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
