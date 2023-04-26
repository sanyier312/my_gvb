package models

import "my_gvb/utils/errmsg"

type Profile struct {
	ID        int    `gorm:"primaryKey" json:"id"`                // ID
	Name      string `gorm:"type:varchar(20)" json:"name"`        // 名称
	Desc      string `gorm:"type:varchar(200)" json:"desc"`       // 个人描述
	Qq        string `gorm:"type:varchar(200)" json:"qq"`         // QQ
	Wechat    string `gorm:"type:varchar(100)" json:"wechat"`     // 微信
	Github    string `gorm:"type:varchar(200)" json:"github"`     // Github
	Email     string `gorm:"type:varchar(200)" json:"email"`      // 邮箱
	Img       string `gorm:"type:varchar(200)" json:"img"`        // 个人图片
	Avatar    string `gorm:"type:varchar(200)" json:"avatar"`     // 头像
	IcpRecord string `gorm:"type:varchar(200)" json:"icp_record"` // ICP备案号
}

// GetProfile 获取个人信息
func GetProfile(id int) (Profile, int) {
	var profile Profile
	err = db.Where(("ID = ?"), id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR

	}
	return profile, errmsg.SUCCSE
}

// UpdateProfile 更新个人信息
func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
