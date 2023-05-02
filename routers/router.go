package routers

import (
	"github.com/gin-gonic/gin"
	v1 "my_gvb/api/v1"
	"my_gvb/middleware"
	"my_gvb/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode) //设置运行模式
	r := gin.New()             //创建路由
	r.Use(middleware.Logger()) //日志中间件
	r.Use(gin.Recovery())      //恢复中间件
	r.Use(middleware.Cors())   //跨域中间件

	//静态资源
	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "./web/front/dist/favicon.ico")
	//前台页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	//后台管理的路由
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken()) //使用中间件
	{
		//用户模块的路由接口
		auth.GET("users", v1.GetUsers)         //查询用户列表
		auth.PUT("user/:id", v1.EditUser)      //编辑用户
		auth.DELETE("user/:id", v1.DeleteUser) //删除用户
		//作者信息的路由接口
		auth.GET("admin/profile/:id", v1.GetProfile) //获取作者信息
		auth.PUT("profile/:id", v1.UpdateProfile)    //更新作者信息
		//修改密码
		auth.PUT("admin/changePassword/:id", v1.ChangePassword)
		//分类模块的路由接口
		auth.GET("admin/category", v1.GetCategory)     //查询分类列表
		auth.POST("category/add", v1.AddCategory)      //添加分类
		auth.PUT("category/:id", v1.EditCategory)      //编辑分类
		auth.DELETE("category/:id", v1.DeleteCategory) //删除分类
		//文章模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArtInfo) //查询指定文章
		auth.GET("admin/article", v1.GetArticleList)      //查询文章列表
		auth.POST("article/add", v1.AddArticle)           //添加文章
		auth.PUT("article/:id", v1.EditArticle)           //编辑文章
		auth.DELETE("article/:id", v1.DeleteArticle)      //删除文章
		//评论模块的路由接口
		auth.GET("comment/list", v1.GetCommentList)       //获取评论列表
		auth.PUT("checkcomment/:id", v1.CheckComment)     //通过评论
		auth.PUT("uncheckcomment/:id", v1.UncheckComment) //撤销评论
		auth.DELETE("delcomment/:id", v1.DeleteComment)   //删除评论
		//上传图片
		auth.POST("upload", v1.UpLoad)
	}

	//前台的路由
	router := r.Group("api/v1") //创建路由组
	{
		//用户模块的路由接口
		router.POST("user/add", v1.AddUser) //添加用户
		//作者信息的路由接口
		router.GET("profile/:id", v1.GetProfile) //获取作者信息

		//文章模块的路由接口
		router.GET("article", v1.GetArticleList)      //查询文章列表
		router.GET("article/info/:id", v1.GetArtInfo) //查询指定文章
		router.GET("article/list/:id", v1.GetCateArt) //查询指定分类的所有文章

		//分类模块的路由接口
		router.GET("category", v1.GetCategory) //查询分类

		//登录模块的路由接口
		router.POST("login", v1.Login)           //后台登录
		router.POST("loginfront", v1.LoginFront) //前台登录

		//评论模块的路由接口
		router.POST("addcomment", v1.AddComment)               //添加评论
		router.GET("comment/info/:id", v1.GetComment)          //获取单个评论
		router.GET("commentfront/:id", v1.GetCommentListFront) //获取评论列表
		router.GET("commentcount/:id", v1.GetCommentCount)     //获取评论数

	}
	r.Run(utils.HttpPort) //运行
}
