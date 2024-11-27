package router

import (
	"gin-gorm-oj/middleware"
	"gin-gorm-oj/service"

	_ "gin-gorm-oj/docs"

	"github.com/gin-gonic/gin"
	swaggerfies "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// 初始化gin
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfies.Handler))

	//路由规则

	//公有方法
	//问题
	r.GET("/problem", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	//用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)
	// 排行榜
	r.GET("/rank-list", service.GetRankList)

	//提交记录
	r.GET("/submit-list", service.GetSubmitList)

	// 管理员私有方法
	authAdmin := r.Group("/admin", middleware.AuthAdminCheck())
	// 问题创建
	authAdmin.POST("/problem-create", service.ProblemCreate)
	// 问题修改
	authAdmin.PUT("/problem-modify", service.ProblemModify)
	// 分类列表
	authAdmin.GET("/category-list", service.GetCategoryList)
	// 分类创建
	authAdmin.POST("/category-create", service.CategoryCreate)
	// 分类删除
	authAdmin.DELETE("/category-delete", service.CategoryDelete)
	// 分类修改
	authAdmin.PUT("/category-modify", service.CategoryModify)

	// 用户私有方法
	authUser := r.Group("/user", middleware.AuthUserCheck())
	// 代码提交
	authUser.POST("/submit", service.Submit)

	return r
}
