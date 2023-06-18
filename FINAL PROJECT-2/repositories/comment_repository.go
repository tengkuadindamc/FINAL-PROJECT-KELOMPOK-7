package repositories

import (
	"fp-2/entities"
	"fp-2/pkg/errs"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(user *entities.User, comment *entities.Comment) (*entities.Comment, errs.MessageErr)
	GetCommentById(commentId uint) (*entities.Comment, errs.MessageErr)
	GetCommentsByUserId(userId uint) ([]entities.Comment, errs.MessageErr)
	GetAllComments() ([]entities.Comment, errs.MessageErr)
	GetCommentsByPhotoId(photoId uint) ([]entities.Comment, errs.MessageErr)
	UpdateComment(comment *entities.Comment, cmtUpdate *entities.Comment) (*entities.Comment, errs.MessageErr)
	DeleteComment(comment *entities.Comment) errs.MessageErr
}

type commentPG struct {
	db *gorm.DB
}

func NewCommentPG(db *gorm.DB) CommentRepository {
	return &commentPG{db: db}
}

func (c *commentPG) CreateComment(user *entities.User, comment *entities.Comment) (*entities.Comment, errs.MessageErr) {
	comment.UserId = user.ID
	if err := c.db.Create(comment).Error; err != nil {
		return nil, errs.NewInternalServerError("Can't create comment")
	}

	return comment, nil
}

func (c *commentPG) GetCommentById(commentId uint) (*entities.Comment, errs.MessageErr) {
	var comment entities.Comment
	err := c.db.First(&comment, commentId).Error

	if err != nil {
		return nil, errs.NewNotFound("comment not found")
	}
	return &comment, nil
}

func (c *commentPG) GetCommentsByUserId(userId uint) ([]entities.Comment, errs.MessageErr) {
	var comment []entities.Comment
	err := c.db.Find(&comment, "user_id=?", userId).Error

	if err != nil {
		return nil, errs.NewNotFound("can't get all comments of this user")
	}
	return comment, nil
}

func (r *commentPG) GetCommentsByPhotoId(photoId uint) ([]entities.Comment, errs.MessageErr) {
	var comments []entities.Comment
	err := r.db.Where("photo_id = ?", photoId).Find(&comments).Error

	if err != nil {
		return nil, errs.NewNotFound("can't get all comments of this photo")
	}
	return comments, nil
}

func (c *commentPG) GetAllComments() ([]entities.Comment, errs.MessageErr) {
	var comments []entities.Comment

	if err := c.db.Find(&comments).Error; err != nil {
		return nil, errs.NewInternalServerError("can't get all comments")
	}
	return comments, nil
}

func (c *commentPG) UpdateComment(comment *entities.Comment, cmtUpdate *entities.Comment) (*entities.Comment, errs.MessageErr) {
	if err := c.db.Model(comment).Where("id = ?", comment.ID).Updates(cmtUpdate).Error; err != nil {
		return nil, errs.NewNotFound("can't update this comment")
	}

	return comment, nil
}

func (c *commentPG) DeleteComment(comment *entities.Comment) errs.MessageErr {
	err := c.db.Delete(comment).Error

	if err != nil {
		return errs.NewInternalServerError("can't delete this comment")
	}
	return nil
}
