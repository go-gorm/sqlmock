package main

import (
	"gorm.io/gorm"
)

// Post model
type Post struct {
	gorm.Model
	Title string
	Body  string
}

// Service for post model
type Service struct {
	db *gorm.DB
}

// Create post
func (service *Service) Create(post *Post) error {

	return service.db.Create(post).Error
}

// Get post
func (service *Service) Get(id uint) (post *Post, err error) {
	post = &Post{}
	err = service.db.First(post, "id = ?", id).Error
	return
}

// Update post
func (service *Service) Update(post *Post, data map[string]interface{}) error {

	return service.db.Model(post).Updates(data).Error
}

// Destroy post
func (service *Service) Destroy(id uint) error {

	return service.db.Delete(&Post{}, "id = ?", id).Error
}
