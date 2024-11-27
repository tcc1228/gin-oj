package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string        `gorm:"coulumn:identity;type:varchar(36);" json:"identity"`                 //唯一标识
	ProblemIdentity string        `gorm:"coulumn:problem_identity;type:varchar(36);" json:"problem_identity"` //问题表的唯一标识
	ProblemBasic    *ProblemBasic `gorm:"foreignKey:identity;references:problem_identity"`                    //关联问题基础表
	UserIdentity    string        `gorm:"coulumn:user_identity;type:varchar(36);" json:"user_identity"`       //用户的唯一标识
	UserBasic       *UserBasic    `gorm:"foreignKey:identity;references:user_identity"`                       //关联用户基础表
	Path            string        `gorm:"coulumn:path;type:varchar(255);" json:"path"`                        //代码存放路径
	Status          int           `gorm:"coulumn:status;type:tinyint;" json:"status"`                         //[0-待判断，1-答案正确，2-答案错误，3-运行超时，4-运行超内存，5-编译错误]
}

func (SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(problemIdentity, userIdentity string, status int) *gorm.DB {
	tx := DB.Model(new(SubmitBasic)).Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
		return db.Omit("content")
	}).Preload("UserBasic")
	if problemIdentity != "" {
		tx.Where("problem_identity = ?", problemIdentity)
	}
	if userIdentity != "" {
		tx.Where("user_identity = ?", userIdentity)
	}
	if status != 0 {
		tx.Where("status = ?", status)
	}
	return tx
}
