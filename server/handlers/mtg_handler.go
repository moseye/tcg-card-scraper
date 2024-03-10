package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"tcg-scraper/models"
	"tcg-scraper/repositories"
	"tcg-scraper/responses"
	s "tcg-scraper/server"
)

type MtgHandler struct {
	server *s.Server
}

func NewMtgHandler(server *s.Server) *MtgHandler {
	return &MtgHandler{server: server}
}

// GetCardDetails godoc
//
//	@Summary		Get Magic:The Gathering Card Details
//	@Description	Returns details on the card closest to the search term
//	@ID				mtg-get
//	@Tags			MTG Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.CreatePostRequest	true	"Post title and content"
//	@Success		201		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/mtg [get]
func (p *MtgHandler) GetCardDetails(c echo.Context) error {
	var posts []models.Post

	postRepository := repositories.NewPostRepository(p.server.DB)
	postRepository.GetPosts(&posts)

	response := responses.NewPostResponse(posts)
	return responses.Response(c, http.StatusOK, response)
}

// DeletePost godoc
//
//	@Summary		Delete post
//	@Description	Delete post
//	@ID				posts-delete
//	@Tags			Posts Actions
//	@Param			id	path		int	true	"Post ID"
//	@Success		204	{object}	responses.Data
//	@Failure		404	{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/posts/{id} [delete]
//func (p *PostHandlers) DeletePost(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	post := models.Post{}
//
//	postRepository := repositories.NewPostRepository(p.server.DB)
//	postRepository.GetPost(&post, id)
//
//	if post.ID == 0 {
//		return responses.ErrorResponse(c, http.StatusNotFound, "Post not found")
//	}
//
//	postService := postservice.NewPostService(p.server.DB)
//	postService.Delete(&post)
//
//	return responses.MessageResponse(c, http.StatusNoContent, "Post deleted successfully")
//}

// GetPosts godoc
//
//	@Summary		Get posts
//	@Description	Get the list of all posts
//	@ID				posts-get
//	@Tags			Posts Actions
//	@Produce		json
//	@Success		200	{array}	responses.PostResponse
//	@Security		ApiKeyAuth
//	@Router			/posts [get]
//func (p *PostHandlers) GetPosts(c echo.Context) error {
//	var posts []models.Post
//
//	postRepository := repositories.NewPostRepository(p.server.DB)
//	postRepository.GetPosts(&posts)
//
//	response := responses.NewPostResponse(posts)
//	return responses.Response(c, http.StatusOK, response)
//}

// UpdatePost godoc
//
//	@Summary		Update post
//	@Description	Update post
//	@ID				posts-update
//	@Tags			Posts Actions
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Post ID"
//	@Param			params	body		requests.UpdatePostRequest	true	"Post title and content"
//	@Success		200		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Failure		404		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/posts/{id} [put]
//func (p *PostHandlers) UpdatePost(c echo.Context) error {
//	updatePostRequest := new(requests.UpdatePostRequest)
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	if err := c.Bind(updatePostRequest); err != nil {
//		return err
//	}
//
//	if err := updatePostRequest.Validate(); err != nil {
//		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
//	}
//
//	post := models.Post{}
//
//	postRepository := repositories.NewPostRepository(p.server.DB)
//	postRepository.GetPost(&post, id)
//
//	if post.ID == 0 {
//		return responses.ErrorResponse(c, http.StatusNotFound, "Post not found")
//	}
//
//	postService := postservice.NewPostService(p.server.DB)
//	postService.Update(&post, updatePostRequest)
//
//	return responses.MessageResponse(c, http.StatusOK, "Post successfully updated")
//}
