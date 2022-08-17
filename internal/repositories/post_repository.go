package repositories

import (
	"github.com/yudegaki/apprunner-memoapp/internal/entities"
	"gorm.io/gorm"
)

type PostRepository struct {
	Conn *gorm.DB
}

type Post struct {
	gorm.Model
	Title string `gorm:"type:varchar(50);not null"`
	Body  string `gorm:"type:varchar(100);not null"`
}

func NewPostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{Conn: conn}
}

func convertRepositoryPostToEntityPost(post *Post) *entities.Post {
	var entityPost entities.Post
	entityPost.SetTitle(post.Title)
	entityPost.SetBody(post.Body)

	return &entityPost
}

func convertRepositoryPostsToEntityPosts(post []*Post) []*entities.Post {
	var entityPosts []*entities.Post
	for _, post := range post {
		entityPosts = append(entityPosts, convertRepositoryPostToEntityPost(post))
	}
	return entityPosts
}

func (r *PostRepository) GetAll() ([]*entities.Post, error) {
	var dbPosts []*Post
	if err := r.Conn.Find(&dbPosts).Error; err != nil {
		return nil, err
	}
	return convertRepositoryPostsToEntityPosts(dbPosts), nil
}

func (r *PostRepository) GetDetail(id uint) (*entities.Post, error) {
	var dbPost Post
	if err := r.Conn.First(&dbPost, id).Error; err != nil {
		return nil, err
	}
	return convertRepositoryPostToEntityPost(&dbPost), nil
}

func (r *PostRepository) Create(post *entities.Post) error {
	if err := r.Conn.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Update(post *entities.Post) error {
	// NOTE: Updatesは構造体の場合はdefault値を更新しない
	// https://gorm.io/docs/update.html#Updates-multiple-columns
	if err := r.Conn.Updates(&post).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Delete(id uint) error {
	if err := r.Conn.Delete(&Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
