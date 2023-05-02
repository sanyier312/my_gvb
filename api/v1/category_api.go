package v1

import (
	"github.com/gin-gonic/gin"
	"my_gvb/models"
	"my_gvb/utils/errmsg"
	"net/http"
	"strconv"
)

// 添加种类
func AddCategory(c *gin.Context) {
	var data models.Category
	_ = c.ShouldBindJSON(&data)
	code = models.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		models.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		code = errmsg.ERROR_CATENAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 查询种类
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := models.GetCategory(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 编辑种类
func EditCategory(c *gin.Context) {
	var data models.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	code = models.CheckUser(data.Name)
	if code == errmsg.SUCCSE {
		models.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}

// 删除种类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = models.DeleteCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
