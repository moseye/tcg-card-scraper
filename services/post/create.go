package post

import "tcg-scraper/models"

func (postService *Service) Create(post *models.Post) {
	postService.DB.Create(post)
}
