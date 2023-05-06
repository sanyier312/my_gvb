package v1

import (
	"github.com/gin-gonic/gin"
	"my_gvb/models"
	"my_gvb/utils/errmsg"
	"my_gvb/utils/validator"
	"net/http"
	"strconv"
)

// 添加用户
func AddUser(c *gin.Context) {
	var data models.User
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort()
		return
	}

	code := models.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		models.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := models.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := models.GetUsers(username, pageSize, pageNum)

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

// 编辑用户
func EditUser(c *gin.Context) {
	var data models.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := models.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCSE {
		models.EditUser(id, &data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	},
	)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := models.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改密码
func ChangePassword(c *gin.Context) {
	var data models.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)

	code := models.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		models.ChangePassword(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
