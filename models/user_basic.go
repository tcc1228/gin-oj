package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity         string `gorm:"coulumn:identity;type:varchar(36);" json:"identity"`            //用户的唯一标识
	Name             string `gorm:"column:name;type:varchar(100);" json:"name"`                    //用户名
	Password         string `gorm:"column:password;type:varchar(32);" json:"password"`             //密码
	Phone            string `gorm:"column:phone;type:varchar(11);" json:"phone"`                   //手机号
	Mail             string `gorm:"column:mail;type:varchar(100);" json:"mail"`                    //邮箱
	FinishProblemNum int64  `gorm:"column:finish_problem_num;type:int;" json:"finish_problem_num"` //完成题目数量
	SubmitNum        int64  `gorm:"column:submit_num;type:int;" json:"submit_num"`                 //提交次数
	IsAdmin          int    `gorm:"column:is_admin;type:tinyint;" json:"is_admin"`                 //是否是管理员【0否，1-是】
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
