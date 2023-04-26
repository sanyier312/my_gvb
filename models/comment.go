package models

import (
	"gorm.io/gorm"
	"my_gvb/utils/errmsg"
)

type Comment struct {
	gorm.Model
	SubComments     []*Comment `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`   //子评论列表
	ParentComment   *Comment   `gorm:"foreignKey:ParentCommentID" json:"parent_comment"` //父级评论列表
	ParentCommentID *uint      `json:"parent_comment_id"`                                //父评论id
	Content         string     `gorm:"size:256" json:"content"`
	ArticleID       uint       `json:"article_id"`                           //文章id
	UserID          uint       `json:"user_id"`                              //评论的用户
	Status          int        `gorm:"type:tinyint;default:1" json:"status"` //评论状态
}

// 增加评论
func AddComment(comment *Comment) int {
	err = db.Create(comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询单个评论
func GetComment(id int) (Comment, int) {
	var comment Comment
	err = db.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return comment, errmsg.ERROR
	}
	return comment, errmsg.SUCCSE
}

// 后台查询评论列表
func GetCommentList(pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Find(&commentList).Count(&total)
	err = db.Model(&commentList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCSE
}

// 查询评论数量
func GetCommentCount(id int) int64 {
	var comment Comment
	var total int64
	db.Find(&comment).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	return total
}

// 前台查询评论列表
func GetCommentListFront(id int, pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Find(&Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	err = db.Model(&Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?",
		id).Where("status = ?", 1).Scan(&commentList).Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCSE
}

// 删除评论
func DeleteComment(id uint) int {
	var comment Comment
	err = db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 评论通过审核 评论数量加一
func CheckComment(id int, data *Comment) int {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id = ?", res.ArticleID).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// UncheckComment 撤下评论  评论数量减一
func UncheckComment(id int, data *Comment) int {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = db.Model(&comment).Where("id = ?", id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id = ?", res.ArticleID).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
