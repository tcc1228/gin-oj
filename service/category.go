package service

import (
	"gin-gorm-oj/define"
	"gin-gorm-oj/helper"
	"gin-gorm-oj/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCategoryList
// @Tags 管理员私有方法
// @Summary 分类列表
// @Param authorization header string true "authorization"
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200",data:""}"
// @Router /admin/category-list [get]
func GetCategoryList(c *gin.Context) {
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetCategoryList Page Strconv Error:", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")

	categoryList := make([]*models.CategoryBasic, 0)
	err = models.DB.Model(new(models.CategoryBasic)).Where("name like ?", "%"+keyword+"%").Count(&count).
		Offset(page).Limit(size).Find(&categoryList).Error
	if err != nil {
		log.Println("GetCategoryList Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "GetCategoryList Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  categoryList,
			"count": count,
		},
	})
}

// CategoryCreate
// @Tags 管理员私有方法
// @Summary 分类创建
// @Param authorization header string true "authorization"
// @Param name formData string true "name"
// @Param parentId formData string false "parentId"
// @Success 200 {string} json "{"code":"200",data:""}"
// @Router /admin/category-create [post]
func CategoryCreate(c *gin.Context) {
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))
	category := &models.CategoryBasic{
		Identity: helper.GetUUID(),
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Create(category).Error
	if err != nil {
		log.Println("CategoryCreate Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "CategoryCreate Error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "创建成功",
	})
}

// CategoryModify
// @Tags 管理员私有方法
// @Summary 分类修改
// @Param authorization header string true "authorization"
// @Param identity formData string true "identity"
// @Param name formData string true "name"
// @Param parentId formData string false "parentId"
// @Success 200 {string} json "{"code":"200",data:""}"
// @Router /admin/category-modify [put]
func CategoryModify(c *gin.Context) {
	identity := c.PostForm("identity")
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parentId"))
	if name == "" || identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	category := &models.CategoryBasic{
		Identity: identity,
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Model(new(models.CategoryBasic)).Where("identity = ?", identity).Updates(category).Error
	if err != nil {
		log.Println("CategoryModify Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "CategoryModify Error",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

// CategoryDelete
// @Tags 管理员私有方法
// @Summary 分类删除
// @Param authorization header string true "authorization"
// @Param identity query string true "identity"
// @Success 200 {string} json "{"code":"200",data:""}"
// @Router /admin/category-delete [delete]
func CategoryDelete(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不正确",
		})
		return
	}
	var cnt int64
	err := models.DB.Model(new(models.ProblemCategory)).Where("category_id = (SELECT id FROM category_basic WHERE identity = ? LIMIT 1)", identity).Count(&cnt).Error
	if err != nil {
		log.Println("Get ProblemCategory Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取分类关联的问题失败",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "该分类下已存在问题，不可删除",
		})
		return
	}
	err = models.DB.Where("identity = ?", identity).Delete(&models.CategoryBasic{}).Error
	// err = models.DB.Where("identity = ?", identity).Delete(new(models.CategoryBasic)).Error
	if err != nil {
		log.Println("CategoryDelete Error:", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
