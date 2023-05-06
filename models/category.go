package models

import (
	"gorm.io/gorm"
	"my_gvb/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 查询分类
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

// CreateCategory 新增分类Category
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// 查询分类列表
func GetCategory(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err = db.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total

}

// GetCateInfo 查询单个分类信息
func GetCateInfo(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCSE
}

// 编辑分类
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCategory(id int) int {
	var category Category
	err = db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
