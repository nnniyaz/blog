package articleHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nnniyaz/blog/internal/domain/article"
	response2 "github.com/nnniyaz/blog/internal/handlers/http/response"
	"github.com/nnniyaz/blog/internal/services/article"
	"github.com/nnniyaz/blog/pkg/logger"
	"net/http"
	"time"
)

type HttpDelivery struct {
	logger  logger.Logger
	service articleService.ApplicationService
}

func NewHttpDelivery(l logger.Logger, s articleService.ApplicationService) *HttpDelivery {
	return &HttpDelivery{logger: l, service: s}
}

type CreateArticleIn struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// CreateArticle godoc
//
//	@Summary		Create article
//	@Description	This can only be done by the logged-in user.
//	@Tags			Articles
//	@Accept			json
//	@Produce		json
//	@Param			data		body		CreateArticleIn		true	"Create Article Structure"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/article [get]
func (hd *HttpDelivery) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var in CreateArticleIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Create(r.Context(), in.Title, in.Content)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, nil)
}

type UpdateArticleIn struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// UpdateArticle godoc
//
//	@Summary		Update article
//	@Description	This can only be done by the logged-in user.
//	@Tags			Articles
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string			true	"Article ID"
//	@Param			data		body		UpdateArticleIn		true	"Update Article Structure"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/article/{id} [put]
func (hd *HttpDelivery) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var in UpdateArticleIn
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	err := hd.service.Update(r.Context(), id, in.Title, in.Content)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, nil)
}

// DeleteArticle godoc
//
//	@Summary		Delete article
//	@Description	This can only be done by the logged-in user.
//	@Tags			Articles
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Article ID"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/article/{id} [delete]
func (hd *HttpDelivery) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Delete(r.Context(), id)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, nil)
}

// RestoreArticle godoc
//
//	@Summary		Restore article
//	@Description	This can only be done by the logged-in user.
//	@Tags			Articles
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Article ID"
//	@Success		200						{object}	response.Success
//	@Failure		default					{object}	response.Error
//	@Router			/article/restore/{id} [put]
func (hd *HttpDelivery) RestoreArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := hd.service.Restore(r.Context(), id)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, nil)
}

type GetArticleOut struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	IsDeleted bool   `json:"isDeleted"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Version   int    `json:"version"`
}

func newGetArticleOut(a *article.Article) *GetArticleOut {
	return &GetArticleOut{
		Id:        a.GetId().String(),
		Title:     a.GetTitle(),
		Content:   a.GetContent(),
		IsDeleted: a.GetIsDeleted(),
		CreatedAt: a.GetCreatedAt().Format(time.RFC3339),
		UpdatedAt: a.GetUpdatedAt().Format(time.RFC3339),
		Version:   a.GetVersion(),
	}
}

// GetArticle godoc
//
//	@Summary		Get article
//	@Description	This can only be done by the logged-in user.
//	@Tags			Articles
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"Article ID"
//	@Success		200						{object}	response.Success{GetArticleOut}
//	@Failure		default					{object}	response.Error
//	@Router			/article/{id} [get]
func (hd *HttpDelivery) GetArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	result, err := hd.service.FindById(r.Context(), id)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, newGetArticleOut(result))
}

type GetAllArticleOut struct {
	Articles []*GetArticleOut `json:"articles"`
	Count    int64            `json:"count"`
}

func newGetAllArticleOut(articles []*article.Article, count int64) *GetAllArticleOut {
	var out []*GetArticleOut
	for _, a := range articles {
		out = append(out, newGetArticleOut(a))
	}
	return &GetAllArticleOut{Articles: out, Count: count}
}

// GetAllArticles godoc
//
//	@Summary		Get all articles
//	@Description	This can only be done by the logged-in user.
//	@Tags			Articles
//	@Accept			json
//	@Produce		json
//	@Param			offset		query		int		false	"Offset"
//	@Param			limit		query		int		false	"Limit"
//	@Param			is_deleted	query		bool	false	"Is deleted"
//	@Param			search		query		string	false	"Search"
//	@Success		200						{object}	response.Success{GetAllArticleOut}
//	@Failure		default					{object}	response.Error
//	@Router			/article [get]
func (hd *HttpDelivery) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	offset := r.Context().Value("offset").(int64)
	limit := r.Context().Value("limit").(int64)
	isDeleted := r.Context().Value("is_deleted").(bool)
	search := r.Context().Value("search").(string)
	articles, count, err := hd.service.FindAll(r.Context(), offset, limit, isDeleted, search)
	if err != nil {
		response2.NewError(hd.logger, w, r, err)
		return
	}
	response2.NewSuccess(hd.logger, w, r, newGetAllArticleOut(articles, count))
}
