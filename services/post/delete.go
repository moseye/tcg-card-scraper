package post

import "tcg-scraper/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}
