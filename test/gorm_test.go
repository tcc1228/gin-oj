package test

import (
	"fmt"
	"gin-gorm-oj/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGormTest(t *testing.T) {
	dsn := "root:tcc200128@tcp(121.4.174.222:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 连接数据库
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.ProblemBasic, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("Problem: %v \n", v)
	}
}
